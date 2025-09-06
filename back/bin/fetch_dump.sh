#!/bin/bash
set -euo pipefail

# Usage:
#   ./fetch_dump.sh <ssh_host> <remote_path> [local_parent_dir] [options]
#
#
# Options:
#   --latest          # si <remote_path> est un dossier: choisir le .tar.gz le plus récent
#   --db <name>       # filtre par préfixe de base: <name>_*.tar.gz
#   --scp             # utiliser scp au lieu de rsync (fallback)
#
#
# Usage notes:
#   ./fetch_dump.sh nseven /home/nseven/backups [local_dir] --latest --db tenjin_preprod
#   ./fetch_dump.sh nseven /home/nseven/backups [local_dir] --latest --db tenjin_prod
#
# Exemples:
#   # 1) Télécharger un fichier dump précis
#   ./fetch_dump.sh user@serveur /home/user/backups/tenjin_20250906-174530.tar.gz
#
#   # 2) Prendre le dernier dump (*.tar.gz) d'un répertoire
#   ./fetch_dump.sh user@serveur /home/user/backups "$HOME/backups" --latest
#
#   # 3) Prendre le dernier dump d'une base spécifique (préfixe "tenjin_")
#   ./fetch_dump.sh user@serveur /home/user/backups "$HOME/backups" --latest --db tenjin
#
#
# Notes:
# - Préfère un chemin ABSOLU côté serveur (évite les ambiguïtés avec ~).
# - Le dossier local par défaut est $HOME/backups.
# - Nécessite rsync; sinon ajoute l’option --scp.

SSH_HOST=${1:-}
REMOTE_PATH_IN=${2:-}
LOCAL_PARENT=${3:-"$HOME/backups"}
shift $(( $#>=3 ? 3 : $# )) || true

if [[ -z "$SSH_HOST" || -z "$REMOTE_PATH_IN" ]]; then
  echo "[ERROR] Usage: $0 <ssh_host> <remote_path> [local_parent_dir] [--latest] [--db <name>] [--scp]"
  exit 1
fi

# Options
USE_LATEST=false
DB_FILTER=""
USE_SCP=false
while [[ $# -gt 0 ]]; do
  case "$1" in
    --latest) USE_LATEST=true; shift ;;
    --db) DB_FILTER="${2:-}"; shift 2 ;;
    --scp) USE_SCP=true; shift ;;
    *)
      echo "[WARNING] Option inconnue: $1"
      exit 1
      ;;
  esac
done

# Normalise le dossier local (gère ~ en tête)
LOCAL_PARENT="${LOCAL_PARENT/#\~/$HOME}"
mkdir -p "$LOCAL_PARENT"

# Détecte si remote_path est un dossier ou un fichier, à distance
echo "[INFO] Inspection côté serveur..."
# shellcheck disable=SC2029
if ssh "$SSH_HOST" "test -d '$REMOTE_PATH_IN'"; then
  REMOTE_IS_DIR=true
elif ssh "$SSH_HOST" "test -f '$REMOTE_PATH_IN'"; then
  REMOTE_IS_DIR=false
else
  echo "[ERROR] Chemin distant introuvable: $REMOTE_PATH_IN"
  exit 1
fi

REMOTE_FILE=""

if [[ "$REMOTE_IS_DIR" == true ]]; then
  if [[ "$USE_LATEST" != true ]]; then
    echo "[ERROR] Le chemin distant est un dossier. Ajoute --latest (et éventuellement --db <name>) pour récupérer le dernier dump."
    exit 1
  fi

  # Construit le motif de recherche
  # Motif par défaut: *.tar.gz ; si DB_FILTER -> <db>_*.tar.gz
  PATTERN="*.tar.gz"
  if [[ -n "$DB_FILTER" ]]; then
    PATTERN="${DB_FILTER}_*.tar.gz"
  fi

  echo "[INFO] Recherche du dernier fichier correspondant à '$PATTERN' dans '$REMOTE_PATH_IN'..."
  # Liste par date décroissante, prend le premier
  REMOTE_FILE=$(ssh "$SSH_HOST" "ls -1t \"${REMOTE_PATH_IN%/}/\"$PATTERN 2>/dev/null | head -n1")

  if [[ -z "$REMOTE_FILE" ]]; then
    echo "[ERROR] Aucun fichier ne correspond au motif dans $REMOTE_PATH_IN"
    exit 1
  fi
  echo "[INFO] Dernier dump détecté: $REMOTE_FILE"
else
  REMOTE_FILE="$REMOTE_PATH_IN"
  echo "[INFO] Fichier à rapatrier: $REMOTE_FILE"
fi

# Téléchargement
BASENAME=$(basename "$REMOTE_FILE")
DEST="$LOCAL_PARENT/$BASENAME"

echo "[INFO] Téléchargement vers: $DEST"

if [[ "$USE_SCP" == true ]]; then
  # Fallback scp
  scp -p "$SSH_HOST:$REMOTE_FILE" "$DEST"
else
  # rsync (reprend en cas de coupure, affiche la progression)
  rsync -avz --partial --progress "$SSH_HOST:$REMOTE_FILE" "$DEST"
fi

echo "[SUCCESS] Terminé."
echo "[INFO] Fichier local: $DEST"
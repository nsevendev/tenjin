#!/bin/bash
set -euo pipefail

# Usage:
#   ./mongo_restore.sh <target_container> <target_db> <archive.tar.gz> [options]
#
# Options:
#   --replace                     # drop les collections avant restauration (équivaut à --drop)
#   --username <user>             # utilisateur d'auth Mongo
#   --password <pass>             # mot de passe d'auth
#   --authenticationDatabase <db> # base d'auth (souvent 'admin')
#   --uri <mongodb_uri>           # URI complet (si fourni, on n'utilise pas --network=container)
#   --tools-image <img>           # image docker contenant mongorestore (defaut: mongo:6.0)
#
# Usage notes:
#   ./mongo_restore.sh tenjin_dev_db tenjin_dev <path to dump.tar.gz> --replace --tools-image mongo:7
#   ./mongo_restore.sh tenjin_preprod_db tenjin_preprod <path to dump.tar.gz> --replace --tools-image mongo:7
#   ./mongo_restore.sh tenjin_prod_db tenjin_prod <path to dump.tar.gz> --replace --tools-image mongo:7
#
# Exemples:
#   # Restaure en fusion (par défaut) via le réseau du container cible
#   ./mongo_restore.sh mongo_target mydb ./backups/mydb_20250906-174530.tar.gz
#
#   # Restaure en remplaçant les collections existantes
#   ./mongo_restore.sh mongo_target mydb ./backups/mydb_...tar.gz --replace
#
#   # Avec authentification
#   ./mongo_restore.sh mongo_target mydb dump.tar.gz --username myuser --password secret --authenticationDatabase admin
#
#   # En se connectant via une URI (pas besoin de partager le réseau du container)
#   ./mongo_restore.sh _ignored mydb dump.tar.gz --uri "mongodb://user:pass@host:27017/?authSource=admin"

TARGET_CONTAINER=${1:-}
TARGET_DB=${2:-}
ARCHIVE_PATH=${3:-}
shift 3 || true

if [[ -z "${TARGET_CONTAINER}" || -z "${TARGET_DB}" || -z "${ARCHIVE_PATH}" ]]; then
  echo "[ERROR] Usage: $0 <target_container> <target_db> <archive.tar.gz> [options]"
  exit 1
fi

# Defaults
REPLACE=false
USERNAME=""
PASSWORD=""
AUTH_DB=""
URI=""
TOOLS_IMAGE="mongo:6.0"   # Image avec les mongo-tools (mongorestore). Ajuste si besoin.

# Parse options
while [[ $# -gt 0 ]]; do
  case "$1" in
    --replace) REPLACE=true; shift ;;
    --username) USERNAME="$2"; shift 2 ;;
    --password) PASSWORD="$2"; shift 2 ;;
    --authenticationDatabase) AUTH_DB="$2"; shift 2 ;;
    --uri) URI="$2"; shift 2 ;;
    --tools-image) TOOLS_IMAGE="$2"; shift 2 ;;
    *)
      echo "[WARNING] Option inconnue: $1"
      exit 1
      ;;
  esac
done

if [[ ! -f "$ARCHIVE_PATH" ]]; then
  echo "[ERROR] Archive introuvable: $ARCHIVE_PATH"
  exit 1
fi

# Dossier temporaire d'extraction
TMP_DIR="$(mktemp -d -t mongorestore-XXXXXX)"
cleanup() {
  rm -rf "$TMP_DIR" || true
}
trap cleanup EXIT

echo "[INFO] Extraction de l'archive: $ARCHIVE_PATH"
tar -xzf "$ARCHIVE_PATH" -C "$TMP_DIR"

# L'archive créée par notre script de dump contient un dossier de la forme <db>_<timestamp>
# On détecte ce dossier (il doit contenir les .bson/.metadata.json)
RESTORE_SUBDIR=""
# On prend le premier dossier trouvé dans le TMP_DIR
for d in "$TMP_DIR"/*; do
  if [[ -d "$d" ]]; then
    RESTORE_SUBDIR="$d"
    break
  fi
done

if [[ -z "$RESTORE_SUBDIR" ]]; then
  echo "[ERROR] Impossible de trouver le dossier de dump dans l'archive."
  exit 1
fi

echo "[INFO] Dossier de données à restaurer: $RESTORE_SUBDIR"

# Prépare les flags mongorestore
RESTORE_FLAGS=( "--db=$TARGET_DB" "--dir=/restore/$(basename "$RESTORE_SUBDIR")" "--stopOnError" )
$REPLACE && RESTORE_FLAGS+=( "--drop" )
[[ -n "$USERNAME" ]] && RESTORE_FLAGS+=( "--username=$USERNAME" )
[[ -n "$PASSWORD" ]] && RESTORE_FLAGS+=( "--password=$PASSWORD" )
[[ -n "$AUTH_DB"  ]] && RESTORE_FLAGS+=( "--authenticationDatabase=$AUTH_DB" )

# Exécution de mongorestore via un container outillage (n'impose pas que le binaire soit dans le container cible)
# Deux modes:
#  1) Sans URI: on partage le namespace réseau du container cible -> --host=localhost
#  2) Avec URI: on se connecte via l'URI (pas de --network=container)
if [[ -z "$URI" ]]; then
  echo "[INFO] Connexion via le réseau du container: $TARGET_CONTAINER"
  docker run --rm \
    --network="container:${TARGET_CONTAINER}" \
    -v "$TMP_DIR:/restore:ro" \
    "$TOOLS_IMAGE" \
    bash -lc "mongorestore --host=localhost ${RESTORE_FLAGS[*]}"
else
  echo "[INFO] Connexion via URI fournie"
  docker run --rm \
    -v "$TMP_DIR:/restore:ro" \
    "$TOOLS_IMAGE" \
    bash -lc "mongorestore --uri='$URI' ${RESTORE_FLAGS[*]}"
fi

echo "[SUCCESS] Restauration terminée."
echo "[INFO] Mode: $([ "$REPLACE" = true ] && echo 'REPLACE (--drop)' || echo 'MERGE (sans drop)')"
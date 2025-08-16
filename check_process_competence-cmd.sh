#!/bin/bash

echo "=== 🔍 Status Compétences ==="
echo "$(date)"
echo ""

# Status du processus
PIDS=$(ps aux | grep "list-competence-complet" | grep -v grep | grep -v "check_process" | awk '{print $2}' | xargs)

if [ -n "$PIDS" ]; then
    echo "✅ Processus en cours (PID: $PIDS)"
    # Compter le nombre de processus
    NB_PROCESS=$(echo $PIDS | wc -w)
    if [ "$NB_PROCESS" -gt 1 ]; then
        echo "⚠️  Attention: $NB_PROCESS processus détectés (peut-être des doublons)"
    fi
else
    echo "❌ Processus arrêté"
fi

echo ""

# Vérifier si le fichier log existe
if [ -f competences.log ]; then
    echo "Statistiques du log :"
    echo "Lignes totales: $(wc -l < competences.log)"
    echo "Taille du fichier: $(du -h competences.log | cut -f1)"
    echo ""

    # Compteurs d'erreurs et de succès
    echo "🎯 Progression :"
    ERREURS=$(grep -c "Erreur\|FATAL\|ERROR" competences.log 2>/dev/null || echo "0")
    SUCCES=$(grep -c "bien ajoutée" competences.log 2>/dev/null || echo "0")
    PROGRESSION=$(grep "Progression:" competences.log | tail -1 2>/dev/null || echo "Aucune progression trouvée")

    echo "Succès: $SUCCES"
    echo "Erreurs: $ERREURS"
    echo "Dernière progression: $PROGRESSION"

    # Calcul du taux de succès si on a des données
    if [ "$SUCCES" -gt 0 ] || [ "$ERREURS" -gt 0 ]; then
        TOTAL=$((SUCCES + ERREURS))
        TAUX=$(echo "scale=1; $SUCCES * 100 / $TOTAL" | bc -l 2>/dev/null || echo "N/A")
        echo "  📈 Taux de succès: ${TAUX}%"
    fi

    echo ""

    # Dernières erreurs s'il y en a
    DERNIERES_ERREURS=$(grep "Erreur\|FATAL\|ERROR" competences.log | tail -3)
    if [ -n "$DERNIERES_ERREURS" ]; then
        echo "Dernières erreurs :"
        echo "$DERNIERES_ERREURS" | sed 's/^/  /'
        echo ""
    fi

    echo "Dernières lignes du log :"
    tail -n 5 competences.log | sed 's/^/  /'

else
    echo "Fichier log (competences.log) introuvable"
fi

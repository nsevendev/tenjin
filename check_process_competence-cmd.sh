#!/bin/bash

echo "=== 🔍 Status Compétences ==="
echo "📅 $(date)"
echo ""

# Vérification propre avec pgrep
if pgrep -f "list-competence-complet" > /dev/null; then
    PID=$(pgrep -f "list-competence-complet" | head -1)  # ← Prendre seulement le premier
    echo "✅ Processus en cours (PID: $PID)"
else
    echo "❌ Processus arrêté"
fi

echo ""

if [ -f competences.log ]; then  # ← Changé selon votre fichier
    echo "📊 Statistiques :"
    LIGNES=$(wc -l < competences.log)
    TAILLE=$(du -h competences.log | cut -f1)
    echo "  📄 Lignes: $LIGNES"
    echo "  📈 Taille: $TAILLE"
    echo ""
    
    echo "🎯 Progression :"
    
    # CORRECTION : Nettoyer les retours à la ligne
    SUCCES=$(grep -c "bien ajoutée" competences.log 2>/dev/null | tr -d '\n' | head -1)
    ERREURS=$(grep -c -E "(Erreur|FATAL|ERROR)" competences.log 2>/dev/null | tr -d '\n' | head -1)
    
    # Valeurs par défaut si vide
    SUCCES=${SUCCES:-0}
    ERREURS=${ERREURS:-0}
    
    echo "  ✅ Succès: $SUCCES"
    echo "  ❌ Erreurs: $ERREURS"
    
    # Calcul sécurisé
    if [ "$SUCCES" -gt 0 ] 2>/dev/null && [ "$ERREURS" -ge 0 ] 2>/dev/null; then
        TOTAL=$((SUCCES + ERREURS))
        if [ "$TOTAL" -gt 0 ]; then
            TAUX=$((SUCCES * 100 / TOTAL))
            echo "  📈 Taux de succès: ${TAUX}%"
        fi
    fi
    
    echo ""
    echo "🔄 Dernière progression :"
    tail -n 50 competences.log | grep "Progression:" | tail -1 | sed 's/^/  /' || echo "  Aucune progression trouvée"
    
    echo ""
    echo "⚠️  Dernières erreurs :"
    DERNIERES_ERREURS=$(tail -n 30 competences.log | grep -E "(Erreur|FATAL|ERROR)" | tail -3)
    if [ -n "$DERNIERES_ERREURS" ]; then
        echo "$DERNIERES_ERREURS" | sed 's/^/  /'
    else
        echo "  Aucune erreur récente"
    fi
    
    echo ""
    echo "📋 Dernières lignes :"
    tail -n 3 competences.log | sed 's/^/  /'
    
else
    echo "❌ Pas de fichier log (competences.log)"
fi

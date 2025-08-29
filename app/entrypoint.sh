#!/bin/sh

if [ ! -d "node_modules" ]; then
    echo "Dossier node_modules introuvable, installation des dépendances..."
    npm ci
else
    echo "Dossier node_modules trouvé, installation ignorée"
fi

exec "$@"

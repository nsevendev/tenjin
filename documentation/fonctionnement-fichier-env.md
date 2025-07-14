# Fonctionnement fichier `.env`
- il y a 3 fichiers `.env` :
    - `.env` : ce fichier à la racine du repo, sert pour le lancement des containeurs docker  
      et contient des infos pour traefik et les ports utilisés par les services (tous les conteneurs)
    - `back/.env` : contient des variable specifiques à l'api
    - `front/.env` : contient des variable specifiques à l'aoolication front end.

**ATTENTION les variables APP_ENV doivent toutes etre avec la meme valeur**  

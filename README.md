# Tenjin

> projet de suivi de formation

## Les projets
- back : api principal en go
- front : application front en qwik partie commerciale
- app : application front en angular partie interne

## Avant installation en mode dev
- Treafik nseven doit etre installé, configuré et lancé
- Creer des copies `.env` depuis `.env.dist` à l'endroit où sont les `.env.dist`
- renseigner les valeurs `1270.0.0.1  tenjin-app.local`, `1270.0.0.1  tenjin-api.local`  
  `127.0.0.1  tenjin.local` dans le fichier `/etc/hosts`  
  (attention utiliser `sudo` pour modifier ce fichier)
- copier/coller `app/src/environment/environment.dist` en `app/src/environment/environment.ts`  
  renseigné laisser les variables existante, celle sans valeur demandé au lead les valeurs manquantes

## Avant d'installer en mode prod ou preprod
- pareil que le mode dev, sauf que les variables dans les `.env` doivent etre renseigner pour la production
  utiliser la preprod comme réference

## Mode dev
- lancer la commande `make up`
- accèder aux logs (commande `make logs-api`) de l'api pour voir l'url du swagger ou l'url de l'api
- accèder aux logs (commande `make logs-front`) de l'application front end pour voir l'url de l'application
- accèder à l'application front end à l'adresse `https://tenjin.local`
- voir toutes les commandes disponibles avec `make`

## Mode prod ou preprod
- mettre les variables des `.env` `APP_ENV=prod` ou `APP_ENV=preprod`
- lancer les commandes comme pour le mode dev, mais renseigner les variables d'environement pour la prod ou preprod

## Lancement des tests
> partie backend (api)
- à la racine du dossier `back` copier/coller le `.env.test.dist` et renommer le en `.env.test`, renseigner les variables vide  
- utiliser les commandes `make` => taper `make` dans le terminal pour voir toute les commandes de test disponibles   

> partie frontend

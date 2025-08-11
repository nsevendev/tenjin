# Tenjin

> projet de suivi de formation

## Avant installation en mode dev
- Treafik nseven doit etre installé, configuré et lancé
- Creer des copies `.env` depuis `.env.dist` à l'endroit où sont les `.env.dist`
- renseigner les valeurs `1270.0.0.1  tenjin.local`, `1270.0.0.1  tenjin-api.local`  
  dans le fichier `/etc/hosts` (attention utiliser `sudo` pour modifier ce fichier)

## Avant d'installer en mode prod ou preprod
- pareil que le mode dev, sauf que les variables dans les `.env` doivent etre renseigner pour la production

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
- creer un .env.test à la racine du dossier `back` mettre les variables APP_ENV=test et VARIABLE_TEST=test  
- utiliser les commandes `make` taper `make` taper dans le terminal pour voir toute les commandes  

> partie frontend

# API ROME Sync – Commandes CLI

Ce CLI permet de **générer les structs Go** à partir de l’OpenAPI ROME,
et de **récupérer, stocker, enrichir** toutes les fiches métiers, compétence ROME en JSON pour exploitation dans l’application.

---

## **Prérequis**

* Go installé (`go version`)
* Le binaire `oapi-codegen` (installer : `go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest`)
> le docker l'install automatiquement
* Avoir accès à l’API ROME France Travail (token/credentials) tout est dans le `.env` du projet coté back
* les fichiers json pour lancer l'ajout dans la bdd. il faut mettre les fichiers dans les dossier : demander les fichiers à l'équipe dev
* `./apirome/version40/data/listmetierdetail/` pour les détails des métiers

---

## **Commandes disponibles**

### `help`

**Affiche l'aide des commandes**

```sh
# makefile
make apir cmd="help"

# shell
go run ./apirome/main.go help 
```

### 1. `generate-model`

**Génère les structs Go (modèle OpenAPI) pour l’API Rome**

```sh
# makefile
make apir cmd="generate-model --input ~/Downloads/rome_4.0_openapi.json --output internal/rome40openapi/model.go --package rome40openapi"

# shell
go run ./apirome/main.go generate-model \
  --input <chemin_fichier_openapi.json> \
  --output <chemin_sortie_go> \
  --package <nom_du_package> # nom du dossier dans lequel seront générées les structs
```

* **--input** : chemin du fichier OpenAPI (JSON ou YAML)
* **--output** : fichier `.go` généré
* **--package** : nom du package Go pour les structs

---

### 2. `list-metier-summary`

**Télécharge le résumé de tous les métiers ROME**
(code + libellé) et l’enregistre dans un fichier JSON.

```sh
# makefile
make apir cmd="list-metier-summary"

# shell
go run main.go list-metier-summary
```

* Fichier créé : `./apirome/version40/data/listmetier/<datetimestamp>.json`
* À utiliser comme base pour télécharger les détails des métiers.

---

### 3. `list-metier-detail`

**Télécharge les détails complets de chaque métier (par code),**
et enregistre le tout dans un gros fichier JSON (un tableau de métiers détaillés).

```sh
# makefile
make apir cmd="list-metier-detail"

# shell
go run main.go list-metier-detail
```

* Récupère le dernier fichier “summary” dans `./apirome/version40/data/listmetier/`
* Pour chaque code, télécharge le détail et l’ajoute à un tableau
* Fichier créé : `./apirome/version40/data/metierdetail/<datetimestamp>.json`
* L’opération prend du temps (plusieurs minutes, dépend du nombre de métiers)

---

## **Workflow**

1. **Générer les structs Go (si besoin)**

  * Utiliser `generate-model` si le schéma OpenAPI évolue

2. **Récupérer la liste des métiers (summary)**

  * `go run main.go list-metier-summary`

3. **Récupérer les fiches métiers détaillées**

  * `go run main.go list-metier-detail`

4. **Exploiter les fichiers JSON générés pour importer ou mapper dans l’app dans la bdd**

## **Workflow**

## METIER

- Prérequis:
    - le `.env` doit avoir les clef api rome à jour

1.**Récupérer la liste des métiers (summary)**

  * `make apir cmd=list-metier-summary` commande rapide

2.**Récupérer les fiches métiers détaillées**

  * `make apir cmd=list-metier-detail` commande longue (3h)

3.**Exploiter les fichiers JSON générés pour importer ou mapper dans l’app dans la bdd**

  * `make apir cmd=import-metier-detail-database` commande rapide

## RNCP

- Prérequis:
    - Avoir le fichier XML RNCP téléchargé depuis la plateforme du gouvernement [ici](https://www.data.gouv.fr/datasets/repertoire-national-des-certifications-professionnelles-et-repertoire-specifique/)
    - Mettre le fichier XML téléchargé dans le dossier `apirome/rncp/data/`
    - Avoir les structs GO à jour avec le fichier XML, les structs sont dans `apirome/rncp/model.go`

1.**Transformer le fichier XML rncp téléchargé depuis la plateform du gouvernement en json**  

  * `make apir cmd=rncp-xml-to-json` commande rapide  

2.**Importer le json précedement créé dans la base de données**  

  * `make apir cmd=import-rncp-database` commande rapide     

## COMPETENCES  

- Prérequis:  
    - le `.env` doit avoir les clef api rome à jour  

1.**Récupérer la liste des compétences (summary)**  

  * `make apir cmd=list-competence` commande rapide  

2.**Récupérer les fiches compétences complete**  

  * `make apir cmd=list-competence-complet` commande longue (20h) (commande en cours de changement peut ne pas fonctionner correctement)  
  * `nohup make apir cmd=list-competence-complet > competences.log 2>&1 < /dev/null &` commande alternative pour que sa fonctionne  

3.**Exploiter les fichiers JSON générés pour importer dans l’app dans la bdd**  

  * `make apir cmd=import-competence-complet-database` commande moyenne  

> pour les logs de la commande de l'étape 2 il y a un script à la racine du projet `check_process_competence-cmd.sh`   
> qui permet de voir les logs de la commande en cours d'exécution.  

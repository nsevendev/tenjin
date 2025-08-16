### mettre les fichier json des data dans :

- rncp/data => fichier xml venant du gouvernement
  (lien ici)[https://www.data.gouv.fr/datasets/repertoire-national-des-certifications-professionnelles-et-repertoire-specifique/]
> permet de creer le ficier json pour mettre en bdd

- rncp/data => fichier json
> permet d'enregistrer les données du rncp dans la base de données

### commandes : 

- ImportJsonInDatabaseCmd
> permet d'importer les données du rncp dans la base de données

- RncpXmlToJsonCmd
> permet de convertir le fichier xml du rncp en json

### flow : 

- si pas de xml rncp demander à un admin de le fournir
- si il y a le fichier xml lancer la commande `RncpXmlToJsonCmd` cela creera le fichier json dans le dossier rncp/data
- si vous avez le json rncp lancer la commande `ImportJsonInDatabaseCmd` pour importer les données dans la base de données


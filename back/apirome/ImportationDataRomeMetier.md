### mettre les fichier json des data dans :

- listmetiersummary => fichier json
> permet de faire la request api pour recuperer le json des métiers

- listmetierdetail => fichier json
> permet d'enregistrer les données des métiers dans la base de données

### commandes : 

- ListMetierSummaryCmd
> recupere le json des métiers (summary) directement depuis l'api rome

- ListMetierDetailCmd
> à l'aide du fichier json listmetiersummary il va recuperer le json des métiers

- ImportDataListMetierDetailCmd
> permet d'importer les données des métiers dans la base de données via le json listmetierdetail

### flow : 

- si pas de json metier summary lancer la commande `ListMetierSummaryCmd`
- si vous avez le json metier summary mais pas le json metier detail lancer la commande `ListMetierDetailCmd`
- si vous avez le json metier detail lancer la commande `ImportDataListMetierDetailCmd`

## **Doc simple pour chaque clé du JSON RNCP**

| Clé                              | Explication                                                                       | Utilité dans Tenjin                                                                               |
| -------------------------------- | --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `id_fiche`                       | ID interne de la fiche dans la base France Compétences                            | Identifiant technique pour faire des liens ou des requêtes internes                               |
| `numero_fiche`                   | Identifiant officiel RNCP (ex: RNCP16749)                                         | Référence unique à utiliser partout (front, api, admin, etc.)                                     |
| `intitule`                       | Intitulé complet de la certification/diplôme                                      | Affichage principal du diplôme/formation                                                          |
| `abrege`                         | Objet avec code et libellé court (ex: MASTER / Master)                            | Pour affichage rapide ou icône/tag “type de diplôme”                                              |
| `etat_fiche`                     | Statut de la fiche (Publiée, etc.)                                                | Pour filtrer les fiches actives/publiées                                                          |
| `nomenclature_europe`            | Niveau européen (EQF) avec code et libellé (ex: NIV7, Niveau 7)                   | Affiche le niveau sur l’échelle européenne (Bac+3, Bac+5, etc.)                                   |
| `codes_nsf`                      | Liste des codes et libellés **NSF** (Nomenclature Spécialisée France)             | Catégorisation par domaine d’étude ou secteur, pour filtrer ou croiser avec d’autres référentiels |
| `certificateurs`                 | Liste d’organismes certificateurs (nom, statut, site web)                         | Affichage du responsable officiel, recherche par établissement                                    |
| `existence_partenaires`          | Booléen, indique s’il y a des partenaires liés à la formation                     | (Rarement utilisé en front, plutôt pour gestion administrative)                                   |
| `activites_visees`               | Description des missions ou activités principales visées par la certification     | Sert à présenter à quoi forme réellement la formation                                             |
| `capacites_attestees`            | Liste ou description des compétences/aptitudes validées par la formation          | Permet d’expliquer ce que l’apprenant sera capable de faire à la sortie                           |
| `secteurs_activite`              | Secteurs où le diplômé peut exercer (banque, assurance, industrie, etc.)          | Pour orienter ou filtrer selon secteurs pro, ou informer l’utilisateur                            |
| `type_emploi_accessibles`        | Liste d’emplois/métiers accessibles                                               | Affichage des débouchés concrets du diplôme                                                       |
| `remarque_type_emploi`           | Info complémentaire sur les emplois (ex : modalités spécifiques, titres réservés) | (Affichage en encart ou info-bulle)                                                               |
| `codes_rome`                     | Liste des métiers ROME visés (code + libellé)                                     | **Lien direct avec tes métiers Tenjin !** Pour faire le mapping métier → formation                |
| `reglementations_activites`      | Infos réglementaires, si certains métiers sont réglementés                        | À afficher en cas de spécificités légales/professionnelles                                        |
| `jurys`                          | Détail des jurys pour chaque modalité (FI, CA, FC, CQ, CL, VAE)                   | (Plutôt utile en gestion ou pour justifier l’accréditation/équivalence)                           |
| `accessible_nouvelle_caledonie`  | Booléen, formation accessible ou non dans ce territoire                           | Filtrer par accessibilité géographique                                                            |
| `accessible_polynesie_francaise` | Idem, pour la Polynésie                                                           | Idem                                                                                              |
| `publication_decret_general`     | Liste des décrets généraux officiels (titre/numéro)                               | Pour preuve réglementaire, affichage légal, etc.                                                  |
| `publication_decret_creation`    | Décret de création du titre ou de l’habilitation                                  | Idem                                                                                              |
| `publication_decret_autre`       | Autres décrets associés                                                           | Idem                                                                                              |
| `type_enregistrement`            | Nature de l’enregistrement au RNCP (ex: “Enregistrement de droit”)                | Pour affichage administratif ou gestion                                                           |
| `objectifs_contexte`             | Objectifs ou contexte de la certification (quand le champ est renseigné)          | À afficher pour enrichir la compréhension du diplôme                                              |
| `actif`                          | Statut actif/inactif dans la base                                                 | Pour filtrer (ne jamais afficher ce qui n’est pas actif)                                          |
| `prerequis_entree_formation`     | Conditions d’accès/prérequis pour s’inscrire à la formation                       | Indispensable pour affichage sur fiche, test d’éligibilité, etc.                                  |

---

## **Flow d’utilisation RNCP / métier dans Tenjin**

1. **L’utilisateur consulte un métier (fiche métier ROME, par exemple Actuaire)**

    * → Il voit les informations métier, débouchés, etc.

2. **Tenjin interroge la base RNCP** pour toutes les certifications liées à ce code ROME (`codes_rome`)

    * → Affiche la liste des diplômes/certifications permettant d’accéder à ce métier

3. **L’utilisateur clique sur une certification/diplôme**

    * → Il voit la fiche RNCP correspondante (titre, organisme, niveau, capacités, secteurs d’activités, etc.)

4. **Le centre de formation (admin Tenjin) crée une formation**

    * → Il peut rechercher dans le RNCP la certification officielle à rattacher à son offre
    * → Il renseigne les prérequis, le certificateur, le niveau, et relie sa formation à un ou plusieurs codes métier ROME

5. **Bénéfices**

    * Toutes les formations sont rattachées à un référentiel officiel : résultats pros, traçabilité, labels, et interopérabilité
    * L’utilisateur final (apprenant) est guidé de manière fiable du métier souhaité → formations certifiantes adaptées → organisme qui propose

---

### **Schéma résumé**

```
Métier (ROME) ⇄ RNCP (certification officielle) ⇄ Formation créée par centre Tenjin
       ↑                        ↑
    (emploi visé)           (mapping officiel, niveau, certificateur…)
```

---

**Ce système permet :**

* Recherche **par métier** ou **par formation**
* **Mapping automatique** via les codes ROME entre métiers et certifications
* Affichage de tous les critères essentiels (niveau, débouchés, certificateur, prérequis, etc.)
* Valorisation de l’offre centre via rattachement à un référentiel reconnu

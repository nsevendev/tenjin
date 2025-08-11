## Explications détaillées des propriétés principales des fiches métiers

### 1. `riasecMajeur` et `riasecMineur`

* **Définition :**
  Ce sont les codes issus du modèle **RIASEC** (ou modèle de Holland), qui catégorise les métiers selon 6 types de profils :

    * **R** (Réaliste),
    * **I** (Investigateur),
    * **A** (Artistique),
    * **S** (Social),
    * **E** (Entreprenant),
    * **C** (Conventionnel).

* **Utilité :**

    * Permet de **recommander des métiers** adaptés à la personnalité de l’utilisateur, via des tests d’orientation.
    * `riasecMajeur` = dominante du métier ; `riasecMineur` = secondaire.
    * Exemple : R (Réaliste) pour un métier manuel, C (Conventionnel) pour la rigueur administrative.

---

### 2. `transitionEcologique` et `transitionEcologiqueDetaillee`

* **Définition :**

    * `transitionEcologique` : booléen, indique si le métier est concerné par la transition écologique (oui/non).
    * `transitionEcologiqueDetaillee` : code ou libellé décrivant le **type d’impact** du métier sur la transition écologique (ex : “EMPLOI\_BLANC” = pas ou peu impacté, “EMPLOI\_VERT”, “EMPLOI\_VERT\_CROISSANT”, etc.).

* **Utilité :**

    * Permet de cibler les métiers impliqués dans la lutte contre le changement climatique ou qui évoluent à cause de ces enjeux (pour orientation, reconversion, politique RH, etc.).
    * Peut aider à mettre en avant les métiers “verts” ou “verdissants”.

---

### 3. `transitionNumerique`

* **Définition :**

    * Booléen indiquant si le métier subit des changements notables dus à la transition numérique (digitalisation).

* **Utilité :**

    * Identifier les métiers en mutation numérique (intéressant pour la formation, l’anticipation des besoins en compétences numériques, etc.).
    * Peut filtrer les métiers “en tension numérique”.

---

### 4. `codeIsco`

* **Définition :**

    * Code ISCO : correspond au **système de classification international des professions** (utilisé par l’OIT).
    * Exemple : “7515” correspond à un métier précis (ex : Boulanger).

* **Utilité :**

    * Facilite la correspondance internationale des métiers.
    * Peut servir pour l’analyse de marché, les échanges de données, la mobilité internationale, etc.

---

### 5. `domaineProfessionnel`

* **Définition :**

    * Objet contenant un code et un libellé pour le domaine professionnel (ex : “Production”).
    * Peut contenir un “grandDomaine” pour la famille de métiers plus large (ex : “Agriculture et Pêche…”).

* **Utilité :**

    * Sert à regrouper les métiers par domaines.
    * Pratique pour le **filtrage**, la **navigation** dans les fiches métiers, ou la construction d’outils d’orientation.

---

### 6. `appellations`

* **Définition :**

    * Liste des **différents intitulés** (appellations) pour ce métier (synonymes, variantes, titres exacts, etc.).
    * Peut contenir des sous-clés, notamment des compétences liées à chaque appellation (`competencesCles`).

* **Utilité :**

    * Permet d’enrichir la recherche (un métier peut être connu sous différents noms).
    * Important pour l’indexation, les outils de recherche, et pour présenter des synonymes à l’utilisateur.

---

### 7. `competencesMobilisees`, `competencesMobiliseesPrincipales`, `competencesMobiliseesEmergentes`

* **Définition :**

    * **`competencesMobiliseesPrincipales`** : les compétences clés absolument nécessaires pour exercer ce métier (le cœur du métier).
    * **`competencesMobiliseesEmergentes`** : compétences récentes ou en train de devenir importantes (souvent en lien avec la transition numérique, écologique, nouveaux outils, etc.).
    * **`competencesMobilisees`** : toutes les compétences associées (parfois un mélange des deux ci-dessus, ou un ensemble plus large).

* **Utilité :**

    * Sert à **présenter** ce qu’il faut savoir faire pour exercer ce métier.
    * Peut guider les choix de formation ou la rédaction de CV.
    * Les “émergentes” sont stratégiques pour l’**avenir du métier**.

---

### 8. `themes`, `centresInterets`, `centresInteretsLies`

* **Définition :**

    * `themes` : grands axes de travail ou sujets associés au métier.
    * `centresInterets` : centres d’intérêt (ce qui peut plaire à la personne qui exerce ce métier).
    * `centresInteretsLies` : centres d’intérêt indirectement liés (peuvent ouvrir vers d’autres métiers similaires).

* **Utilité :**

    * Sert à personnaliser l’orientation (conseiller des métiers selon les centres d’intérêts d’un utilisateur).
    * Peut être affiché comme “ce métier est fait pour ceux qui aiment : travailler en équipe, résoudre des problèmes, etc.”.

---

### 9. `secteursActivites`, `secteursActivitesLies`

* **Définition :**

    * Secteurs d’activité où le métier est exercé (`secteursActivites`).
    * Autres secteurs voisins, où des compétences proches sont mobilisées (`secteursActivitesLies`).

* **Utilité :**

    * Pour élargir la recherche d’emploi ou de formation.
    * Pratique pour expliquer la **polyvalence** ou la mobilité possible du métier.

---

### 10. `divisionsNaf`

* **Définition :**

    * Liste des codes **NAF** (Nomenclature d’Activité Française) correspondant aux secteurs dans lesquels s’exerce ce métier.
    * Par exemple : “A01Z” pour agriculture.

* **Utilité :**

    * Sert au croisement avec les entreprises (trouver quels secteurs recrutent ce métier).
    * Utilisé pour les statistiques, la data analyse, ou la conformité réglementaire.

---

### 11. `formacodes`

* **Définition :**

    * Liste des **codes de formation** (référentiel de formation utilisé par Pôle emploi/CNCP).
    * Permet d’identifier les cursus/formations qui mènent à ce métier.

* **Utilité :**

    * Très utile pour l’**orientation**, la recommandation de formation.
    * Permet de proposer des liens directs entre fiches métier et offres de formation.

---

### 12. `contextesTravail`

* **Définition :**

    * Décrit les **conditions de travail** spécifiques au métier : environnement (extérieur/intérieur), horaires, pénibilité, port de charges, etc.

* **Utilité :**

    * Sert à informer et à “casser les idées reçues”.
    * Peut permettre de faire un **match** avec les attentes/préférences de l’utilisateur.

---

### 13. `transitionDemographique`

* **Définition :**

    * Indique si le métier est concerné par les enjeux démographiques (vieillissement de la population, évolution du marché du travail, etc.).

* **Utilité :**

    * Sert à anticiper les **évolutions du métier** (besoins futurs, opportunités, etc.).
    * Important pour la politique RH, l’orientation, la veille stratégique.


package jobs

// AppellationDTO Structure pour les appellations métier
type AppellationDTO struct {
	Code                          string             `json:"code" bson:"code"`
	Libelle                       string             `json:"libelle" bson:"libelle"`
	LibelleCourt                  string             `json:"libelleCourt" bson:"libelleCourt"`
	TransitionEcologique          bool               `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionEcologiqueDetaillee string             `json:"transitionEcologiqueDetaillee" bson:"transitionEcologiqueDetaillee"`
	Classification                string             `json:"classification" bson:"classification"`
	CompetencesCles               []CompetenceCleDTO `json:"competencesCles" bson:"competencesCles"`
}

// CompetenceCleDTO Structure pour les compétences clés d'une appellation
type CompetenceCleDTO struct {
	Competence CompetenceDTO `json:"competence" bson:"competence"`
	Frequence  int           `json:"frequence" bson:"frequence"`
}

// CompetenceDTO Structure pour les compétences
type CompetenceDTO struct {
	Type         string `json:"type" bson:"type"`
	Code         string `json:"code" bson:"code"`
	Libelle      string `json:"libelle" bson:"libelle"`
	CodeOgr      string `json:"codeOgr" bson:"codeOgr"`
	RiasecMajeur string `json:"riasecMajeur,omitempty" bson:"riasecMajeur,omitempty"`
	RiasecMineur string `json:"riasecMineur,omitempty" bson:"riasecMineur,omitempty"`
}

// CentreInteretDTO Structure pour les centres d'intérêt
type CentreInteretDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// CentreInteretLieDTO Structure pour les centres d'intérêt liés
type CentreInteretLieDTO struct {
	CentreInteret CentreInteretDTO `json:"centreInteret" bson:"centreInteret"`
	Principal     bool             `json:"principal" bson:"principal"`
}

// ContexteTravailDTO Structure pour les contextes de travail
type ContexteTravailDTO struct {
	Code      string `json:"code" bson:"code"`
	Libelle   string `json:"libelle" bson:"libelle"`
	Categorie string `json:"categorie" bson:"categorie"`
}

// DivisionNafDTO Structure pour les divisions NAF
type DivisionNafDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// DomaineProfessionnelDTO Structure pour le domaine professionnel
type DomaineProfessionnelDTO struct {
	Code         string          `json:"code" bson:"code"`
	Libelle      string          `json:"libelle" bson:"libelle"`
	GrandDomaine GrandDomaineDTO `json:"grandDomaine" bson:"grandDomaine"`
}

// GrandDomaineDTO Structure pour les grands domaines
type GrandDomaineDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// FormacodeDTO Structure pour les formacodes
type FormacodeDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// SecteurActiviteDTO Structure pour les secteurs d'activité
type SecteurActiviteDTO struct {
	Code            string              `json:"code" bson:"code"`
	Libelle         string              `json:"libelle" bson:"libelle"`
	SecteurActivite *SecteurActiviteDTO `json:"secteurActivite,omitempty" bson:"secteurActivite,omitempty"`
}

// SecteurActiviteLieDTO Structure pour les secteurs d'activité liés
type SecteurActiviteLieDTO struct {
	SecteurActivite SecteurActiviteDTO `json:"secteurActivite" bson:"secteurActivite"`
	Principal       bool               `json:"principal" bson:"principal"`
}

// ThemeDTO Structure pour les thèmes
type ThemeDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

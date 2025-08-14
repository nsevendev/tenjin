package jobs

// AppellationDTO représente une appellation de métier
type AppellationDTO struct {
	Code                          string             `json:"code"`
	Libelle                       string             `json:"libelle"`
	LibelleCourt                  string             `json:"libelleCourt"`
	TransitionEcologique          bool               `json:"transitionEcologique"`
	TransitionEcologiqueDetaillee string             `json:"transitionEcologiqueDetaillee"`
	Classification                string             `json:"classification"` // PRINCIPALE, SYNONYME
	CompetencesCles               []CompetenceCleDTO `json:"competencesCles"`
}

// CompetenceCleDTO représente une compétence clé avec sa fréquence
type CompetenceCleDTO struct {
	Competence CompetenceDTO `json:"competence"`
	Frequence  int           `json:"frequence"`
}

// CompetenceDTO représente une compétence
type CompetenceDTO struct {
	Type         string `json:"type"` // COMPETENCE-DETAILLEE, MACRO-SAVOIR-FAIRE, SAVOIR, MACRO-SAVOIR-ETRE-PROFESSIONNEL
	Code         string `json:"code"`
	Libelle      string `json:"libelle"`
	CodeOgr      string `json:"codeOgr"`
	RiasecMajeur string `json:"riasecMajeur,omitempty"`
	RiasecMineur string `json:"riasecMineur,omitempty"`
}

// CentreInteretDTO représente un centre d'intérêt
type CentreInteretDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

// CentreInteretLieDTO représente un centre d'intérêt lié avec son importance
type CentreInteretLieDTO struct {
	CentreInteret CentreInteretDTO `json:"centreInteret"`
	Principal     bool             `json:"principal"`
}

// ContexteTravailDTO représente un contexte de travail
type ContexteTravailDTO struct {
	Code      string `json:"code"`
	Libelle   string `json:"libelle"`
	Categorie string `json:"categorie"` // CONDITIONS_TRAVAIL, HORAIRE_ET_DUREE_TRAVAIL, STATUT_EMPLOI
}

// DivisionNafDTO représente une division NAF
type DivisionNafDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

// DomaineProfessionnelDTO représente un domaine professionnel
type DomaineProfessionnelDTO struct {
	Code         string          `json:"code"`
	Libelle      string          `json:"libelle"`
	GrandDomaine GrandDomaineDTO `json:"grandDomaine"`
}

// GrandDomaineDTO représente un grand domaine professionnel
type GrandDomaineDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

// FormacodeDTO représente un code de formation
type FormacodeDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

// SecteurActiviteDTO représente un secteur d'activité
type SecteurActiviteDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

// SecteurActiviteLieDTO représente un secteur d'activité lié avec son importance
type SecteurActiviteLieDTO struct {
	Principal       bool               `json:"principal"`
	SecteurActivite SecteurActiviteDTO `json:"secteurActivite"`
}

// ThemeDTO représente un thème de métier
type ThemeDTO struct {
	Code    string `json:"code"`
	Libelle string `json:"libelle"`
}

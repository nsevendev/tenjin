package jobs

// JobDTO représente un métier complet selon l'API ROME metier
type JobDTO struct {
	ID                               string                  `json:"_id"`
	AccesEmploi                      string                  `json:"accesEmploi"`
	Appellations                     []AppellationDTO        `json:"appellations"`
	CentresInterets                  []CentreInteretDTO      `json:"centresInterets"`
	CentresInteretsLies              []CentreInteretLieDTO   `json:"centresInteretsLies"`
	Code                             string                  `json:"code"`
	CodeIsco                         string                  `json:"codeIsco"`
	CompetencesMobilisees            []CompetenceDTO         `json:"competencesMobilisees"`
	CompetencesMobiliseesEmergentes  []CompetenceDTO         `json:"competencesMobiliseesEmergentes"`
	CompetencesMobiliseesPrincipales []CompetenceDTO         `json:"competencesMobiliseesPrincipales"`
	ContextesTravail                 []ContexteTravailDTO    `json:"contextesTravail"`
	Definition                       string                  `json:"definition"`
	DivisionsNaf                     []DivisionNafDTO        `json:"divisionsNaf"`
	DomaineProfessionnel             DomaineProfessionnelDTO `json:"domaineProfessionnel"`
	EmploiCadre                      bool                    `json:"emploiCadre"`
	EmploiReglemente                 bool                    `json:"emploiReglemente"`
	Formacodes                       []FormacodeDTO          `json:"formacodes"`
	Libelle                          string                  `json:"libelle"`
	Obsolete                         bool                    `json:"obsolete"`
	RiasecMajeur                     string                  `json:"riasecMajeur"`
	RiasecMineur                     string                  `json:"riasecMineur"`
	SecteursActivites                []SecteurActiviteDTO    `json:"secteursActivites"`
	SecteursActivitesLies            []SecteurActiviteLieDTO `json:"secteursActivitesLies"`
	Themes                           []ThemeDTO              `json:"themes"`
	TransitionDemographique          bool                    `json:"transitionDemographique"`
	TransitionEcologique             bool                    `json:"transitionEcologique"`
	TransitionEcologiqueDetaillee    string                  `json:"transitionEcologiqueDetaillee"`
	TransitionNumerique              bool                    `json:"transitionNumerique"`
}

// JobSummaryDTO représente une version courte d'un métier (pour les recherches et listes)
// TODO : peut etre a modifier au besoin
type JobSummaryDTO struct {
	ID                      string `json:"_id"`
	Code                    string `json:"code"`
	Libelle                 string `json:"libelle"`
	Definition              string `json:"definition"`
	RiasecMajeur            string `json:"riasecMajeur"`
	RiasecMineur            string `json:"riasecMineur"`
	TransitionEcologique    bool   `json:"transitionEcologique"`
	TransitionNumerique     bool   `json:"transitionNumerique"`
	TransitionDemographique bool   `json:"transitionDemographique"`
	EmploiCadre             bool   `json:"emploiCadre"`
	EmploiReglemente        bool   `json:"emploiReglemente"`
}

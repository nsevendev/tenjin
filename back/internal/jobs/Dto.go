package jobs

import "go.mongodb.org/mongo-driver/bson/primitive"

// JobDTO Structure complète d'un métier ROME
type JobDTO struct {
	ID                               primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	AccesEmploi                      string                  `json:"accesEmploi" bson:"accesEmploi"`
	Appellations                     []AppellationDTO        `json:"appellations" bson:"appellations"`
	CentresInterets                  []CentreInteretDTO      `json:"centresInterets" bson:"centresInterets"`
	CentresInteretsLies              []CentreInteretLieDTO   `json:"centresInteretsLies" bson:"centresInteretsLies"`
	Code                             string                  `json:"code" bson:"code"`
	CodeIsco                         string                  `json:"codeIsco" bson:"codeIsco"`
	CompetencesMobilisees            []CompetenceDTO         `json:"competencesMobilisees" bson:"competencesMobilisees"`
	CompetencesMobiliseesEmergentes  []CompetenceDTO         `json:"competencesMobiliseesEmergentes" bson:"competencesMobiliseesEmergentes"`
	CompetencesMobiliseesPrincipales []CompetenceDTO         `json:"competencesMobiliseesPrincipales" bson:"competencesMobiliseesPrincipales"`
	ContextesTravail                 []ContexteTravailDTO    `json:"contextesTravail" bson:"contextesTravail"`
	Definition                       string                  `json:"definition" bson:"definition"`
	DivisionsNaf                     []DivisionNafDTO        `json:"divisionsNaf" bson:"divisionsNaf"`
	DomaineProfessionnel             DomaineProfessionnelDTO `json:"domaineProfessionnel" bson:"domaineProfessionnel"`
	EmploiCadre                      bool                    `json:"emploiCadre" bson:"emploiCadre"`
	EmploiReglemente                 bool                    `json:"emploiReglemente" bson:"emploiReglemente"`
	Formacodes                       []FormacodeDTO          `json:"formacodes" bson:"formacodes"`
	Libelle                          string                  `json:"libelle" bson:"libelle"`
	Obsolete                         bool                    `json:"obsolete" bson:"obsolete"`
	RiasecMajeur                     string                  `json:"riasecMajeur" bson:"riasecMajeur"`
	RiasecMineur                     string                  `json:"riasecMineur" bson:"riasecMineur"`
	SecteursActivites                []SecteurActiviteDTO    `json:"secteursActivites" bson:"secteursActivites"`
	SecteursActivitesLies            []SecteurActiviteLieDTO `json:"secteursActivitesLies" bson:"secteursActivitesLies"`
	Themes                           []ThemeDTO              `json:"themes" bson:"themes"`
	TransitionDemographique          bool                    `json:"transitionDemographique" bson:"transitionDemographique"`
	TransitionEcologique             bool                    `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionEcologiqueDetaillee    string                  `json:"transitionEcologiqueDetaillee" bson:"transitionEcologiqueDetaillee"`
	TransitionNumerique              bool                    `json:"transitionNumerique" bson:"transitionNumerique"`
}

// JobSummaryDTO Version résumée pour les listes
type JobSummaryDTO struct {
	ID                      primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Code                    string                  `json:"code" bson:"code"`
	Libelle                 string                  `json:"libelle" bson:"libelle"`
	Definition              string                  `json:"definition" bson:"definition"`
	RiasecMajeur            string                  `json:"riasecMajeur" bson:"riasecMajeur"`
	RiasecMineur            string                  `json:"riasecMineur" bson:"riasecMineur"`
	TransitionEcologique    bool                    `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionNumerique     bool                    `json:"transitionNumerique" bson:"transitionNumerique"`
	TransitionDemographique bool                    `json:"transitionDemographique" bson:"transitionDemographique"`
	EmploiCadre             bool                    `json:"emploiCadre" bson:"emploiCadre"`
	EmploiReglemente        bool                    `json:"emploiReglemente" bson:"emploiReglemente"`
	DomaineProfessionnel    DomaineProfessionnelDTO `json:"domaineProfessionnel" bson:"domaineProfessionnel"`
}

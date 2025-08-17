package skills

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// SkillDTO Structure complète d'une compétence
type SkillDTO struct {
	ID                     primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Type                   string              `json:"type" bson:"type"`
	Libelle                string              `json:"libelle" bson:"libelle"`
	RiasecMineur           string              `json:"riasecMineur" bson:"riasecMineur"`
	RiasecMajeur           string              `json:"riasecMajeur" bson:"riasecMajeur"`
	Obsolete               bool                `json:"obsolete" bson:"obsolete"`
	Code                   string              `json:"code" bson:"code"`
	CodeOgr                string              `json:"codeOgr" bson:"codeOgr"`
	TransitionEcologique   bool                `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionNumerique    bool                `json:"transitionNumerique" bson:"transitionNumerique"`
	Transferable           bool                `json:"transferable" bson:"transferable"`
	QualiteProfessionnelle string              `json:"qualiteProfessionnelle" bson:"qualiteProfessionnelle"`
	SousCategorie          string              `json:"sousCategorie" bson:"sousCategorie"`
	Definition             string              `json:"definition" bson:"definition"`
	CodeArborescence       string              `json:"codeArborescence" bson:"codeArborescence"`
	Maturite               string              `json:"maturite" bson:"maturite"`
	MacroCompetence        *MacroCompetenceDTO `json:"macroCompetence" bson:"macroCompetence"`
	CompetenceEsco         *CompetenceEscoDTO  `json:"competenceEsco" bson:"competenceEsco"`
	Competences            []SkillSimpleDTO    `json:"competences" bson:"competences"`
	Objectif               *ObjectifDTO        `json:"objectif" bson:"objectif"`
	CategorieSavoir        *CategorieSavoirDTO `json:"categorieSavoir" bson:"categorieSavoir"`
	DateFin                *time.Time          `json:"dateFin" bson:"dateFin"`
}

// SkillSummaryDTO Version résumée pour les listes
type SkillSummaryDTO struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type                 string             `json:"type" bson:"type"`
	Code                 string             `json:"code" bson:"code"`
	Libelle              string             `json:"libelle" bson:"libelle"`
	RiasecMajeur         string             `json:"riasecMajeur" bson:"riasecMajeur"`
	RiasecMineur         string             `json:"riasecMineur" bson:"riasecMineur"`
	TransitionEcologique bool               `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionNumerique  bool               `json:"transitionNumerique" bson:"transitionNumerique"`
	Obsolete             bool               `json:"obsolete" bson:"obsolete"`
	SousCategorie        string             `json:"sousCategorie" bson:"sousCategorie"`
	Transferable         bool               `json:"transferable" bson:"transferable"`
}

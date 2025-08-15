package skills

// SkillSimpleDTO Structure simplifiée pour les références
type SkillSimpleDTO struct {
	Type                 string `json:"type" bson:"type"`
	Libelle              string `json:"libelle" bson:"libelle"`
	RiasecMineur         string `json:"riasecMineur" bson:"riasecMineur"`
	RiasecMajeur         string `json:"riasecMajeur" bson:"riasecMajeur"`
	Code                 string `json:"code" bson:"code"`
	TransitionEcologique bool   `json:"transitionEcologique" bson:"transitionEcologique"`
	TransitionNumerique  bool   `json:"transitionNumerique" bson:"transitionNumerique"`
	CodeOgr              string `json:"codeOgr" bson:"codeOgr"`
}

// MacroCompetenceDTO Structure pour les macro-compétences
type MacroCompetenceDTO struct {
	Type                   string           `json:"type" bson:"type"`
	Libelle                string           `json:"libelle" bson:"libelle"`
	Transferable           bool             `json:"transferable" bson:"transferable"`
	QualiteProfessionnelle string           `json:"qualiteProfessionnelle" bson:"qualiteProfessionnelle"`
	SousCategorie          string           `json:"sousCategorie" bson:"sousCategorie"`
	Code                   string           `json:"code" bson:"code"`
	RiasecMineur           string           `json:"riasecMineur" bson:"riasecMineur"`
	RiasecMajeur           string           `json:"riasecMajeur" bson:"riasecMajeur"`
	Definition             string           `json:"definition" bson:"definition"`
	CodeArborescence       string           `json:"codeArborescence" bson:"codeArborescence"`
	CodeOgr                string           `json:"codeOgr" bson:"codeOgr"`
	Maturite               string           `json:"maturite" bson:"maturite"`
	Objectif               *ObjectifDTO     `json:"objectif" bson:"objectif"`
	Competences            []SkillSimpleDTO `json:"competences" bson:"competences"`
}

// CompetenceEscoDTO Structure pour les compétences ESCO
type CompetenceEscoDTO struct {
	Libelle string `json:"libelle" bson:"libelle"`
	Uri     string `json:"uri" bson:"uri"`
}

// ObjectifDTO Structure pour les objectifs
type ObjectifDTO struct {
	Libelle          string    `json:"libelle" bson:"libelle"`
	Code             string    `json:"code" bson:"code"`
	CodeArborescence string    `json:"codeArborescence" bson:"codeArborescence"`
	Enjeu            *EnjeuDTO `json:"enjeu" bson:"enjeu"`
}

// EnjeuDTO Structure pour les enjeux
type EnjeuDTO struct {
	Libelle           string                `json:"libelle" bson:"libelle"`
	Code              string                `json:"code" bson:"code"`
	CodeArborescence  string                `json:"codeArborescence" bson:"codeArborescence"`
	DomaineCompetence *DomaineCompetenceDTO `json:"domaineCompetence" bson:"domaineCompetence"`
}

// DomaineCompetenceDTO Structure pour les domaines de compétence
type DomaineCompetenceDTO struct {
	Libelle          string `json:"libelle" bson:"libelle"`
	Code             string `json:"code" bson:"code"`
	CodeArborescence string `json:"codeArborescence" bson:"codeArborescence"`
}

// CategorieSavoirDTO Structure pour les catégories de savoir
type CategorieSavoirDTO struct {
	Libelle   string        `json:"libelle" bson:"libelle"`
	Code      string        `json:"code" bson:"code"`
	Categorie *CategorieDTO `json:"categorie" bson:"categorie"`
}

// CategorieDTO Structure pour les catégories
type CategorieDTO struct {
	Libelle string `json:"libelle" bson:"libelle"`
	Code    string `json:"code" bson:"code"`
}

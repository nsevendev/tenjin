package rncp

// RNCPAbregeDDTO Structure pour l'abrégé du diplôme
type RNCPAbregeDDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// RNCPNSFDTO Structure pour les codes NSF (Nomenclature des Spécialités de Formation)
type RNCPNSFDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

// RNCPCertificateurDTO Structure pour les organismes certificateurs
type RNCPCertificateurDTO struct {
	Nom          string `json:"nom" bson:"nom"`
	Etat         string `json:"etat" bson:"etat"`
	SiteInternet string `json:"site_internet" bson:"site_internet"`
}

// RNCPCodeROMEDTO Structure pour les codes ROME associés
type RNCPCodeROMEDTO struct {
	Code    string `json:"code" bson:"code"`
	Libelle string `json:"libelle" bson:"libelle"`
}

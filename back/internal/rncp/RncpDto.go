package rncp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// RNCPCertificationDTO Structure complète d'une certification RNCP
type RNCPCertificationDTO struct {
	ID                    primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	IDFiche               string                 `json:"id_fiche" bson:"id_fiche"`
	NumeroFiche           string                 `json:"numero_fiche" bson:"numero_fiche"`
	Intitule              string                 `json:"intitule" bson:"intitule"`
	Abrege                RNCPAbregeDDTO         `json:"abrege" bson:"abrege"`
	EtatFiche             string                 `json:"etat_fiche" bson:"etat_fiche"`
	Niveau                string                 `json:"niveau" bson:"niveau"`
	NiveauLibelle         string                 `json:"niveau_libelle" bson:"niveau_libelle"`
	CodesNSF              []RNCPNSFDTO           `json:"codes_nsf" bson:"codes_nsf"`
	Certificateurs        []RNCPCertificateurDTO `json:"certificateurs" bson:"certificateurs"`
	ActivitesVisees       string                 `json:"activites_visees" bson:"activitesVisees"`
	CapacitesAttestees    string                 `json:"capacites_attestees" bson:"capacitesAttestees"`
	SecteursActivite      string                 `json:"secteurs_activite" bson:"secteursActivite"`
	TypeEmploiAccessibles string                 `json:"type_emploi_accessibles" bson:"typeEmploiAccessibles"`
	CodesROME             []RNCPCodeROMEDTO      `json:"codes_rome" bson:"codes_rome"`
	Actif                 bool                   `json:"actif" bson:"actif"`
	ImportedAt            time.Time              `json:"imported_at" bson:"imported_at"`
	UpdatedAt             time.Time              `json:"updated_at" bson:"updated_at"`
}

// RNCPCertificationSummaryDTO Version résumée pour les listes
type RNCPCertificationSummaryDTO struct {
	ID             primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	NumeroFiche    string                 `json:"numero_fiche" bson:"numero_fiche"`
	Intitule       string                 `json:"intitule" bson:"intitule"`
	Niveau         string                 `json:"niveau" bson:"niveau"`
	NiveauLibelle  string                 `json:"niveau_libelle" bson:"niveau_libelle"`
	EtatFiche      string                 `json:"etat_fiche" bson:"etat_fiche"`
	Actif          bool                   `json:"actif" bson:"actif"`
	CodesROME      []RNCPCodeROMEDTO      `json:"codes_rome" bson:"codes_rome"`
	Abrege         RNCPAbregeDDTO         `json:"abrege" bson:"abrege"`
	Certificateurs []RNCPCertificateurDTO `json:"certificateurs" bson:"certificateurs"`
}

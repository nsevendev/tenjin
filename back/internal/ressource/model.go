package ressource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model  d'une ressource
type Model struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	URL            string             `bson:"url" json:"url" validate:"required,url"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	Size           int64              `bson:"size" json:"size" validate:"required"`          // en octets
	MimeType       string             `bson:"mime_type" json:"mimeType" validate:"required"` // ex: image/png
	Type           string             `bson:"type" json:"type" validate:"required,oneof=pdf video img"`
	AssociatedTo   primitive.ObjectID `bson:"associated_to" json:"associatedTo" validate:"required"`                                        // ID de l'entité liée
	AssociatedType string             `bson:"associated_type" json:"associatedType" validate:"required,oneof=user session institute offer"` // Type de lien
	UploadedBy     primitive.ObjectID `bson:"uploaded_by" json:"uploadedBy" validate:"required"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt" validate:"required"`
	// ajouter la taille
	// ajouter à qui ou quoi il a été associé
	// le type : pdf, video image etc
}

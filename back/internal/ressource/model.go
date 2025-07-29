package ressource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ressource struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	URL            string             `bson:"url" json:"url" validate:"required,url"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	Size           int64              `bson:"size" json:"size" validate:"required"`
	MimeType       string             `bson:"mime_type" json:"mimeType" validate:"required"`
	Type           string             `bson:"type" json:"type" validate:"required"`
	AssociatedTo   primitive.ObjectID `bson:"associated_to" json:"associatedTo" validate:"required"`
	AssociatedType string             `bson:"associated_type" json:"associatedType" validate:"required,oneof=user session institute offer quiz recruiter evaluation competence "`
	UploadedBy     primitive.ObjectID `bson:"uploaded_by" json:"uploadedBy" validate:"required"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt" validate:"required"`
}

package ressources

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
	AssociationID  primitive.ObjectID `bson:"association_id" json:"associationId" validate:"required"`
	AssociatedType string             `bson:"associated_type" json:"associatedType" validate:"required,oneof=user session institute offer quiz recruiter evaluation competence "`
	UploaderID     primitive.ObjectID `bson:"uploader_id" json:"uploaderId" validate:"required"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time         `bson:"updated_at" json:"updatedAt"`
}

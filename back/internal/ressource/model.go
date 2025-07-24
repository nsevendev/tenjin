package ressource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelUser struct {
	ID         primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
	Type       string             `bson:"type" json:"type" validate:"required, oneof=cv cover_letter portfolio identity certificate photo"` // les ressources du profil ou des éléments en liens avec le User (constantes/fields.FieldsProfile)
	URL        string             `bson:"url" json:"url" validate:"required,url"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	UploadedBy primitive.ObjectID `bson:"uploaded_by" json:"uploaded_by" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at" validate:"required"`
}

type ModelTraining struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Type       string             `bson:"type" json:"type" validate:"required,oneof=pdf video slide exercise other"` // les ressources des cours ou d'une session
	URL        string             `bson:"url" json:"url" validate:"required,url"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	UploadedBy primitive.ObjectID `bson:"uploaded_by" json:"uploaded_by" validate:"required"` // formateur ou admin
	SessionID  primitive.ObjectID `bson:"session_id,omitempty" json:"sessionId,omitempty"`    // facultatif, si lié à une session
	CreatedAt  time.Time          `bson:"created_at" json:"created_at" validate:"required"`
}

package ressource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model de Base d'une ressource
type Model struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	URL        string             `bson:"url" json:"url" validate:"required,url"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	UploadedBy primitive.ObjectID `bson:"uploaded_by" json:"uploaded_by" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at" validate:"required"`
}

type ModelUser struct {
	Model `bson:",inline"`

	Type   string             `bson:"type" json:"type" validate:"required, oneof=cv cover_letter portfolio identity certificate photo"` // les ressources du profil ou des éléments en liens avec le User (constantes/fields.FieldsProfile)
	UserID primitive.ObjectID `bson:"user_id" json:"userId" validate:"required"`
}

type ModelTraining struct {
	Model `bson:",inline"`

	Type      string             `bson:"type" json:"type" validate:"required,oneof=pdf video slide exercise other"` // les ressources des cours ou d'une session
	SessionID primitive.ObjectID `bson:"session_id,omitempty" json:"sessionId,omitempty"`
}

type ModelInstitute struct {
	Model `bson:",inline"`

	Type        string             `bson:"type" json:"type" validate:"required,oneof=logo regulation legal marketing guide report"` // les ressources lié à un Organisme
	InstituteID primitive.ObjectID `bson:"institute_id,omitempty" json:"instituteId,omitempty" validate:"required"`
}

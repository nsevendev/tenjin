package ressource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ressource struct {
	ID         primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
	Type       string             `bson:"type" json:"type" validate:"required"`
	URL        string             `bson:"url" json:"url" validate:"required,url"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	UploadedBy primitive.ObjectID `bson:"uploaded_by" json:"uploaded_by" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at" validate:"required"`
}

package institute

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Institute struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	BusinessName  string               `bson:"business_name" json:"business_name" validate:"required"`
	Siret         string               `bson:"siret" json:"siret" validate:"required"`
	Address       string               `bson:"address" json:"address" validate:"required"`
	ZipCode       string               `bson:"zip_code" json:"zip_code" validate:"required"`
	City          string               `bson:"city" json:"city" validate:"required"`
	ContactEmails []string             `bson:"contact_emails" json:"contact_emails"`
	Formations    []primitive.ObjectID `bson:"formations" json:"formations"`
	Users         []primitive.ObjectID `bson:"users" json:"users" validate:"required"`
	CreatedAt     time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at" json:"updated_at"`
}

package company

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	BusinessName  string               `bson:"business_name" json:"business_name" validate:"required"`
	Siret         string               `bson:"siret" json:"siret" validate:"required"`
	Siren         string               `bson:"siren" json:"siren" validate:"required"`
	Sector        string               `bson:"sector" json:"sector" validate:"required"`
	CompType      string               `bson:"comp_type" json:"comp_type" validate:"required"`
	Address       string               `bson:"address" json:"address" validate:"required"`
	ZipCode       string               `bson:"zip_code" json:"zip_code" validate:"required"`
	City          string               `bson:"city" json:"city" validate:"required"`
	ContactEmails []string             `bson:"contact_emails" json:"contact_emails"`
	Formations    []primitive.ObjectID `bson:"formations" json:"formations"`
	Users         []primitive.ObjectID `bson:"users" json:"users" validate:"required"`
	CreatedAt     primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt     primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}

type CompanyRetrieveDto struct {
	Siret string `json:"siret" binding:"required" validate:"required"`
	Siren string `json:"siren" binding:"required" validate:"required"`
}

type CompanyCreateDto struct {
	BusinessName  string               `bson:"business_name" json:"business_name" validate:"required"`
	Siret         string               `bson:"siret" json:"siret" validate:"required"`
	Siren         string               `bson:"siren" json:"siren" validate:"required"`
	Sector        string               `bson:"sector" json:"sector" validate:"required"`
	CompType      string               `bson:"comp_type" json:"comp_type" validate:"required"`
	Address       string               `bson:"address" json:"address" validate:"required"`
	ZipCode       string               `bson:"zip_code" json:"zip_code" validate:"required"`
	City          string               `bson:"city" json:"city" validate:"required"`
	ContactEmails []string             `bson:"contact_emails" json:"contact_emails"`
	Formations    []primitive.ObjectID `bson:"formations" json:"formations"`
	Users         []primitive.ObjectID `bson:"users" json:"users" validate:"required"`
}

func (c *Company) SetTimeStamps() {
	now := primitive.NewDateTimeFromTime(time.Now())
	if c.CreatedAt == 0 {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
}
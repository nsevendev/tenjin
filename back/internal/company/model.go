package company

import (
	"tenjin/back/internal/addresses"
	"tenjin/back/internal/phones"
	"tenjin/back/internal/utils/constantes"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*type Company struct {
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
}*/

type Company struct {
	ID            primitive.ObjectID       `bson:"_id,omitempty" json:"id"`
	BusinessName  string                   `bson:"business_name" json:"businessName" validate:"required,min=2,max=200"`
	Siret         string                   `bson:"siret" json:"siret" validate:"required,len=14,numeric"`
	Addresses     []addresses.Address      `bson:"addresses" json:"addresses" validate:"required,dive"`
	ContactEmails []string                 `bson:"contact_emails" json:"contactEmails" validate:"required,min=1,dive,email"`
	Phones        []phones.Phone           `bson:"phones" json:"phones" validate:"required,min=1,dive"`
	Status        constantes.StatusState   `bson:"status" json:"status" validate:"required,oneof=enable disable suspended"`
	Type          constantes.TypeInstitute `bson:"type" json:"type" validate:"required,oneof=public private association"`
	LogoUrl       *string                  `bson:"logo_url" json:"logoUrl" validate:"omitempty,url"`
	FormationIDs  []primitive.ObjectID     `bson:"formation_ids" json:"formationIds"`
	UserIDs       []primitive.ObjectID     `bson:"user_ids" json:"userIds"`
	CreatedAt     primitive.DateTime       `bson:"created_at" json:"createdAt"`
	UpdatedAt     primitive.DateTime       `bson:"updated_at" json:"updatedAt"`
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
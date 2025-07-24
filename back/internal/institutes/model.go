package institutes

import (
	"tenjin/back/internal/addresses"
	"tenjin/back/internal/phones"
	"tenjin/back/internal/utils/constantes"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Institute représente un organisme de formation
type Institute struct {
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
	CreatedAt     time.Time                `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time                `bson:"updated_at" json:"updatedAt"`
}

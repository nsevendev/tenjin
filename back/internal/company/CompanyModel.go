package company

import (
	"errors"
	"tenjin/back/internal/addresses"
	"tenjin/back/internal/phones"
	"tenjin/back/internal/utils/constantes"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BusinessName string             `bson:"business_name" json:"businessName" validate:"required,min=2,max=200"`
	Siret        string             `bson:"siret" json:"siret" validate:"required,len=14,numeric"`

	// Gestion hiérarchique des filiales
	ParentCompanyID *primitive.ObjectID  `bson:"parent_company_id,omitempty" json:"parentCompanyId"` // null = maison-mère
	ChildrenIDs     []primitive.ObjectID `bson:"children_ids" json:"childrenIds"`                    // filiales

	Addresses     []addresses.Address      `bson:"addresses" json:"addresses" validate:"required,dive"`
	ContactEmails []string                 `bson:"contact_emails" json:"contactEmails" validate:"required,min=1,dive,email"`
	Phones        []phones.Phone           `bson:"phones" json:"phones" validate:"required,min=1,dive"`
	Status        constantes.StatusState   `bson:"status" json:"status" validate:"required,oneof=enable disable suspended"`
	Type          constantes.TypeInstitute `bson:"type" json:"type" validate:"required,oneof=public private association"`
	LogoUrl       *string                  `bson:"logo_url" json:"logoUrl" validate:"omitempty,url"`

	// Relations à la maille de chaque filiale
	FormationIDs []primitive.ObjectID `bson:"formation_ids" json:"formationIds"`
	UserIDs      []primitive.ObjectID `bson:"user_ids" json:"userIds"`

	CreatedAt primitive.DateTime `bson:"created_at" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updatedAt"`
}

func (c *Company) SetTimeStamps() {
	now := primitive.NewDateTimeFromTime(time.Now())
	if c.CreatedAt == 0 {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
}

// IsParentCompany vérifie si c'est une maison-mère
func (c *Company) IsParentCompany() bool {
	return c.ParentCompanyID == nil
}

// IsSubsidiary vérifie si c'est une filiale
func (c *Company) IsSubsidiary() bool {
	return c.ParentCompanyID != nil
}

// ValidateSiretCoherence valide que le SIRET d'une filiale est cohérent avec sa maison-mère
func (c *Company) ValidateSiretCoherence(parentSiret string) error {
	if c.IsParentCompany() {
		return nil // Pas de validation nécessaire pour une maison-mère
	}

	if len(c.Siret) != 14 || len(parentSiret) != 14 {
		return errors.New("SIRET must be 14 characters long")
	}

	// Les 9 premiers chiffres (SIREN) doivent être identiques
	companySiren := c.Siret[:9]
	parentSiren := parentSiret[:9]

	if companySiren != parentSiren {
		return errors.New("subsidiary SIRET must have the same SIREN as parent company")
	}

	return nil
}

// AddSubsidiary ajoute une filiale à la liste des enfants
func (c *Company) AddSubsidiary(subsidiaryID primitive.ObjectID) {
	// Éviter les doublons
	for _, id := range c.ChildrenIDs {
		if id == subsidiaryID {
			return
		}
	}
	c.ChildrenIDs = append(c.ChildrenIDs, subsidiaryID)
}

// RemoveSubsidiary retire une filiale de la liste des enfants
func (c *Company) RemoveSubsidiary(subsidiaryID primitive.ObjectID) {
	for i, id := range c.ChildrenIDs {
		if id == subsidiaryID {
			c.ChildrenIDs = append(c.ChildrenIDs[:i], c.ChildrenIDs[i+1:]...)
			break
		}
	}
}

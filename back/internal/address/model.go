package address

import (
	"tenjin/back/internal/utils/constantes"
)

type Model struct {
	Number  string             `bson:"number" json:"number" validate:"required,numeric"`
	Street  string             `bson:"street" json:"street" validate:"required,min=5,max=200"`
	ZipCode string             `bson:"zip_code" json:"zipCode" validate:"required,len=5,numeric"`
	City    string             `bson:"city" json:"city" validate:"required,min=2,max=100"`
	Country constantes.Country `bson:"country" json:"country" validate:"required,oneof=france"`
}

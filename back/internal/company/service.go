package company

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type companyService struct {
	collection *mongo.Collection
}

type CompanyServiceInterface interface {
}

func NewCompanyService(db *mongo.Database) CompanyServiceInterface {
	return &companyService{
		collection: db.Collection("companys"),
	}
}
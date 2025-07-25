package insee

import (
	"context"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyService struct {
	collection *mongo.Collection
}

type CompanyServiceInterface interface {
	Create(ctx context.Context, companyCreateDto CompanyCreateDto) (*Company, error)
}

func NewCompanyService(db *mongo.Database) CompanyServiceInterface {
	return &companyService{
		collection: db.Collection("company"),
	}
}

func (s *companyService) Create(ctx context.Context, companyCreateDto CompanyCreateDto) (*Company, error) {
	company := &Company{
		Address:   companyCreateDto.Address,
		BusinessName: companyCreateDto.BusinessName,
		City:   companyCreateDto.City,
		Siret: companyCreateDto.Siret,
		ZipCode:   companyCreateDto.ZipCode,
	}

	result, err := s.collection.InsertOne(ctx, company)
	if err != nil {
		logger.Ef("Erreur lors de la création de la company : %v", err)
		return nil, err
	}

	company.ID = result.InsertedID.(primitive.ObjectID)

	return company, nil
}
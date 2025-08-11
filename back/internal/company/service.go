package company

import (
	"context"
	"fmt"

	"tenjin/back/internal/insee"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyService struct {
	collection *mongo.Collection
}

type CompanyServiceInterface interface {
	RetrieveCompanyInfo(ctx context.Context, siret string, siren string) (*insee.CompanyInfo, error)
}

func NewCompanyService(db *mongo.Database) CompanyServiceInterface {
	return &companyService{
		collection: db.Collection("companys"),
	}
}

func (s *companyService) RetrieveCompanyInfo(ctx context.Context, siret string, siren string) (*insee.CompanyInfo, error) {
	if siret == "" {
		return nil, fmt.Errorf("le SIRET est requis")
	}

	if siren == "" {
		return nil, fmt.Errorf("le SIREN est requis")
	}

	companyInfo, err := insee.CheckSiretExists(siret, siren)
	if err != nil {
		return nil, fmt.Errorf("echec lors de la recuperation des donnees INSEE : %w", err)
	}

	if companyInfo == nil {
		return nil, fmt.Errorf("aucune entreprise trouvee pour le SIRET %s", siret)
	}

	return companyInfo, nil
}

func (s *companyService) Create(ctx context.Context, dto CompanyCreateDto) (*Company, error) {
	company := &Company{
		BusinessName:  dto.BusinessName,
		Siret:         dto.Siret,
		Siren:         dto.Siren,
		Sector:        dto.Sector,
		CompType:      dto.CompType,
		Address:       dto.Address,
		ZipCode:       dto.ZipCode,
		City:          dto.City,
		ContactEmails: dto.ContactEmails,
		Formations:    dto.Formations,
		Users:         dto.Users,
	}

	result, err := s.collection.InsertOne(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la creation de l'entreprise : %w", err)
	}

	company.ID = result.InsertedID.(primitive.ObjectID)
	
	return company, nil
}
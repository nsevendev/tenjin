package company

import (
	"context"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"tenjin/back/internal/utils/mongohelpers"

	"tenjin/back/internal/insee"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyService struct {
	collection  *mongo.Collection
	mongoHelper mongohelpers.Helper
}

type CompanyServiceInterface interface {
	RetrieveCompanyInfo(ctx context.Context, siret string, siren string) (*insee.CompanyInfo, error)
	Create(ctx context.Context, dto CompanyCreateDto) (*Company, error)
}

func NewCompanyService(db *mongo.Database, helper mongohelpers.Helper) CompanyServiceInterface {
	return &companyService{
		collection:  db.Collection("company"),
		mongoHelper: helper,
	}
}

func (s *companyService) RetrieveCompanyInfo(ctx context.Context, siret string, siren string) (*insee.CompanyInfo, error) {
	if siret == "" {
		logger.Ef("le SIRET est requis")
		return nil, fmt.Errorf("le SIRET est requis")
	}

	if siren == "" {
		logger.Ef("le SIREN est requis")
		return nil, fmt.Errorf("le SIREN est requis")
	}

	companyInfo, err := insee.CheckSiretExists(siret, siren)
	if err != nil {
		logger.Ef("echec lors de la recuperation des donnees INSEE pour le SIRET %s et le SIREN %s : %v", siret, siren, err)
		return nil, fmt.Errorf("echec lors de la recuperation des donnees INSEE : %w", err)
	}

	if companyInfo == nil {
		logger.Wf("aucune entreprise trouvee pour le SIRET %s", siret)
		return nil, fmt.Errorf("aucune entreprise trouvee pour le SIRET %s", siret)
	}

	logger.Sf("entreprise trouvee pour le SIRET %s : %v", siret, companyInfo)

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

	s.mongoHelper.SetTimestamps(company)

	result, err := s.collection.InsertOne(ctx, company)
	if err != nil {
		logger.Ef("erreur lors de la creation de l'entreprise : %v", err)
		return nil, fmt.Errorf("erreur lors de la creation de l'entreprise : %w", err)
	}

	company.ID = result.InsertedID.(primitive.ObjectID)

	return company, nil
}

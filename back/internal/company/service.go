package company

import (
	"context"
	"fmt"

	"tenjin/back/internal/insee"

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

/* func (s *companyService) Create(ctx context.Context, companyCreateDto CompanyCreateDto) (*Company, error) {
	company := &Company{
		BusinessName:   companyCreateDto.BusinessName,
		Siret: companyCreateDto.Siret,
		Siren: companyCreateDto.Siren,
		Sector: companyCreateDto.Sector,
		CompType: companyCreateDto.CompType,
		Address: companyCreateDto.Address,
		ZipCode: companyCreateDto.ZipCode,
		City: companyCreateDto.City,
		ContactEmails: companyCreateDto.ContactEmails,
		Formations: companyCreateDto.Formations,
		Users: companyCreateDto.Users,
	}
} */

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
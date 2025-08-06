package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreateCompanyCollection = mignosql.Migration{
	Name: "20250806120000_create_company_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{
					"business_name", "siret", "siren", "sector",
					"comp_type", "address", "zip_code", "city", "users",
				},
				"properties": bson.M{
					"business_name": bson.M{
						"bsonType":    "string",
						"description": "Raison sociale de l'entreprise",
					},
					"siret": bson.M{
						"bsonType":    "string",
						"description": "Numéro SIRET",
					},
					"siren": bson.M{
						"bsonType":    "string",
						"description": "Numéro SIREN",
					},
					"sector": bson.M{
						"bsonType":    "string",
						"description": "Secteur d'activité",
					},
					"comp_type": bson.M{
						"bsonType":    "string",
						"description": "Type d'entreprise",
					},
					"address": bson.M{
						"bsonType":    "string",
						"description": "Adresse complète",
					},
					"zip_code": bson.M{
						"bsonType":    "string",
						"description": "Code postal",
					},
					"city": bson.M{
						"bsonType":    "string",
						"description": "Ville",
					},
					"contact_emails": bson.M{
						"bsonType": "array",
						"description": "Emails de contact",
						"items": bson.M{
							"bsonType": "string",
						},
					},
					"formations": bson.M{
						"bsonType": "array",
						"description": "Formations associées",
						"items": bson.M{
							"bsonType": "objectId",
						},
					},
					"users": bson.M{
						"bsonType": "array",
						"description": "Utilisateurs associés",
						"items": bson.M{
							"bsonType": "objectId",
						},
					},
					"created_at": bson.M{
						"bsonType":    "date",
						"description": "Date de création",
					},
					"updated_at": bson.M{
						"bsonType":    "date",
						"description": "Date de mise à jour",
					},
				},
			},
		}

		opts := options.CreateCollection().SetValidator(validator)
		err := db.CreateCollection(ctx, "company", opts)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return err
			}
		}

		// Index unique sur le SIRET
		indexModel := mongo.IndexModel{
			Keys: bson.M{"siret": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("idx_unique_siret"),
		}

		_, err = db.Collection("company").Indexes().CreateOne(ctx, indexModel)
		return err
	},
}

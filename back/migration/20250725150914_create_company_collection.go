package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreateCompanyCollection = mignosql.Migration{
	Name: "20250725120000_create_company_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"business_name", "siret", "address", "zip_code", "city", "created_at", "updated_at"},
				"properties": bson.M{
					"business_name": bson.M{
						"bsonType":    "string",
						"minLength":   1,
						"description": "Nom de l'entreprise",
					},
					"siret": bson.M{
						"bsonType":    "string",
						"minLength":   14,
						"maxLength":   14,
						"description": "Numéro SIRET, 14 caractères",
					},
					"address": bson.M{
						"bsonType":    "string",
						"minLength":   1,
						"description": "Adresse de l'entreprise",
					},
					"zip_code": bson.M{
						"bsonType":    "string",
						"minLength":   5,
						"maxLength":   5,
						"description": "Code postal",
					},
					"city": bson.M{
						"bsonType":    "string",
						"minLength":   1,
						"description": "Ville",
					},
					"created_at": bson.M{
						"bsonType":    "date",
						"description": "Date de création du document",
					},
					"updated_at": bson.M{
						"bsonType":    "date",
						"description": "Date de mise à jour du document",
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

		indexModel := mongo.IndexModel{
			Keys:    bson.M{"siret": 1},
			Options: options.Index().SetUnique(true).SetName("idx_unique_siret"),
		}

		_, err = db.Collection("company").Indexes().CreateOne(ctx, indexModel)
		return err
	},
}

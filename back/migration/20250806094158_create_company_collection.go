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

		err := db.CreateCollection(ctx, "company", nil)
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

package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreateUserCollection = mignosql.Migration{
	Name: "20250826120000_create_user_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		err := db.CreateCollection(ctx, "user", nil)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return err
			}
		}

		emailIndex := mongo.IndexModel{
			Keys: bson.M{"email": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("idx_unique_email"),
		}

		statusIndex := mongo.IndexModel{
			Keys: bson.M{"status": 1},
			Options: options.Index().
				SetName("idx_status"),
		}

		roleIndex := mongo.IndexModel{
			Keys: bson.M{"roles": 1},
			Options: options.Index().
				SetName("idx_roles"),
		}

		_, err = db.Collection("user").Indexes().CreateMany(ctx, []mongo.IndexModel{
			emailIndex,
			statusIndex,
			roleIndex,
		})
		return err
	},
}

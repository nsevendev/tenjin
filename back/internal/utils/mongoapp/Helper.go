package mongoapp

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type Helper interface {
	SetTimestamps(model interface{})
	FindByID(ctx context.Context, coll *mongo.Collection, id string, result interface{}) error
	ExistsByID(ctx context.Context, coll *mongo.Collection, id string) (bool, error)
}

type helper struct{}

func NewHelper() Helper {
	return &helper{}
}

// ==================== Helpers pour les timestamps ==================== //

// SetTimestamps met à jour les champs CreatedAt et UpdatedAt de n'importe quelle struct
// La struct doit avoir les champs CreatedAt et UpdatedAt de type primitive.DateTime
func (h helper) SetTimestamps(model interface{}) {
	now := primitive.NewDateTimeFromTime(time.Now())

	v := reflect.ValueOf(model).Elem()

	if createdAt := v.FieldByName("CreatedAt"); createdAt.IsValid() && createdAt.CanSet() {
		if createdAt.Interface().(primitive.DateTime) == 0 {
			createdAt.Set(reflect.ValueOf(now))
		}
	}

	if updatedAt := v.FieldByName("UpdatedAt"); updatedAt.IsValid() && updatedAt.CanSet() {
		updatedAt.Set(reflect.ValueOf(now))
	}
}

// ==================== Helpers pour les requêtes courantes ==================== //

// FindByID trouve un document par son ID
func (h helper) FindByID(ctx context.Context, coll *mongo.Collection, id string, result interface{}) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return coll.FindOne(ctx, bson.M{"_id": objID}).Decode(result)
}

// ExistsByID vérifie si un document existe par son ID
func (h helper) ExistsByID(ctx context.Context, coll *mongo.Collection, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	count, err := coll.CountDocuments(ctx, bson.M{"_id": objID})
	return count > 0, err
}

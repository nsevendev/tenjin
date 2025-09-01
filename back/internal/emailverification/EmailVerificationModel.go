package emailverification

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailVerification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Token     string             `bson:"token"`
	Expiry    time.Time          `bson:"expiry"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}
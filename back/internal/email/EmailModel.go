package mail

import (
	"tenjin/back/internal/utils/constantes"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mail struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	To        string             `bson:"to" json:"to"`
	Subject   string             `bson:"subject" json:"subject"`
	Body      string             `bson:"body" json:"body"`
	Type      constantes.TypeMail `json:"type" validate:"required,oneof=welcome inscription reset_password"`
	MetaName  string             `bson:"meta_name,omitempty" json:"meta_name,omitempty"`
	S3Path    string             `bson:"s3_path,omitempty" json:"s3_path,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
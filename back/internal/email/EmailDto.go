package mail

import (
	"tenjin/back/internal/utils/constantes"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailCreateDto struct {
	UserID   	primitive.ObjectID 	`json:"userId" validate:"required"`
	To       	string             	`json:"to" validate:"required,email"`
	Subject  	string             	`json:"subject" validate:"required,min=1,max=200"`
	Body     	string             	`json:"body" validate:"required"`
	Type     	constantes.TypeMail `json:"type" validate:"required,oneof=welcome inscription reset_password"`
	MetaName 	*string            	`json:"metaName,omitempty"`
	S3Path   	*string            	`json:"s3Path,omitempty"`
}
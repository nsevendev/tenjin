package mail

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailCreateDto struct {
	UserID   primitive.ObjectID `json:"userId" validate:"required"`
	To       string             `json:"to" validate:"required,email"`
	Subject  string             `json:"subject" validate:"required,min=1,max=200"`
	Body     string             `json:"body" validate:"required"`
	Type     string             `json:"type" validate:"required"`
	MetaName *string            `json:"metaName,omitempty"`
	S3Path   *string            `json:"s3Path,omitempty"`
}
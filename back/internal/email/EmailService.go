package mail

import (
	"errors"
	"fmt"

	"tenjin/back/internal/utils/mongohelpers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailService struct {
	mongoHelper mongohelpers.Helper
	db          *mongo.Database
}

func NewMailService(mongoHelper mongohelpers.Helper, db *mongo.Database) *MailService {
	return &MailService{
		mongoHelper: mongoHelper,
		db:          db,
	}
}

func (s *MailService) Create(ctx *gin.Context, dto MailCreateDto) (*Mail, error) {
	if dto.UserID.IsZero() {
		return nil, errors.New("UserID est obligatoire")
	}
	if dto.To == "" {
		return nil, errors.New("To (adresse email) est obligatoire")
	}

	mail := &Mail{
		UserID:   dto.UserID,
		To:       dto.To,
		Subject:  dto.Subject,
		Body:     dto.Body,
		Type:     dto.Type,
		MetaName: nil,
		S3Path:   nil,
	}

	s.mongoHelper.SetTimestamps(mail)

	if dto.MetaName != nil {
		mail.MetaName = dto.MetaName
	}
	if dto.S3Path != nil {
		mail.S3Path = dto.S3Path
	}

	result, err := s.db.Collection("mails").InsertOne(ctx, mail)
	if err != nil {
		return nil, fmt.Errorf("impossible de creer le mail : %v", err)
	}

	mail.ID = result.InsertedID.(primitive.ObjectID)
	return mail, nil
}

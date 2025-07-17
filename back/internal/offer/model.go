package offer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OfferResponse struct {
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required"`
	SharedFields []string           `bson:"shared_fields,omitempty" json:"shared_fields,omitempty"`
	ReplyDate    *time.Time         `bson:"reply_date,omitempty" json:"reply_date,omitempty"`
}

type Offer struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty" validate:"required"`
	RecruiterID   primitive.ObjectID   `bson:"recruiter_id" json:"recruiter_id" validate:"required"`
	Title         string               `bson:"title" json:"title" validate:"required"`
	Message       *string              `bson:"message,omitempty" json:"message,omitempty"`
	AttachmentURL *string              `bson:"attachment_url,omitempty" json:"attachment_url,omitempty"`
	CandidateIDs  []primitive.ObjectID `bson:"candidate_ids,omitempty" json:"candidate_ids,omitempty"`
	Status        string               `bson:"status" json:"status" validate:"required,oneof=sent viewed accepted declined"`
	QuizID        *primitive.ObjectID  `bson:"quiz_id,omitempty" json:"quiz_id,omitempty"`
	Responses     []OfferResponse      `bson:"responses" json:"responses" validate:"required,dive"`
	CreatedAt     *time.Time           `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

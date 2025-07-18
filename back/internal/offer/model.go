package offer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OfferResponse struct {
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required" binding:"required"`
	Status       string             `bson:"status" json:"status" validate:"required" binding:"required"`
	SharedFields []string           `bson:"shared_fields" json:"shared_fields"` // optionnel, pas de binding ni validate required
	ReplyDate    *time.Time         `bson:"reply_date" json:"reply_date"`       // optionnel, pointeur pour nil possible
}

type Offer struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id" validate:"required" binding:"required"`
	RecruiterID   primitive.ObjectID   `bson:"recruiter_id" json:"recruiter_id" validate:"required" binding:"required"`
	Title         string               `bson:"title" json:"title" validate:"required" binding:"required"`
	Message       string               `bson:"message" json:"message"` // pas obligatoire donc pas de binding/validate required
	AttachmentURL *string              `bson:"attachment_url" json:"attachment_url"`
	CandidateIDs  []primitive.ObjectID `bson:"candidate_ids" json:"candidate_ids"`
	Status        string               `bson:"status" json:"status" validate:"required,oneof=sent viewed accepted declined" binding:"required"`
	QuizID        *primitive.ObjectID  `bson:"quiz_id" json:"quiz_id"`
	Responses     []OfferResponse      `bson:"responses" json:"responses" validate:"required,dive" binding:"required"`
	CreatedAt     *time.Time           `bson:"created_at" json:"created_at"`
}

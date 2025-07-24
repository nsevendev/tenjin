package offer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	RecruiterID   primitive.ObjectID  `bson:"recruiter_id" json:"recruiterId" validate:"required"`
	CompanyID     primitive.ObjectID  `bson:"company_id" json:"companyId" validate:"required"`
	Title         string              `bson:"title" json:"title" validate:"required"`
	Description   string              `bson:"description" json:"description" validate:"required"`
	AttachmentURL *string             `bson:"attachment_url" json:"attachmentUrl" validate:"omitempty,url"` // si présent, doit être une URL valide
	QuizID        *primitive.ObjectID `bson:"quiz_id" json:"quizId"`
	QuizRequired  *bool               `bson:"quiz_required" json:"quizRequired" validate:"required"`
	StartDateJob  time.Time           `bson:"start_date_job" json:"startDateJob" validate:"required"`
	Salary        string              `bson:"salary" json:"salary" validate:"required"`
	ExpiredAt     time.Time           `bson:"expired_at" json:"expiredAt" validate:"required"`
	EndDate       time.Time           `bson:"end_date" json:"endDate" validate:"required,gtfield=StartDateJob"` // EndDate doit être après StartDateJob
	Status        string              `bson:"status" json:"status" validate:"required,oneof=enable expired disable archived"`
	EmploiType    string              `bson:"emploi_type" json:"emploiType" validate:"required,oneof=CDI CDD Alternance Stage Freelance"`
	CreatedAt     time.Time           `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time           `bson:"updated_at" json:"updatedAt"`
}

// OfferSent - Table intermédiaire : représente l'envoi d'une offre à un candidat spécifique
type ModelSent struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OfferID     primitive.ObjectID `bson:"offer_id" json:"offerId" validate:"required"`
	CandidateID primitive.ObjectID `bson:"candidate_id" json:"candidateId" validate:"required"`
	RecruiterID primitive.ObjectID `bson:"recruiter_id" json:"recruiterId" validate:"required"`
	CompanyID   primitive.ObjectID `bson:"company_id" json:"companyId" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required,oneof=sent viewed responded"`
	SentAt      time.Time          `bson:"sent_at" json:"sentAt" validate:"required"`
	ViewedAt    *time.Time         `bson:"viewed_at" json:"viewedAt"`
	RespondedAt *time.Time         `bson:"responded_at" json:"respondedAt"`
	Message     *string            `bson:"message" json:"message"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}

// OfferResponse - Réponse du candidat à une offre spécifique
type ModelResponse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OfferSentID  primitive.ObjectID `bson:"offer_sent_id" json:"offerSentId" validate:"required"`
	OfferID      primitive.ObjectID `bson:"offer_id" json:"offerId" validate:"required"`
	CompanyID    primitive.ObjectID `bson:"company_id" json:"companyId" validate:"required"`
	CandidateID  primitive.ObjectID `bson:"candidate_id" json:"candidateId" validate:"required"`
	RecruiterID  primitive.ObjectID `bson:"recruiter_id" json:"recruiterId" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=accepted declined"`
	SharedFields []string           `bson:"shared_fields" json:"sharedFields" validate:"dive,oneof=email phone cv linkedin github skills experience location identity"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
}

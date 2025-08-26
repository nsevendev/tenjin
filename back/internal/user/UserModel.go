package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                   primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Firstname            string                 `bson:"firstname" json:"firstname" validate:"required,min=2,max=100"`
	Lastname             string                 `bson:"lastname" json:"lastname" validate:"required,min=2,max=100"`
	Email                string                 `bson:"email" json:"email" validate:"required,email"`
	Roles                []string               `bson:"roles" json:"roles" validate:"required,dive,oneof=student trainer manager admin recruiter"`
	Organizations        []primitive.ObjectID   `bson:"organizations" json:"organizations"`
	Sessions             []primitive.ObjectID   `bson:"sessions" json:"sessions"`
	CompetenceRecords    []CompetenceRecord     `bson:"competence_records" json:"competenceRecords"`
	ExternalExperiences  []ExternalExperience   `bson:"external_experiences" json:"externalExperiences"`
	Status               string                 `bson:"status" json:"status" validate:"omitempty,oneof=training employed jobseeker"`
	Availability         []AvailabilityPeriod   `bson:"availability" json:"availability"`
	ReceivedOffers       []primitive.ObjectID   `bson:"received_offers" json:"receivedOffers"`
	PendingShareRequests []ShareRequest         `bson:"pending_share_requests" json:"pendingShareRequests"`
	QuizResults          []QuizResult           `bson:"quiz_results" json:"quizResults"`
	Chats                []primitive.ObjectID   `bson:"chats" json:"chats"`
	CreatedAt            primitive.DateTime     `bson:"created_at" json:"createdAt"`
	UpdatedAt            primitive.DateTime     `bson:"updated_at" json:"updatedAt"`
}

type CompetenceRecord struct {
	CompetenceID primitive.ObjectID `bson:"competence_id" json:"competenceId"`
	History      []CompetenceEvent  `bson:"history" json:"history"`
}

type CompetenceEvent struct {
	Date        primitive.DateTime   `bson:"date" json:"date"`
	Level       string               `bson:"level" json:"level"`
	ValidatedBy primitive.ObjectID   `bson:"validated_by" json:"validatedBy"`
	SessionID   primitive.ObjectID   `bson:"session_id" json:"sessionId"`
	Notes       string               `bson:"notes" json:"notes,omitempty"`
}

type ExternalExperience struct {
	Title       string               `bson:"title" json:"title"`
	Description string               `bson:"description" json:"description"`
	Date        primitive.DateTime   `bson:"date" json:"date"`
	Proofs      []primitive.ObjectID `bson:"proofs" json:"proofs"`
}

type AvailabilityPeriod struct {
	StartDate primitive.DateTime `bson:"start_date" json:"startDate"`
	EndDate   primitive.DateTime `bson:"end_date" json:"endDate"`
	Type      string             `bson:"type" json:"type"`
}

type ShareRequest struct {
	OfferID         primitive.ObjectID `bson:"offer_id" json:"offerId"`
	FieldsRequested []string           `bson:"fields_requested" json:"fieldsRequested"`
	Status          string             `bson:"status" json:"status" validate:"required,oneof=pending accepted rejected"`
}

type QuizResult struct {
	QuizID  primitive.ObjectID `bson:"quiz_id" json:"quizId"`
	Result  string             `bson:"result" json:"result"`
	Details map[string]any     `bson:"details" json:"details"`
}

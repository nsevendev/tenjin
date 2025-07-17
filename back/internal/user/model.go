package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompetenceHistory struct {
	Date        time.Time          `bson:"date" validate:"required"`
	Level       string             `bson:"level" validate:"required"`
	ValidatedBy primitive.ObjectID `bson:"validated_by" validate:"required"`
	SessionID   primitive.ObjectID `bson:"session_id" validate:"required"`
	Notes       string             `bson:"notes" validate:"required"`
}

type CompetenceRecord struct {
	CompetenceID primitive.ObjectID  `bson:"competence_id"`
	History      []CompetenceHistory `bson:"history"`
}

type ExternalExperience struct {
	Title       string               `bson:"title" validate:"required"`
	Description string               `bson:"description" validate:"required"`
	Date        time.Time            `bson:"date" validate:"required"`
	Proofs      []primitive.ObjectID `bson:"proofs" validate:"required"`
}

type Availability struct {
	StartDate time.Time `bson:"start_date"`
	EndDate   time.Time `bson:"end_date"`
	Type      string    `bson:"type" validate:"required"`
}

type PendingShareRequest struct {
	OfferID         primitive.ObjectID `bson:"offer_id" validate:"required"`
	FieldsRequested []string           `bson:"fields_requested" validate:"required"`
	Status          string             `bson:"status" validate:"required"` // pending | accepted | rejected
}

type QuizResult struct {
	QuizID  primitive.ObjectID `bson:"quiz_id" validate:"required"`
	Result  string             `bson:"result" validate:"required"`
	Details interface{}        `bson:"details" validate:"required"` // Can be a map or anything else
}

type User struct {
	ID                   primitive.ObjectID    `bson:"_id,omitempty" validate:"required"`
	Firstname            string                `bson:"firstname" validate:"required"`
	Lastname             string                `bson:"lastname" validate:"required"`
	Email                string                `bson:"email" validate:"required"`
	Roles                []string              `bson:"roles" validate:"required"`
	Organizations        []primitive.ObjectID  `bson:"organizations"`
	Sessions             []primitive.ObjectID  `bson:"sessions"`
	CompetenceRecords    []CompetenceRecord    `bson:"competence_records"`
	ExternalExperiences  []ExternalExperience  `bson:"external_experiences"`
	Status               string                `bson:"status" validate:"required"`
	Availability         []Availability        `bson:"availability" validate:"required"`
	ReceivedOffers       []primitive.ObjectID  `bson:"received_offers"`
	PendingShareRequests []PendingShareRequest `bson:"pending_share_requests"`
	QuizResults          []QuizResult          `bson:"quiz_results"`
	Chats                []primitive.ObjectID  `bson:"chats"`
	CreatedAt            time.Time             `bson:"created_at"`
	UpdatedAt            time.Time             `bson:"updated_at"`
}

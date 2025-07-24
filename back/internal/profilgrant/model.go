package profilgrant

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileGrant struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	CandidateID   primitive.ObjectID  `bson:"candidate_id" validate:"required"`
	RecruiterID   *primitive.ObjectID `bson:"recruiter_id" json:"recruiterID"`
	CompanyID     *primitive.ObjectID `bson:"company_id" json:"companyID"`
	Audience      *string             `bson:"audience" json:"audience" validate:"omitempty,oneof=recruiter institute company public"`
	ScopeType     string              `bson:"scope_type" json:"scopeType" validate:"required,oneof=recruiter institute company audience"`
	GrantedFields []string            `bson:"granted_fields" json:"grantedFields" validate:"required,dive,oneof=email cv"`
	Revoked       bool                `bson:"revoked" json:"revoked" validate:"required"`
	ExpiresAt     *time.Time          `bson:"expires_at" json:"expiresAt"`
	CreatedAt     time.Time           `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time           `bson:"updated_at" json:"updatedAt"`
}

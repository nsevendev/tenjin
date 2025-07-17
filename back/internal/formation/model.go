package formation

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Formation struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title          string               `bson:"title" json:"title" validate:"required"`
	Description    string               `bson:"description" json:"description" validate:"required"`
	InstituteID    primitive.ObjectID   `bson:"institute_id" json:"institute_id" validate:"required"` 
	CourseIDs      []primitive.ObjectID `bson:"course_ids,omitempty" json:"course_ids,omitempty"`
	CompetenceIDs  []primitive.ObjectID `bson:"competence_ids,omitempty" json:"competence_ids,omitempty"`
	ExternalJobRef string               `bson:"external_job_ref,omitempty" json:"external_job_ref,omitempty"`
	SessionIDs     []primitive.ObjectID `bson:"sessions,omitempty" json:"sessions,omitempty"`
	Meta           map[string]interface{} `bson:"meta,omitempty" json:"meta,omitempty"`
	CreatedAt      time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
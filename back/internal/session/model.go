package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	FormationID       primitive.ObjectID   `bson:"formation_id" json:"formation_id" validate:"required"`
	InstituteID       primitive.ObjectID   `bson:"organization_id" json:"organization_id" validate:"required"`
	Title             string               `bson:"title" json:"title" validate:"required"`
	StartDate         time.Time            `bson:"start_date" json:"start_date" validate:"required"`
	EndDate           time.Time            `bson:"end_date" json:"end_date" validate:"required"`
	Users             []primitive.ObjectID `bson:"users,omitempty" json:"users,omitempty"`
	Trainers          []primitive.ObjectID `bson:"trainers,omitempty" json:"trainers,omitempty"`
	CourseIDs         []primitive.ObjectID `bson:"course_ids,omitempty" json:"course_ids,omitempty"`
	Resources         []primitive.ObjectID `bson:"resources,omitempty" json:"resources,omitempty"`
	Evaluations       []primitive.ObjectID `bson:"evaluations,omitempty" json:"evaluations,omitempty"`
	Quizzes           []primitive.ObjectID `bson:"quizzes,omitempty" json:"quizzes,omitempty"`
	Chats             []primitive.ObjectID `bson:"chats,omitempty" json:"chats,omitempty"`
	CalendarID        *primitive.ObjectID  `bson:"calendar_id,omitempty" json:"calendar_id,omitempty"`
	AttendanceSheetID *primitive.ObjectID  `bson:"attendance_sheet_id,omitempty" json:"attendance_sheet_id,omitempty"`
	CreatedAt         time.Time            `bson:"created_at" json:"created_at"`
}

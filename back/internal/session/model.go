package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID                primitive.ObjectID   `bson:"_id" json:"id`
	FormationID       primitive.ObjectID   `bson:"formation_id" json:"formation_id" validate:"required"`
	InstituteID       primitive.ObjectID   `bson:"organization_id" json:"organization_id" validate:"required"`
	Title             string               `bson:"title" json:"title" validate:"required"`
	StartDate         time.Time            `bson:"start_date" json:"start_date" validate:"required"`
	EndDate           time.Time            `bson:"end_date" json:"end_date" validate:"required"`
	Users             []primitive.ObjectID `bson:"users" json:"users"`
	Trainers          []primitive.ObjectID `bson:"trainers" json:"trainers"`
	CourseIDs         []primitive.ObjectID `bson:"course_ids" json:"course_ids"`
	Resources         []primitive.ObjectID `bson:"resources" json:"resources"`
	Evaluations       []primitive.ObjectID `bson:"evaluations" json:"evaluations"`
	Quizzes           []primitive.ObjectID `bson:"quizzes" json:"quizzes"`
	Chats             []primitive.ObjectID `bson:"chats" json:"chats"`
	CalendarID        *primitive.ObjectID  `bson:"calendar_id" json:"calendar_id"`
	AttendanceSheetID *primitive.ObjectID  `bson:"attendance_sheet_id" json:"attendance_sheet_id"`
	CreatedAt         time.Time            `bson:"created_at" json:"created_at"`
}

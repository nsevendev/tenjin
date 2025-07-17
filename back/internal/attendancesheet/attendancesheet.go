package attendancesheet

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceRecord struct {
	UserID primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Date   time.Time          `bson:"date" json:"date" validate:"required"`
	Status string             `bson:"status" json:"status" validate:"required,oneof=present absent late excused"`
	Notes  *string            `bson:"notes,omitempty" json:"notes,omitempty"`
}

type Attendance struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" validate:"required"`
	SessionID primitive.ObjectID `bson:"session_id" json:"session_id" validate:"required"`
	Records   []AttendanceRecord `bson:"records" json:"records" validate:"required,dive"`
	CreatedAt *time.Time         `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

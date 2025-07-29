package attendancesheet

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceRecord struct {
	UserID     primitive.ObjectID  `bson:"user_id" json:"user_id" validate:"required"`
	Date       time.Time           `bson:"date" json:"date" validate:"required"`
	Status     string              `bson:"status" json:"status" validate:"required,oneof=present absent late excused"`
	Comment    *string             `bson:"comment" json:"comment" `
	RecordedBy primitive.ObjectID  `bson:"recorded_by" json:"recorded_by" validate:"required"`
	UpdatedAt  *time.Time          `bson:"updated_at" json:"updated_at"`
	UpdatedBy  *primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt  time.Time           `bson:"created_at" json:"created_at" validate:"required"`
	CreatedBy  primitive.ObjectID  `bson:"created_by" json:"created_by" validate:"required"`
}

type Attendance struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id" validate:"required"`
	SessionID primitive.ObjectID  `bson:"session_id" json:"session_id" validate:"required"`
	Records   []AttendanceRecord  `bson:"records" json:"records" validate:"required,dive"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at" validate:"required"`
	CreatedBy primitive.ObjectID  `bson:"created_by" json:"created_by" validate:"required"`
	UpdatedAt *time.Time          `bson:"updated_at" json:"updated_at"`
	UpdatedBy *primitive.ObjectID `bson:"updated_by" json:"updated_by"`
}

type AttendanceJustificatif struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	SessionID    primitive.ObjectID `bson:"session_id" json:"session_id" validate:"required"`
	FileURL      string             `bson:"file_url" json:"file_url" validate:"required,url"`
	Date         time.Time          `bson:"date" json:"date" validate:"required"`
	MediaType    string             `bson:"media_type" json:"media_type" validate:"required"`
	DocumentType string             `bson:"document_type" json:"document_type" validate:"required,oneof=justificatif autorisation"`
	Reason       *string            `bson:"reason,omitempty" json:"reason,omitempty"`
	UploadedAt   time.Time          `bson:"uploaded_at" json:"uploaded_at"  validate:"required"`
	UploadedBy   primitive.ObjectID `bson:"uploaded_by" json:"uploaded_by"  validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=pending approved rejected"`
	Comment      *string            `bson:"comment,omitempty" json:"comment,omitempty"`
}

// Constants for attendance status dans STATUS

// Type StatusAttendanceRecord string
// const (
// Present StatusAttendanceRecord = "present"
// Absent  StatusAttendanceRecord = "absent"
// Late    StatusAttendanceRecord = "late"
// Excused StatusAttendanceRecord = "excused"
// )

// Type JustificatifStatus string
// const (
// 	Pending  JustificatifStatus = "pending"
// 	Approved JustificatifStatus = "approved"
// 	Rejected JustificatifStatus = "rejected"
// )

// Type DocumentType string
// const (
// 	Justificatif DocumentType = "justificatif"
// 	Autorisation  DocumentType = "autorisation"
// )

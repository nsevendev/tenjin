package attendancesheet

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceRecord struct {
	UserID      primitive.ObjectID  `bson:"user_id" json:"user_id" validate:"required"`
	Date        time.Time           `bson:"date" json:"date" validate:"required"`
	Status      string              `bson:"status" json:"status" validate:"required,oneof=present absent late excused"`
	Comment     *string             `bson:"comment" json:"comment"`
	RegistrarID *primitive.ObjectID `bson:"registrar_id" json:"registrarId"`
	CreatorID   primitive.ObjectID  `bson:"creator_id" json:"creatorId" validate:"required"`
	UpdaterID   *primitive.ObjectID `bson:"updater_id" json:"updaterId"`
	CreatedAt   time.Time           `bson:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time          `bson:"updated_at" json:"updatedAt"`
}

type Attendance struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	SessionID primitive.ObjectID  `bson:"session_id" json:"sessionId" validate:"required"`
	Records   []AttendanceRecord  `bson:"records" json:"records" validate:"required,dive"`
	CreatorID primitive.ObjectID  `bson:"creator_id" json:"creatorId" validate:"required"`
	UpdaterID *primitive.ObjectID `bson:"updater_id" json:"updaterId"`
	CreatedAt time.Time           `bson:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt *time.Time          `bson:"updated_at" json:"updatedAt"`
}

type AttendanceJustificatif struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"userId" validate:"required"`
	SessionID    primitive.ObjectID `bson:"session_id" json:"sessionId" validate:"required"`
	FileURL      string             `bson:"file_url" json:"fileUrl" validate:"required,url"`
	Date         time.Time          `bson:"date" json:"date" validate:"required"`
	MediaType    string             `bson:"media_type" json:"mediaType" validate:"required"`
	DocumentType string             `bson:"document_type" json:"documentType" validate:"required,oneof=justificatif autorisation"`
	Reason       *string            `bson:"reason" json:"reason"`
	UploadedAt   time.Time          `bson:"uploaded_at" json:"uploadedAt"  validate:"required"`
	UploaderID   primitive.ObjectID `bson:"uploader_id" json:"uploaderId" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=pending approved rejected"`
	Comment      *string            `bson:"comment" json:"comment"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt    *time.Time         `bson:"updated_at" json:"updatedAt"`
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

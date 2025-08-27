package formations

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Formation représente une formation exercée par un organisme de formation
type Formation struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title               string               `bson:"title" json:"title" validate:"required,min=2,max=200"`
	Description         string               `bson:"description" json:"description" validate:"required,min=10,max=2000"`
	IsActive            bool                 `bson:"is_active" json:"isActive"`
	Duration            int                  `bson:"duration" json:"duration" validate:"required,min=1"` // en secondes
	MaxParticipant      *int                 `bson:"max_participant" json:"maxParticipant"`
	DocumentUrls        []string             `bson:"document_urls" json:"documentUrls" validate:"omitempty,dive,url"`
	CompanyID           primitive.ObjectID   `bson:"company_id" json:"companyID" validate:"required"`
	FormationSessionIDs []primitive.ObjectID `bson:"formation_session_ids" json:"formationSessionIds"`
	JobID               primitive.ObjectID   `bson:"job_id" json:"jobId" validate:"required"`
	
	// Soft delete avec traçabilité
	IsDeleted       bool                 `bson:"is_deleted" json:"isDeleted"`
	DeletedAt       *time.Time           `bson:"deleted_at,omitempty" json:"deletedAt"`
	DeletedBy       *primitive.ObjectID  `bson:"deleted_by,omitempty" json:"deletedBy"`
	HasParticipants bool                 `bson:"has_participants" json:"hasParticipants"`
	
	CreatedAt       time.Time            `bson:"created_at" json:"createdAt"`
	UpdatedAt       time.Time            `bson:"updated_at" json:"updatedAt"`
}

// FormationSession creer des session sur une base de formation deja defini, session 2023, 2024, etc ...
type FormationSession struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FormationID       primitive.ObjectID   `bson:"formation_id" json:"formationId" validate:"required"`
	Title             string               `bson:"title" json:"title" validate:"required"`
	IsActive          bool                 `bson:"is_active" json:"isActive"` // disable seulement teacher_admin et teacher_user
	StudentIDs        []primitive.ObjectID `bson:"student_ids" json:"studentIds"`
	TeacherIDs        []primitive.ObjectID `bson:"teacher_ids" json:"teacherIds"`
	CreatorID         primitive.ObjectID   `bson:"creator_id" json:"creatorId" validate:"required"` // ID de l'utilisateur qui a créé la session
	CourseIDs         []primitive.ObjectID `bson:"course_ids" json:"courseIds"`
	ChatID            primitive.ObjectID   `bson:"chat_id" json:"chatId"`
	CalendarID        primitive.ObjectID   `bson:"calendar_id" json:"calendarId"`
	AttendanceSheetID primitive.ObjectID   `bson:"attendance_sheet_id" json:"attendanceSheetId"`
	StartDate         time.Time            `bson:"start_date" json:"startDate" validate:"required"`
	EndDate           time.Time            `bson:"end_date" json:"endDate" validate:"required"`
	StageStartDate    *time.Time           `bson:"stage_start_date" json:"stageStartDate"`
	StageEndDate      *time.Time           `bson:"stage_end_date" json:"stageEndDate"`
	CreatedAt         time.Time            `bson:"created_at" json:"createdAt"`
	UpdatedAt         time.Time            `bson:"updated_at" json:"updatedAt"`
}

// SoftDelete marque la formation comme supprimée
func (f *Formation) SoftDelete(deletedBy primitive.ObjectID) {
	now := time.Now()
	f.IsDeleted = true
	f.DeletedAt = &now
	f.DeletedBy = &deletedBy
	f.IsActive = false
}

// IsAccessible vérifie si la formation est accessible (non supprimée et active)
func (f *Formation) IsAccessible() bool {
	return !f.IsDeleted && f.IsActive
}

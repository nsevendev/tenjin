package crm

import (
	"time"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   primitive.ObjectID       `bson:"_id,omitempty" json:"id"`
	Firstname            string                   `bson:"firstname" json:"firstname"`
	Lastname             string                   `bson:"lastname" json:"lastname"`
	Email                string                   `bson:"email" json:"email"`
	Username             string                   `bson:"username" json:"username"`
	Password             string                   `bson:"password" json:"-"`
	Roles                []string                 `bson:"roles" json:"roles"`
	Status               string                   `bson:"status" json:"status"`
	Organizations        []primitive.ObjectID     `bson:"organizations" json:"organizations"`
	Sessions             []primitive.ObjectID     `bson:"sessions" json:"sessions"`
	CompetenceRecords    []CompetenceRecord       `bson:"competence_records" json:"competenceRecords"`
	ExternalExperiences  []ExternalExperience     `bson:"external_experiences" json:"externalExperiences"`
	ReceivedOffers       []primitive.ObjectID     `bson:"received_offers" json:"receivedOffers"`
	PendingShareRequests []ShareRequest           `bson:"pending_share_requests" json:"pendingShareRequests"`
	QuizResults          []QuizResult             `bson:"quiz_results" json:"quizResults"`
	Chats                []primitive.ObjectID     `bson:"chats" json:"chats"`

	//gestion multi-filiales
	CurrentCompanyID     primitive.ObjectID       `bson:"current_company_id" json:"currentCompanyId"`
	CompanyHistory       []CompanyAssignment      `bson:"company_history" json:"companyHistory"`
	ParticipationHistory []FormationParticipation `bson:"participation_history" json:"participationHistory"`

	// pour la vérif de l'email
	EmailVerified        bool                     `bson:"email_verified" json:"emailVerified"`

	CreatedAt            primitive.DateTime       `bson:"created_at" json:"createdAt"`
	UpdatedAt            primitive.DateTime       `bson:"updated_at" json:"updatedAt"`
}

type CompetenceRecord struct {
	CompetenceID primitive.ObjectID `bson:"competence_id" json:"competenceId"`
	History      []CompetenceEvent  `bson:"history" json:"history"`
}

type CompetenceEvent struct {
	Date        primitive.DateTime `bson:"date" json:"date"`
	Level       string             `bson:"level" json:"level"`
	ValidatedBy primitive.ObjectID `bson:"validated_by" json:"validatedBy"`
	SessionID   primitive.ObjectID `bson:"session_id" json:"sessionId"`
	Notes       string             `bson:"notes,omitempty" json:"notes,omitempty"`
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
	Status          string             `bson:"status" json:"status"`
}

type QuizResult struct {
	QuizID  primitive.ObjectID `bson:"quiz_id" json:"quizId"`
	Result  string             `bson:"result" json:"result"`
	Details map[string]any     `bson:"details" json:"details"`
}

// CompanyAssignment représente l'affectation d'un user à une company (historique des transferts)
type CompanyAssignment struct {
	CompanyID     primitive.ObjectID  `bson:"company_id" json:"companyId"`
	CompanyName   string              `bson:"company_name" json:"companyName"`     // Copie pour historique
	Role          string              `bson:"role" json:"role"`
	StartDate     primitive.DateTime  `bson:"start_date" json:"startDate"`
	EndDate       *primitive.DateTime `bson:"end_date,omitempty" json:"endDate"`   // null = affectation actuelle
	TransferredBy *primitive.ObjectID `bson:"transferred_by,omitempty" json:"transferredBy"`
	Reason        string              `bson:"reason,omitempty" json:"reason"`      // Raison du transfert
}

// FormationParticipation représente la participation d'un user à une formation (avec copie des infos essentielles)
type FormationParticipation struct {
	FormationID       primitive.ObjectID  `bson:"formation_id" json:"formationId"`
	FormationTitle    string              `bson:"formation_title" json:"formationTitle"`
	FormationSessionID *primitive.ObjectID `bson:"formation_session_id,omitempty" json:"formationSessionId"`
	CompanyID         primitive.ObjectID  `bson:"company_id" json:"companyId"`
	CompanyName       string              `bson:"company_name" json:"companyName"`
	StartDate         time.Time           `bson:"start_date" json:"startDate"`
	EndDate           *time.Time          `bson:"end_date,omitempty" json:"endDate"`
	Status            string              `bson:"status" json:"status"` // completed, ongoing, dropped, cancelled
	CompletionRate    *float64            `bson:"completion_rate,omitempty" json:"completionRate"` // % de completion
	Notes             string              `bson:"notes,omitempty" json:"notes"`
	CreatedAt         time.Time           `bson:"created_at" json:"createdAt"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Ef("Erreur de hashage du mot de passe: %v", err)
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// TransferToCompany effectue le transfert d'un user vers une nouvelle company
func (u *User) TransferToCompany(newCompanyID primitive.ObjectID, newCompanyName, newRole string, transferredBy primitive.ObjectID, reason string) {
	now := primitive.NewDateTimeFromTime(time.Now())
	
	// Fermer l'affectation actuelle
	if len(u.CompanyHistory) > 0 {
		for i := range u.CompanyHistory {
			if u.CompanyHistory[i].EndDate == nil {
				u.CompanyHistory[i].EndDate = &now
				break
			}
		}
	}
	
	// Créer la nouvelle affectation
	newAssignment := CompanyAssignment{
		CompanyID:     newCompanyID,
		CompanyName:   newCompanyName,
		Role:          newRole,
		StartDate:     now,
		TransferredBy: &transferredBy,
		Reason:        reason,
	}
	
	u.CompanyHistory = append(u.CompanyHistory, newAssignment)
	u.CurrentCompanyID = newCompanyID
}

// AddFormationParticipation ajoute une participation à une formation dans l'historique
func (u *User) AddFormationParticipation(participation FormationParticipation) {
	participation.CreatedAt = time.Now()
	u.ParticipationHistory = append(u.ParticipationHistory, participation)
}

// GetCurrentCompanyAssignment retourne l'affectation actuelle du user
func (u *User) GetCurrentCompanyAssignment() *CompanyAssignment {
	for i := range u.CompanyHistory {
		if u.CompanyHistory[i].EndDate == nil {
			return &u.CompanyHistory[i]
		}
	}
	return nil
}

// GetActiveFormations retourne les formations en cours
func (u *User) GetActiveFormations() []FormationParticipation {
	var active []FormationParticipation
	for _, participation := range u.ParticipationHistory {
		if participation.Status == "ongoing" {
			active = append(active, participation)
		}
	}
	return active
}

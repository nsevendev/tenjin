package crm

import (
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	Role      string             `bson:"role" json:"role"`
	Username  string             `bson:"username" json:"username"`
	
	// Gestion multi-filiales
	CurrentCompanyID     primitive.ObjectID        `bson:"current_company_id" json:"currentCompanyId"`
	CompanyHistory       []CompanyAssignment       `bson:"company_history" json:"companyHistory"`
	ParticipationHistory []FormationParticipation  `bson:"participation_history" json:"participationHistory"`
	
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
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

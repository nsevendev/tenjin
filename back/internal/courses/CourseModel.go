package courses

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tenjin/back/internal/utils/constantes"
	"time"
)

// Course - Représente un cours de formation
type Course struct {
	ID            primitive.ObjectID         `bson:"_id,omitempty" json:"id"`                             // ID MongoDB
	Title         string                     `bson:"title" json:"title" validate:"required"`              // Titre du cours
	Description   *string                    `bson:"description" json:"description"`                      // Description courte
	ContentBlocks []CourseContentBlock       `bson:"content_blocks" json:"contentBlocks" validate:"dive"` // Blocs de contenu (texte, PDF, etc.)
	CompetenceIDs []primitive.ObjectID       `bson:"competence_ids" json:"competenceIds"`                 // Compétences abordées (IDs) en relation seulement avec les competences de la session / formation
	ResourceIDs   []primitive.ObjectID       `bson:"resource_ids" json:"resourceIds"`                     // Autres ressources (IDs) liées au cours d'une maniere générale
	QuizIDs       []primitive.ObjectID       `bson:"quiz_ids" json:"quizIds"`                             // Liste des quizzes liés au cours
	TypeAccess    constantes.TypeAccessLevel `bson:"type_access" json:"typeAccess" validate:"required"`   // Niveau d'accès, permet de restreindre l'accès au cours au besoin du teacher
	SessionID     primitive.ObjectID         `bson:"session_id" json:"sessionId" validate:"required"`     // Session rattachée, si applicable
	AuthorID      primitive.ObjectID         `bson:"author_id" json:"authorId" validate:"required"`       // Utilisateur ayant créé le cours
	CreatedAt     time.Time                  `bson:"created_at" json:"createdAt"`                         // Date création
	UpdatedAt     time.Time                  `bson:"updated_at" json:"updatedAt"`                         // Date maj
}

// CourseContentBlock - Représente un bloc de contenu dans un cours
// exemple :
/*
"content_blocks": [
  { "type": "text", "data": "Bienvenue sur ce cours, voici l'introduction..." },
  { "type": "pdf",  "url": "https://.../cours-01.pdf", "name": "Cours complet PDF" },
  { "type": "image", "url": "https://.../schema.png", "caption": "Schéma explicatif" },
  { "type": "text", "data": "Passons à la partie pratique..." }
]
*/
type CourseContentBlock struct {
	Type        constantes.TypeCourseContentBlock `bson:"type" json:"type" validate:"text file embed quiz link code audio other"`
	Data        *string                           `bson:"data" json:"data"`         // Pour le texte pur ou embed
	Title       *string                           `bson:"title" json:"title"`       // Titre contextuel, facultatif
	SubTitle    *string                           `bson:"subtitle" json:"subtitle"` // Sous-titre contextuel, facultatif
	Note        *string                           `bson:"note" json:"note"`         // Petite note ou légende
	ResourceIDs []primitive.ObjectID              `bson:"resource_ids" json:"resourceIds"`
}

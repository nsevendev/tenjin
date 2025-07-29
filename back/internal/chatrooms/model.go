package chatrooms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ChatHub représente un espace global de discussion (équivalent d’un serveur Discord)
// C’est le "hub" où plusieurs salons (channels) sont regroupés.
type ChatHub struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`                  
	Name         *string              `bson:"name" validate:"required"`      // Nom du hub (ex: "Centre de formation Numerica") - Obligatoire
	Description  *string              `bson:"description" validate:"required"` // Description textuelle - Obligatoire
	InstituteID  primitive.ObjectID   `bson:"institute_id" validate:"required"` // Référence à l’institute - Obligatoire
	ChannelIDs   []primitive.ObjectID `bson:"channel_ids,omitempty"`          // Liste des IDs des salons (channels) rattachés au hub - facultatif
	CreatedAt    time.Time            `bson:"created_at"`                     // Date de création
}

// Types de salon possibles (a voir pour cette partie)
const (
	ChannelTypeInfo      = "info"      // Salon d’info, pas de débat possible
	ChannelTypePrivate   = "private"   // Salon privé, accès limité
	ChannelTypeModerated = "moderated" // Salon modéré, messages contrôlés par des modérateurs
	ChannelTypeRecruiter = "recruiter" // Salon spécial recruteur, avec auto-suppression après 1 mois
)

const (
    StatusActive   = "active"
    StatusArchived = "archived"
    StatusDeleted  = "deleted"
)


// ChatChannel représente un salon de discussion dans un ChatHub
type ChatChannel struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`                
	HubID          	primitive.ObjectID   `bson:"hub_id" validate:"required"`  // ID du hub auquel appartient ce salon
	SessionID       *primitive.ObjectID  `bson:"session_id,omitempty"`          // Optionnel : lien vers une session formation, ou nil sinon
	Name            *string              `bson:"name" validate:"required"`     // Nom du salon (ex : "Général", "Recrutement")
	Type            *string              `bson:"type" validate:"required"`     // Type de salon (info, private, moderated, recruiter) - Obligatoire
	IsPrivate       bool                 `bson:"is_private"`                   // Indique si le salon est privé (accès restreint) - facultatif
	Participants    []primitive.ObjectID `bson:"participants,omitempty"`       // Liste des IDs des utilisateurs présents dans ce salon
	Moderators      []primitive.ObjectID `bson:"moderators,omitempty"`         // Liste des modérateurs, si le salon est modéré
	AllowedUserIDs  []primitive.ObjectID `bson:"allowed_user_ids,omitempty"`   // Liste des utilisateurs autorisés si salon privé (peut être vide sinon)
	MessageIDs      []primitive.ObjectID `bson:"messages,omitempty"`           // Liste des IDs des messages envoyés dans ce salon
	LastMessage 	*LastMessagePreview  `bson:"last_message,omitempty"`	   // le dernier message relié a la struct lastmessagepreview
	CreatedAt       time.Time            `bson:"created_at"`                   // Date de création du salon
	AutoDeleteAt    *time.Time           `bson:"auto_delete_at,omitempty"`     // Date de suppression automatique (ex : salon recruteur 1 mois)
	Status          *string              `bson:"status" validate:"required"`  // Statut du salon : active, archived ou deleted - Obligatoire
}

// ChatMessage représente un message envoyé dans un ChatChannel
type ChatMessage struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	ChannelID       primitive.ObjectID  `bson:"channel_id" validate:"required"`
	SenderID        primitive.ObjectID  `bson:"sender_id" validate:"required"` //ID de la personne qui a envoyé le msg
	Content         string              `bson:"content" validate:"required"` // contenu du message, obligatoire
	CreatedAt       time.Time           `bson:"created_at"`
	UpdatedAt       *time.Time          `bson:"updated_at,omitempty"`       // Date de la dernière modif
	IsEdited        *bool               `bson:"is_edited,omitempty"`        // Message édité ou non (si je veux par exemple affiché "message modifié")
	ParentMessageID *primitive.ObjectID `bson:"parent_message_id,omitempty"`// Pour les réponses à un message
}

//  stocker le dernier message pour éventuellement l'afficher dans la preview du canal 
type LastMessagePreview struct {
	ID        primitive.ObjectID `bson:"id" validate:"required"`          // ID du message - obligatoire
	Content   string             `bson:"content" validate:"required"`     // Contenu textuel du message - obligatoire
	SenderID  primitive.ObjectID `bson:"sender_id" validate:"required"`   // ID de l’expéditeur - obligatoire
	CreatedAt time.Time          `bson:"created_at"`  					 // Date d’envoi 
}

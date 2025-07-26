package chatrooms 

type ChatRoom struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         *string              `bson:"name"`                  // Obligatoire
	Description  *string              `bson:"description"`           // Obligatoire
	InstituteID  primitive.ObjectID   `bson:"institute_id"`          // Obligatoire
	IsPrivate    *bool                `bson:"is_private"`            // true = accès restreint
	ChannelIDs   []primitive.ObjectID `bson:"channel_ids,omitempty"` // Liste des salons liés
	CreatedAt    time.Time            `bson:"created_at"`
}
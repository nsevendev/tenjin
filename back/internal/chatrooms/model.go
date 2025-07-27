package chatrooms 

type ChatRoom struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         *string              `bson:"name"`                  // Obligatoire
	Description  *string              `bson:"description"`           // Obligatoire
	InstituteID  primitive.ObjectID   `bson:"institute_id"`          // Obligatoire
	ChannelIDs   []primitive.ObjectID `bson:"channel_ids,omitempty"` // Liste des salons liés
	CreatedAt    time.Time            `bson:"created_at"`
}


type ChatChannel struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`  // ID du salon
	RoomID          primitive.ObjectID   `bson:"room_id"`  // id de la chatroom associée
	Name            *string              `bson:"name"` 
	IsPrivate       *bool                `bson:"is_private"` // est ce que le salon est privé ??
	AllowedUserIDs  []primitive.ObjectID `bson:"allowed_user_ids,omitempty"` // user autorisés si il l'est 
	CreatedAt       time.Time            `bson:"created_at"`
}

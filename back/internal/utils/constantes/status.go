package constantes

type StatusState string

const (
	StateEnable    StatusState = "enable"
	StateDisable   StatusState = "disable"
	StateSuspended StatusState = "suspended"
	StateExpired   StatusState = "expired"
	StateArchived  StatusState = "archived"
)

type StatusMessage string

const (
	MessageSent      StatusMessage = "sent"
	MessageViewed    StatusMessage = "viewed"
	MessageResponded StatusMessage = "responded"
)

type StatusOfferResponse string

const (
	OfferResponseAccepted StatusOfferResponse = "accepted"
	OfferResponseDeclined StatusOfferResponse = "declined"
)

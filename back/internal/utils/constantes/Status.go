package constantes

type StatusState string

const (
	StatusStateEnable    StatusState = "enable"
	StatusStateDisable   StatusState = "disable"
	StatusStateSuspended StatusState = "suspended"
	StatusStateExpired   StatusState = "expired"
	StatusStateArchived  StatusState = "archived"
)

type StatusMessage string

const (
	StatusMessageSent      StatusMessage = "sent"
	StatusMessageViewed    StatusMessage = "viewed"
	StatusMessageResponded StatusMessage = "responded"
)

type StatusOfferResponse string

const (
	StatusOfferResponseAccepted StatusOfferResponse = "accepted"
	StatusOfferResponseDeclined StatusOfferResponse = "declined"
)

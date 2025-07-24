package constantes

type StatusActivate string

const (
	Enable    StatusActivate = "enable"
	Disable   StatusActivate = "disable"
	Suspended StatusActivate = "suspended"
	Expired   StatusActivate = "expired"
	Archived  StatusActivate = "archived"
)

type StatusMessage string

const (
	Sent      StatusMessage = "sent"
	Viewed    StatusMessage = "viewed"
	Responded StatusMessage = "responded"
)

type StatusOfferResponse string

const (
	Accepted StatusOfferResponse = "accepted"
	Declined StatusOfferResponse = "declined"
)

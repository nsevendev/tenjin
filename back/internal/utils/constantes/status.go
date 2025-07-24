package constantes

type StatusActivate string

const (
	Enable    StatusActivate = "enable"
	Disable   StatusActivate = "disable"
	Suspended StatusActivate = "suspended"
	Expired   StatusActivate = "expired"
	Archived  StatusActivate = "archived"
)

type MessageStatus string

const (
	Sent      MessageStatus = "sent"
	Viewed    MessageStatus = "viewed"
	Responded MessageStatus = "responded"
)

type OfferResponseStatus string

const (
	Accepted OfferResponseStatus = "accepted"
	Declined OfferResponseStatus = "declined"
)

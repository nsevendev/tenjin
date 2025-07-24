package constantes

type StatusActivate string

const (
	Enable    StatusActivate = "enable"
	Disable   StatusActivate = "disable"
	Suspended StatusActivate = "suspended"
)

type OfferStatus string

const (
	Enabled  OfferStatus = "enabled"
	Expired  OfferStatus = "expired"
	Disabled OfferStatus = "disabled"
	Archived OfferStatus = "archived"
)

type OfferSentStatus string

const (
	Sent      OfferSentStatus = "sent"
	Viewed    OfferSentStatus = "viewed"
	Responded OfferSentStatus = "responded"
)

type OfferResponseStatus string

const (
	Accepted OfferResponseStatus = "accepted"
	Declined OfferResponseStatus = "declined"
)

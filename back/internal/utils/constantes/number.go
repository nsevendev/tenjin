package constantes

type PhoneNumberType string

const (
	PhoneMobile PhoneNumberType = "mobile"
	PhoneFixe   PhoneNumberType = "fixe"
	PhoneFax    PhoneNumberType = "fax"
	PhoneOther  PhoneNumberType = "autre"
)

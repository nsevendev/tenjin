package constantes

type PhoneNumberType string

const (
	Mobile     PhoneNumberType = "mobile"
	Fixe       PhoneNumberType = "fixe"
	Fax        PhoneNumberType = "fax"
	PhoneOther PhoneNumberType = "autre"
)

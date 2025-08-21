package constantes

type PhoneNumberType string

const (
	PhoneNumberMobile PhoneNumberType = "mobile"
	PhoneNumberFixe   PhoneNumberType = "fixe"
	PhoneNumberFax    PhoneNumberType = "fax"
	PhoneNumberOther  PhoneNumberType = "autre"
)

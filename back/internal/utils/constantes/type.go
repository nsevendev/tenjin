package constantes

type TypeInstitute string

const (
	Public      TypeInstitute = "public"
	Private     TypeInstitute = "private"
	Association TypeInstitute = "association"
)

type TypeAddress string

const (
	Invoice    TypeAddress = "invoice"
	Shipping   TypeAddress = "shipping"
	HeadOffice TypeAddress = "headOffice"
	Other      TypeAddress = "other"
	Temporary  TypeAddress = "temporaire"
)

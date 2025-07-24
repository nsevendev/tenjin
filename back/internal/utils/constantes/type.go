package constantes

type TypeInstitute string

const (
	Public      TypeInstitute = "public"
	Private     TypeInstitute = "private"
	Association TypeInstitute = "association"
)

type TypeScope string

const (
	ScopeRecruiter TypeScope = "recruiter"
	ScopeInstitute TypeScope = "institute"
	ScopeCompagny  TypeScope = "company"
	ScopeAudience  TypeScope = "audience"
)

type TypeAudience string

const (
	AudienceRecruiter TypeAudience = "recruiter"
	AudienceInstitute TypeAudience = "institute"
	AudienceCompagny  TypeAudience = "company"
	AudiencePublic    TypeAudience = "public"
)

type TypeEmploi string

const (
	Cdi        TypeEmploi = "CDI"
	Cdd        TypeEmploi = "CDD"
	Alternance TypeEmploi = "Alternance"
	Stage      TypeEmploi = "Stage"
	Freelance  TypeEmploi = "Freelance"
)

type TypeAddress string

const (
	Invoice    TypeAddress = "invoice"
	Shipping   TypeAddress = "shipping"
	HeadOffice TypeAddress = "headOffice"
	Other      TypeAddress = "other"
	Temporary  TypeAddress = "temporaire"
)

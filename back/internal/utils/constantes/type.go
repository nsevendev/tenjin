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

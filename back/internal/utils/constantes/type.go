package constantes

type TypeInstitute string

const (
	InstitutePublic      TypeInstitute = "public"
	InstitutePrivate     TypeInstitute = "private"
	InstituteAssociation TypeInstitute = "association"
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
	EmploiCdi        TypeEmploi = "CDI"
	EmploiCdd        TypeEmploi = "CDD"
	EmploiAlternance TypeEmploi = "Alternance"
	EmploiStage      TypeEmploi = "Stage"
	EmploiFreelance  TypeEmploi = "Freelance"
)

type TypeAddress string

const (
	AddressInvoice    TypeAddress = "invoice"
	AddressShipping   TypeAddress = "shipping"
	AddressHeadOffice TypeAddress = "headOffice"
	AddressOther      TypeAddress = "other"
	AddressTemporary  TypeAddress = "temporaire"
)

// TypeAccessLevel - Niveaux d'accès possibles
type TypeAccessLevel string

const (
	AccessPrivate  TypeAccessLevel = "private"
	AccessTeachers TypeAccessLevel = "teachers"
	AccessSession  TypeAccessLevel = "session"
)

type TypeCourseContentBlock string

const (
	ContentBlockText  TypeCourseContentBlock = "text"
	ContentBlockFile  TypeCourseContentBlock = "file"
	ContentBlockEmbed TypeCourseContentBlock = "embed"
	ContentBlockQuiz  TypeCourseContentBlock = "quiz"
	ContentBlockLink  TypeCourseContentBlock = "link"
	ContentBlockCode  TypeCourseContentBlock = "code"
	ContentBlockAudio TypeCourseContentBlock = "audio"
	ContentBlockOther TypeCourseContentBlock = "other"
)

type TypeRessourceAssociated string

const (
	User       TypeRessourceAssociated = "user"
	Session    TypeRessourceAssociated = "session"
	Institute  TypeRessourceAssociated = "institute"
	Offer      TypeRessourceAssociated = "offer"
	Quiz       TypeRessourceAssociated = "quiz"
	Recruiter  TypeRessourceAssociated = "recruiter"
	Evaluation TypeRessourceAssociated = "evaluation"
	Competence TypeRessourceAssociated = "competence"
)

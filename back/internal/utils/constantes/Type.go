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
	TypeAddressInvoice    TypeAddress = "invoice"
	TypeAddressShipping   TypeAddress = "shipping"
	TypeAddressHeadOffice TypeAddress = "headOffice"
	TypeAddressOther      TypeAddress = "other"
	TypeAddressTemporary  TypeAddress = "temporaire"
)

// TypeAccessLevel - Niveaux d'accès possibles
type TypeAccessLevel string

const (
	TypeAccessPrivate  TypeAccessLevel = "private"
	TypeAccessTeachers TypeAccessLevel = "teachers"
	TypeAccessSession  TypeAccessLevel = "session"
)

type TypeCourseContentBlock string

const (
	TypeContentBlockText  TypeCourseContentBlock = "text"
	TypeContentBlockFile  TypeCourseContentBlock = "file"
	TypeContentBlockEmbed TypeCourseContentBlock = "embed"
	TypeContentBlockQuiz  TypeCourseContentBlock = "quiz"
	TypeContentBlockLink  TypeCourseContentBlock = "link"
	TypeContentBlockCode  TypeCourseContentBlock = "code"
	TypeContentBlockAudio TypeCourseContentBlock = "audio"
	TypeContentBlockOther TypeCourseContentBlock = "other"
)

type TypeRessourceAssociated string

const (
	TypeRessourceUser       TypeRessourceAssociated = "user"
	TypeRessourceSession    TypeRessourceAssociated = "session"
	TypeRessourceInstitute  TypeRessourceAssociated = "institute"
	TypeRessourceOffer      TypeRessourceAssociated = "offer"
	TypeRessourceQuiz       TypeRessourceAssociated = "quiz"
	TypeRessourceRecruiter  TypeRessourceAssociated = "recruiter"
	TypeRessourceEvaluation TypeRessourceAssociated = "evaluation"
	TypeRessourceCompetence TypeRessourceAssociated = "competence"
)

type TypeChannel string

const (
	TypeChannelInfo      TypeChannel = "info"      // Salon d’info, pas de débat possible
	TypeChannelPrivate   TypeChannel = "private"   // Salon privé, accès limité
	TypeChannelModerated TypeChannel = "moderated" // Salon modéré, messages contrôlés par des modérateurs
	TypeChannelRecruiter TypeChannel = "recruiter" // Salon spécial recruteur, avec auto-suppression après 1 mois
)

type TypeMail string

const (
	MailWelcome TypeMail = "welcome" 				// mail de bienvenue
	MailRegister   TypeMail = "register"			// mail de validation du compte
	MailResetPassword TypeMail = "reset_password" 	// mail de reset du password
)
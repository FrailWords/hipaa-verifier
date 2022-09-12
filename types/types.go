package types

// PhiField Protected-Health-Information(PHI) fields
type PhiField string

const (
	SSN         PhiField = "SSN"
	GivenName            = "FirstName"
	FamilyName           = "LastName"
	BirthDate            = "DOB"
	PhoneNumber          = "PhoneNumber"
	Address              = "Address"
)

const (
	SsnRegex         = "^(?!0{3})(?!6{3})[0-8]\\d{2}-(?!0{2})\\d{2}-(?!0{4})\\d{4}$"
	DateRegex        = "^\\d{4}-\\d{2}-\\d{2}$"
	PhoneNumberRegex = "^(?:\\(?([0-9]{3})\\)?[-.●]?)?([0-9]{3})[-.●]?([0-9]{4})$"
	SsnMask          = "XXX-XXX-XXXX"
	DateMask         = "XXXX-XX-XX"
	PhoneNumberMask  = "(XXX)-XXX-XXXX"
	Sha512Regex      = "^\\w{128}$"
)

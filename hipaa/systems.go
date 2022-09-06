package hipaa

import "hipaa-verifier/types"

type DataRecordType int64

const (
	FhirRecord DataRecordType = 0
)

type InterServiceRule struct {
	Source           string
	Destination      string
	RecordType       DataRecordType
	DataProtectRules map[types.PhiField]ProtectRule
	DataVerifyRules  map[types.PhiField]VerifyRule
}

var RegisteredServices = []InterServiceRule{
	{
		Source:           "FhirPatientData",
		Destination:      "FhirPatientAnalytics",
		RecordType:       FhirRecord,
		DataProtectRules: ProtectRules,
		DataVerifyRules:  VerifyRules,
	},
}

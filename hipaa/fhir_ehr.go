package hipaa

import (
	"hipaa-verifier/types"
)

type FhirProtectRule struct {
	Field   types.PhiField
	Protect func(record *types.EhrRecord)
	Verify  func(record types.EhrRecord) bool
}

var protectAndVerifyRules = []FhirProtectRule{
	{
		Field: types.BirthDate,
		Protect: func(record *types.EhrRecord) {
			existingValue := record.BirthDate
			protectRule := ProtectRules[types.BirthDate]
			protectFn := ProtectFunctionsMap[protectRule.Fn]
			newValue := protectFn(existingValue, protectRule.Args...)
			record.BirthDate = newValue
		},
		Verify: func(record types.EhrRecord) bool {
			value := types.BirthDate
			verifyRule := VerifyRules[types.BirthDate]
			verifyFn := VerifyFunctionsMap[verifyRule.Fn]
			return !verifyFn(value, verifyRule.Args...)
		},
	},
	{
		Field: types.PhoneNumber,
		Protect: func(record *types.EhrRecord) {
			telecomValues := record.Telecom
			protectRule := ProtectRules[types.PhoneNumber]
			for idx, telecom := range telecomValues {
				if telecom.System == "phone" {
					existingValue := telecomValues[idx].Value
					protectFn := ProtectFunctionsMap[protectRule.Fn]
					telecomValues[idx].Value = protectFn(existingValue, protectRule.Args...)
				}
			}
		},
		Verify: func(record types.EhrRecord) bool {
			telecomValues := record.Telecom
			verifyRule := VerifyRules[types.PhoneNumber]
			for idx, telecom := range telecomValues {
				if telecom.System == "phone" {
					value := telecomValues[idx].Value
					verifyFn := VerifyFunctionsMap[verifyRule.Fn]
					return !verifyFn(value, verifyRule.Args...)
				}
			}
			return false
		},
	},
}

func identifyAndProtect(data *types.EhrRecord) {
	for _, rule := range protectAndVerifyRules {
		rule.Protect(data)
	}
}

func ProtectDataRecords(data []*types.EhrRecord) []*types.EhrRecord {
	for _, record := range data {
		identifyAndProtect(record)
	}
	return data
}

func identifyAndVerify(data types.EhrRecord) bool {
	var allFieldsVerified = false
	for _, rule := range protectAndVerifyRules {
		allFieldsVerified = allFieldsVerified || rule.Verify(data)
	}
	return allFieldsVerified
}

func VerifyDataRecords(data []types.EhrRecord) bool {
	var allFieldsProtectedAndVerified = true
	for _, record := range data {
		allFieldsProtectedAndVerified = allFieldsProtectedAndVerified || identifyAndVerify(record)
	}
	return allFieldsProtectedAndVerified
}

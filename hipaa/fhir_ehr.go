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
			value := record.BirthDate
			verifyRule := VerifyRules[types.BirthDate]
			verifyFn := VerifyFunctionsMap[verifyRule.Fn]
			return verifyFn(value, false, verifyRule.Args...)
		},
	},
	{
		Field: types.GivenName,
		Protect: func(record *types.EhrRecord) {
			existingValue := record.Name[0].Given[0]
			protectRule := ProtectRules[types.GivenName]
			protectFn := ProtectFunctionsMap[protectRule.Fn]
			newValue := protectFn(existingValue, protectRule.Args...)
			record.Name[0].Given[0] = newValue
		},
		Verify: func(record types.EhrRecord) bool {
			value := record.Name[0].Given[0]
			verifyRule := VerifyRules[types.GivenName]
			verifyFn := VerifyFunctionsMap[verifyRule.Fn]
			return verifyFn(value, true, verifyRule.Args...)
		},
	},
	{
		Field: types.FamilyName,
		Protect: func(record *types.EhrRecord) {
			existingValue := record.Name[0].Family
			protectRule := ProtectRules[types.FamilyName]
			protectFn := ProtectFunctionsMap[protectRule.Fn]
			newValue := protectFn(existingValue, protectRule.Args...)
			record.Name[0].Family = newValue
		},
		Verify: func(record types.EhrRecord) bool {
			value := record.Name[0].Family
			verifyRule := VerifyRules[types.FamilyName]
			verifyFn := VerifyFunctionsMap[verifyRule.Fn]
			return verifyFn(value, true, verifyRule.Args...)
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
					return verifyFn(value, false, verifyRule.Args...)
				}
			}
			return true
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
	var allFieldsVerified = true
	for _, rule := range protectAndVerifyRules {
		allFieldsVerified = allFieldsVerified && rule.Verify(data)
	}
	return allFieldsVerified
}

func VerifyDataRecords(data []types.EhrRecord) bool {
	var allFieldsProtectedAndVerified = true
	for _, record := range data {
		allFieldsProtectedAndVerified = allFieldsProtectedAndVerified && identifyAndVerify(record)
	}
	return allFieldsProtectedAndVerified
}

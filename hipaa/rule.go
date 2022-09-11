package hipaa

import (
	"hipaa-verifier/types"
	"hipaa-verifier/utils"
	"regexp"
)

type ProtectType int64

const (
	Mask ProtectType = iota
	HashValue
)

type VerifyType int64

const (
	RegexMatcher VerifyType = iota
)

type ProtectRule struct {
	Type ProtectType
	Fn   ProtectFunctions
	Args []string
}

func MaskingFn(value string, args ...string) string {
	return args[0]
}

func HashFn(value string, _ ...string) string {
	return utils.Sha512HashOfValue(value)
}

type ProtectFunctions string

const (
	MaskingFunction ProtectFunctions = "MaskingFn"
	HashingFunction                  = "HashFn"
)

var ProtectFunctionsMap = map[ProtectFunctions]func(string, ...string) string{
	MaskingFunction: MaskingFn,
	HashingFunction: HashFn,
}

// ProtectRules - definition
var ProtectRules = map[types.PhiField]ProtectRule{
	types.SSN: {
		Type: Mask,
		Fn:   MaskingFunction,
		Args: []string{types.SsnMask},
	},
	types.BirthDate: {
		Type: Mask,
		Fn:   MaskingFunction,
		Args: []string{types.DateMask},
	},
	types.PhoneNumber: {
		Type: Mask,
		Fn:   MaskingFunction,
		Args: []string{types.PhoneNumberMask},
	},
	types.GivenName: {
		Type: HashValue,
		Fn:   HashingFunction,
	},
	types.FamilyName: {
		Type: HashValue,
		Fn:   HashingFunction,
	},
}

type VerifyRule struct {
	FieldName types.PhiField
	Type      VerifyType
	Fn        VerifyFunctions
	Args      []string
}

func MatchRegexFn(value string, args ...string) bool {
	match, _ := regexp.Match(args[0], []byte(value))
	return match
}

type VerifyFunctions string

const (
	MatchRegularExpression VerifyFunctions = "MatchRegexFn"
)

var VerifyFunctionsMap = map[VerifyFunctions]func(string, ...string) bool{
	MatchRegularExpression: MatchRegexFn,
}

// VerifyRules - definition
var VerifyRules = map[types.PhiField]VerifyRule{
	types.SSN: {
		Type: RegexMatcher,
		Fn:   MatchRegularExpression,
		Args: []string{types.SsnRegex},
	},
	types.BirthDate: {
		Type: RegexMatcher,
		Fn:   MatchRegularExpression,
		Args: []string{types.DateRegex},
	},
	types.PhoneNumber: {
		Type: RegexMatcher,
		Fn:   MatchRegularExpression,
		Args: []string{types.PhoneNumberRegex},
	},
}

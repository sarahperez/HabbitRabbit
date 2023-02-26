package password

//Copyright 2020 Joseph Matthias Goh (@zephinzer > @usvc)

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// LowercaseCharacters defines lowercase characters
	LowercaseCharacters = "abcdefghijklmnopqrstuvwxyz"

	// NumericCharacters defines numerical characters
	NumericCharacters = "1234567890"

	// SpecialCharacters defines special characters
	SpecialCharacters = "~`!@#$%^&*()_-=+[{]}\\|;:'\"<>./?"

	// UppercaseCharacters defines uppercase characters
	UppercaseCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// DefaultPasswordCustomSpecial defines the default character set
	// that is used to define special characters
	DefaultPasswordCustomSpecial string = SpecialCharacters

	// DefaultPasswordMinimumLowercaseCount defines the default for
	// the number of lower-cased characters required
	DefaultPasswordMinimumLowercaseCount int = 0

	// DefaultPasswordMinimumUppercaseCount defines the default for
	// the number of upper-cased characters required
	DefaultPasswordMinimumUppercaseCount int = 0

	// DefaultPasswordMaximumLength defines the default for the length
	// of the password
	DefaultPasswordMaximumLength int = 64

	// DefaultPasswordMinimumLength defines the default for the minimum
	// length of the password
	DefaultPasswordMinimumLength int = 8

	// DefaultPasswordMinimumNumericCount defines the default
	// number of numeric characters in the password
	DefaultPasswordMinimumNumericCount int = 0

	// DefaultPasswordMinimumSpecialCount defines the default
	// number of special characters in the password
	DefaultPasswordMinimumSpecialCount int = 0

	// StringTypeNumeric defines a numeric character type
	StringTypeNumeric = "NUMERIC"
	// StringTypeSpecial defines a special character type
	StringTypeSpecial = "SPECIAL"
	// StringTypeLowercase defines a lowercase character type
	StringTypeLowercase = "LOWERCASE"
	// StringTypeUppercase defines a uppercase character type
	StringTypeUppercase = "UPPERCASE"
	// StringTypeUnknown defines an unknown character type
	StringTypeUnknown = "UNKNOWN"
)

// Policy defines possible configurations for password
// requirements
type Policy struct {
	MaximumLength         int
	MinimumLength         int
	MinimumLowercaseCount int
	MinimumUppercaseCount int
	MinimumNumericCount   int
	MinimumSpecialCount   int
	CustomSpecial         []byte
}

// GetDefaultPolicy returns a Policy with
// its values set to the default
func GetDefaultPolicy() Policy {
	return Policy{
		MaximumLength:         DefaultPasswordMaximumLength,
		MinimumLength:         DefaultPasswordMinimumLength,
		MinimumLowercaseCount: DefaultPasswordMinimumLowercaseCount,
		MinimumUppercaseCount: DefaultPasswordMinimumUppercaseCount,
		MinimumNumericCount:   DefaultPasswordMinimumNumericCount,
		MinimumSpecialCount:   DefaultPasswordMinimumSpecialCount,
		CustomSpecial:         []byte(DefaultPasswordCustomSpecial),
	}
}

// StringMetadata provides metadata about a string, use GetStringMetadata to
// generate the metadata
type StringMetadata struct {
	Length     int
	Lowercases strings.Builder
	Uppercases strings.Builder
	Numerics   strings.Builder
	Specials   strings.Builder
	Unknowns   strings.Builder
	PrefixType string
	SuffixType string
}

// GetStringMetadata returns a populated StringMetadata structure
// that provides meta-data about the provided plaintext string for
// further processing by validators
func GetStringMetadata(plaintext string, customSpecial ...[]byte) StringMetadata {
	stringMetadata := StringMetadata{
		Length:     len(plaintext),
		Lowercases: strings.Builder{},
		Uppercases: strings.Builder{},
		Numerics:   strings.Builder{},
		Specials:   strings.Builder{},
		Unknowns:   strings.Builder{},
	}
	for i := 0; i < len(plaintext); i++ {
		currentCharacter := plaintext[i]
		specials := []byte(SpecialCharacters)
		if len(customSpecial) > 0 && len(customSpecial[0]) > 0 {
			specials = customSpecial[0]
		}
		var characterType string
		var builder *strings.Builder
		if bytes.Contains([]byte(specials), []byte{currentCharacter}) {
			characterType = StringTypeSpecial
			builder = &stringMetadata.Specials
		} else if bytes.Contains([]byte(LowercaseCharacters), []byte{currentCharacter}) {
			characterType = StringTypeLowercase
			builder = &stringMetadata.Lowercases
		} else if bytes.Contains([]byte(UppercaseCharacters), []byte{currentCharacter}) {
			characterType = StringTypeUppercase
			builder = &stringMetadata.Uppercases
		} else if bytes.Contains([]byte(NumericCharacters), []byte{currentCharacter}) {
			characterType = StringTypeNumeric
			builder = &stringMetadata.Numerics
		} else {
			characterType = StringTypeUnknown
			builder = &stringMetadata.Unknowns
		}
		if i == 0 {
			stringMetadata.PrefixType = characterType
		} else if i == len(plaintext)-1 {
			stringMetadata.SuffixType = characterType
		}
		builder.WriteByte(currentCharacter)
	}
	return stringMetadata
}

// Validate validates a provided plaintext password using the
// default PasswordPolicy or a custom policy if it's provided
func Validate(plaintext string, customPolicy ...Policy) error {
	policy := GetDefaultPolicy()
	if len(customPolicy) > 0 {
		policy = customPolicy[0]
	}

	passwordMetadata := GetStringMetadata(plaintext, policy.CustomSpecial)
	switch true {
	case passwordMetadata.Length < policy.MinimumLength:
		return fmt.Errorf("provided password requires at least %v characters", policy.MinimumLength)
	case passwordMetadata.Length > policy.MaximumLength:
		return fmt.Errorf("provided password exceeds the maximum length of %v characters", policy.MaximumLength)
	case passwordMetadata.Lowercases.Len() < policy.MinimumLowercaseCount:
		return fmt.Errorf("provided password requires at least %v lower-cased characters", policy.MinimumLowercaseCount)
	case passwordMetadata.Uppercases.Len() < policy.MinimumUppercaseCount:
		return fmt.Errorf("provided password requires at least %v upper-cased characters", policy.MinimumUppercaseCount)
	case passwordMetadata.Numerics.Len() < policy.MinimumNumericCount:
		return fmt.Errorf("provided password requires at least %v numeric characters", policy.MinimumNumericCount)
	case passwordMetadata.Specials.Len() < policy.MinimumSpecialCount:
		return fmt.Errorf("provided password requires at least %v special characters", policy.MinimumSpecialCount)
	}
	return nil
}

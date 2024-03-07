/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api


import (
	"fmt"
)



type Gender string

// List of Gender
const (
	MALE Gender = "MALE"
	FEMALE Gender = "FEMALE"
	OTHER Gender = "OTHER"
)

// AllowedGenderEnumValues is all the allowed values of Gender enum
var AllowedGenderEnumValues = []Gender{
	"MALE",
	"FEMALE",
	"OTHER",
}

// validGenderEnumValue provides a map of Genders for fast verification of use input
var validGenderEnumValues = map[Gender]struct{}{
	"MALE": {},
	"FEMALE": {},
	"OTHER": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Gender) IsValid() bool {
	_, ok := validGenderEnumValues[v]
	return ok
}

// NewGenderFromValue returns a pointer to a valid Gender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGenderFromValue(v string) (Gender, error) {
	ev := Gender(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for Gender: valid values are %v", v, AllowedGenderEnumValues)
	}
}



// AssertGenderRequired checks if the required fields are not zero-ed
func AssertGenderRequired(obj Gender) error {
	return nil
}

// AssertGenderConstraints checks if the values respects the defined constraints
func AssertGenderConstraints(obj Gender) error {
	return nil
}

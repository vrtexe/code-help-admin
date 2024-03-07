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



type SubmissionStatus string

// List of SubmissionStatus
const (
	ACCEPTED SubmissionStatus = "ACCEPTED"
	DECLINED SubmissionStatus = "DECLINED"
)

// AllowedSubmissionStatusEnumValues is all the allowed values of SubmissionStatus enum
var AllowedSubmissionStatusEnumValues = []SubmissionStatus{
	"ACCEPTED",
	"DECLINED",
}

// validSubmissionStatusEnumValue provides a map of SubmissionStatuss for fast verification of use input
var validSubmissionStatusEnumValues = map[SubmissionStatus]struct{}{
	"ACCEPTED": {},
	"DECLINED": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubmissionStatus) IsValid() bool {
	_, ok := validSubmissionStatusEnumValues[v]
	return ok
}

// NewSubmissionStatusFromValue returns a pointer to a valid SubmissionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubmissionStatusFromValue(v string) (SubmissionStatus, error) {
	ev := SubmissionStatus(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for SubmissionStatus: valid values are %v", v, AllowedSubmissionStatusEnumValues)
	}
}



// AssertSubmissionStatusRequired checks if the required fields are not zero-ed
func AssertSubmissionStatusRequired(obj SubmissionStatus) error {
	return nil
}

// AssertSubmissionStatusConstraints checks if the values respects the defined constraints
func AssertSubmissionStatusConstraints(obj SubmissionStatus) error {
	return nil
}

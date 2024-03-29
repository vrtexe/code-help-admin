/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api




type SubmissionEntry struct {

	Language string `json:"language"`

	Code string `json:"code"`

	Id int32 `json:"id,omitempty"`

	TimeSubmitted string `json:"timeSubmitted"`

	Status SubmissionStatus `json:"status"`
}

// AssertSubmissionEntryRequired checks if the required fields are not zero-ed
func AssertSubmissionEntryRequired(obj SubmissionEntry) error {
	elements := map[string]interface{}{
		"language": obj.Language,
		"code": obj.Code,
		"timeSubmitted": obj.TimeSubmitted,
		"status": obj.Status,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSubmissionEntryConstraints checks if the values respects the defined constraints
func AssertSubmissionEntryConstraints(obj SubmissionEntry) error {
	return nil
}

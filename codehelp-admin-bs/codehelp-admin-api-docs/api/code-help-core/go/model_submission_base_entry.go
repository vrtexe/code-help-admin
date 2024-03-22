/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api




type SubmissionBaseEntry struct {

	Language string `json:"language"`

	Code string `json:"code"`
}

// AssertSubmissionBaseEntryRequired checks if the required fields are not zero-ed
func AssertSubmissionBaseEntryRequired(obj SubmissionBaseEntry) error {
	elements := map[string]interface{}{
		"language": obj.Language,
		"code": obj.Code,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSubmissionBaseEntryConstraints checks if the values respects the defined constraints
func AssertSubmissionBaseEntryConstraints(obj SubmissionBaseEntry) error {
	return nil
}
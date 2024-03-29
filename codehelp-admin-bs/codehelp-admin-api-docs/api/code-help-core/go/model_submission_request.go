/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api




type SubmissionRequest struct {

	Language string `json:"language"`

	Code string `json:"code"`

	ProblemId int32 `json:"problemId"`
}

// AssertSubmissionRequestRequired checks if the required fields are not zero-ed
func AssertSubmissionRequestRequired(obj SubmissionRequest) error {
	elements := map[string]interface{}{
		"language": obj.Language,
		"code": obj.Code,
		"problemId": obj.ProblemId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSubmissionRequestConstraints checks if the values respects the defined constraints
func AssertSubmissionRequestConstraints(obj SubmissionRequest) error {
	return nil
}

/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api




type ContestProblem struct {

	Id int32 `json:"id,omitempty"`

	Problem ProblemEntry `json:"problem"`

	Score int32 `json:"score"`
}

// AssertContestProblemRequired checks if the required fields are not zero-ed
func AssertContestProblemRequired(obj ContestProblem) error {
	elements := map[string]interface{}{
		"problem": obj.Problem,
		"score": obj.Score,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertProblemEntryRequired(obj.Problem); err != nil {
		return err
	}
	return nil
}

// AssertContestProblemConstraints checks if the values respects the defined constraints
func AssertContestProblemConstraints(obj ContestProblem) error {
	return nil
}

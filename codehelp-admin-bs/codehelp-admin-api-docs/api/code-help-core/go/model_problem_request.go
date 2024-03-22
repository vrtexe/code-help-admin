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
	"os"
)



type ProblemRequest struct {

	Category Category `json:"category,omitempty"`

	Title string `json:"title"`

	Difficulty Difficulty `json:"difficulty"`

	Markdown string `json:"markdown,omitempty"`

	StarterCode *os.File `json:"starterCode"`

	RunnerCode *os.File `json:"runnerCode"`

	TestCases []*os.File `json:"testCases"`
}

// AssertProblemRequestRequired checks if the required fields are not zero-ed
func AssertProblemRequestRequired(obj ProblemRequest) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"difficulty": obj.Difficulty,
		"starterCode": obj.StarterCode,
		"runnerCode": obj.RunnerCode,
		"testCases": obj.TestCases,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCategoryRequired(obj.Category); err != nil {
		return err
	}
	return nil
}

// AssertProblemRequestConstraints checks if the values respects the defined constraints
func AssertProblemRequestConstraints(obj ProblemRequest) error {
	return nil
}
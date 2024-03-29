/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type ModeratorRequest struct {

	Username string `json:"username"`
}

// AssertModeratorRequestRequired checks if the required fields are not zero-ed
func AssertModeratorRequestRequired(obj ModeratorRequest) error {
	elements := map[string]interface{}{
		"username": obj.Username,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertModeratorRequestConstraints checks if the values respects the defined constraints
func AssertModeratorRequestConstraints(obj ModeratorRequest) error {
	return nil
}

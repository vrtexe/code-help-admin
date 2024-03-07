/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type CommentRequest struct {

	Content string `json:"content"`
}

// AssertCommentRequestRequired checks if the required fields are not zero-ed
func AssertCommentRequestRequired(obj CommentRequest) error {
	elements := map[string]interface{}{
		"content": obj.Content,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCommentRequestConstraints checks if the values respects the defined constraints
func AssertCommentRequestConstraints(obj CommentRequest) error {
	return nil
}

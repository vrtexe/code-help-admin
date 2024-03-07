/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type CommentReplies struct {

	Count int32 `json:"count"`

	Replies []Comment `json:"replies,omitempty"`
}

// AssertCommentRepliesRequired checks if the required fields are not zero-ed
func AssertCommentRepliesRequired(obj CommentReplies) error {
	elements := map[string]interface{}{
		"count": obj.Count,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Replies {
		if err := AssertCommentRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCommentRepliesConstraints checks if the values respects the defined constraints
func AssertCommentRepliesConstraints(obj CommentReplies) error {
	return nil
}

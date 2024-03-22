/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type Communities struct {

	Communities []Community `json:"communities"`
}

// AssertCommunitiesRequired checks if the required fields are not zero-ed
func AssertCommunitiesRequired(obj Communities) error {
	elements := map[string]interface{}{
		"communities": obj.Communities,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Communities {
		if err := AssertCommunityRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCommunitiesConstraints checks if the values respects the defined constraints
func AssertCommunitiesConstraints(obj Communities) error {
	return nil
}
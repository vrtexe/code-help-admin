/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type ShortCommunities struct {

	Communities []ShortCommunity `json:"communities"`
}

// AssertShortCommunitiesRequired checks if the required fields are not zero-ed
func AssertShortCommunitiesRequired(obj ShortCommunities) error {
	elements := map[string]interface{}{
		"communities": obj.Communities,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Communities {
		if err := AssertShortCommunityRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertShortCommunitiesConstraints checks if the values respects the defined constraints
func AssertShortCommunitiesConstraints(obj ShortCommunities) error {
	return nil
}

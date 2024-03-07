/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api




type CommunityRequestCategories struct {

	Uids []string `json:"uids,omitempty"`
}

// AssertCommunityRequestCategoriesRequired checks if the required fields are not zero-ed
func AssertCommunityRequestCategoriesRequired(obj CommunityRequestCategories) error {
	return nil
}

// AssertCommunityRequestCategoriesConstraints checks if the values respects the defined constraints
func AssertCommunityRequestCategoriesConstraints(obj CommunityRequestCategories) error {
	return nil
}

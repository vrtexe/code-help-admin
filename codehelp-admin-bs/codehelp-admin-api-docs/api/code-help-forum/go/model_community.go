/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api


import (
	"time"
)



type Community struct {

	Name string `json:"name"`

	Description string `json:"description"`

	Image string `json:"image,omitempty"`

	Categories []Category `json:"categories"`

	Admin User `json:"admin"`

	Posts []ShortPost `json:"posts"`

	Joined bool `json:"joined"`

	Moderators []User `json:"moderators"`

	Created time.Time `json:"created"`
}

// AssertCommunityRequired checks if the required fields are not zero-ed
func AssertCommunityRequired(obj Community) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"description": obj.Description,
		"categories": obj.Categories,
		"admin": obj.Admin,
		"posts": obj.Posts,
		"joined": obj.Joined,
		"moderators": obj.Moderators,
		"created": obj.Created,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Categories {
		if err := AssertCategoryRequired(el); err != nil {
			return err
		}
	}
	if err := AssertUserRequired(obj.Admin); err != nil {
		return err
	}
	for _, el := range obj.Posts {
		if err := AssertShortPostRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Moderators {
		if err := AssertUserRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCommunityConstraints checks if the values respects the defined constraints
func AssertCommunityConstraints(obj Community) error {
	return nil
}

/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api




type Category struct {

	Name string `json:"name,omitempty"`
}

// AssertCategoryRequired checks if the required fields are not zero-ed
func AssertCategoryRequired(obj Category) error {
	return nil
}

// AssertCategoryConstraints checks if the values respects the defined constraints
func AssertCategoryConstraints(obj Category) error {
	return nil
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger Potion maker - OpenAPI 3.0
 *
 * Сервис изготовления зелий
 *
 * API version: 1.0.0
 */

package openapi




type Ingredient struct {

	Id string `json:"id"`

	Name string `json:"name"`
}

// AssertIngredientRequired checks if the required fields are not zero-ed
func AssertIngredientRequired(obj Ingredient) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertIngredientConstraints checks if the values respects the defined constraints
func AssertIngredientConstraints(obj Ingredient) error {
	return nil
}
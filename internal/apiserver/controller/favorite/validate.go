package favorite

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a policy object is valid.
func (c *FavoriteActionRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

// Validate validates that a policy object is valid.
func (c *FavoriteListRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

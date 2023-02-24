package user

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a policy object is valid.
func (c *UserGetRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

// Validate validates that a policy object is valid.
func (c *UserRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

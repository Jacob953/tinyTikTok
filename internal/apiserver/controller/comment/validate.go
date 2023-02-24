package comment

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a policy object is valid.
func (c *CommentRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

// Validate validates that a secret object is valid.
func (c *CommentCreateRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

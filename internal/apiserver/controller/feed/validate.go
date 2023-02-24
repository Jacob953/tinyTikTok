package feed

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a secret object is valid.
func (c *FeedRequest) Validate() field.ErrorList {
	val := validation.NewValidator(c)

	return val.Validate()
}

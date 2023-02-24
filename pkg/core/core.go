package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
)

// ErrResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), Response{
			StatusCode: coder.Code(),
			StatusMsg:  coder.String(),
		})

		return
	}

	c.JSON(http.StatusOK, data)
}

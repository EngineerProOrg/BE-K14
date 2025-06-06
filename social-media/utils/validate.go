package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate[T any](c *gin.Context) (*T, bool) {
	var req T

	if err := c.ShouldBindJSON(&req); err != nil {
		// Validation errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(map[string]string)
			for _, fe := range ve {
				field := strings.ToLower(fe.Field())

				var msg string
				switch fe.Tag() {

				case "required":
					msg = "is required"

				case "email":
					msg = "must be a valid email"

				case "gt":
					msg = "must be greater than 0"

				case "notblank":
					msg = "must be not blank"

				default:
					msg = fmt.Sprintf("failed on '%s'", fe.Tag())
				}

				errs[field] = msg
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
			return nil, false
		}

		// The other exceptions (Invalid request format)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return nil, false
	}

	return &req, true
}

package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteConflictResponse(errors []error, c *gin.Context) {
	writeErrorResponse(errors, c, http.StatusConflict)
}

func WriteBadRequestResponse(errors []error, c *gin.Context) {
	writeErrorResponse(errors, c, http.StatusBadRequest)
}

func writeErrorResponse(errors []error, c *gin.Context, statusCode int) {
	c.JSON(statusCode, gin.H{
		"errors": errorsToStrings(errors),
	})
}

func errorsToStrings(errs []error) []string {
	out := make([]string, len(errs))
	for i, e := range errs {
		out[i] = e.Error()
	}
	return out
}

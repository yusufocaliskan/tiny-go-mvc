package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check for the role
// For some page that should be protected
// from other roles
func ForceOnlyRole(allowedRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, _ := ctx.Get("UserRole")

		//Do it only for non-admin roles
		if userRole != "admin" && userRole != allowedRole {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not allowed permission"})
			ctx.Abort()
		}

		ctx.Next()
	}
}

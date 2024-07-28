package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"saml_sso/internal/services"
)

// UserRoutes sets up the user-related routes
func MemberRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/profile/:stytch_member_id", func(c *gin.Context) {
		services.GetUserProfile(c, db)
	})
}

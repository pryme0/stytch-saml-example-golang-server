package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"saml_sso/internal/services"
)

func AuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/authenticate", func(c *gin.Context) {
		services.Authenticate(c, db)
	})
	r.POST("/signup", func(c *gin.Context) {
		services.SignUp(c, db)
	})

	r.PUT("/tenants/update/saml-connection/:id", func(c *gin.Context) {
		services.UpdateSamlConnection(c, db)
	})

}

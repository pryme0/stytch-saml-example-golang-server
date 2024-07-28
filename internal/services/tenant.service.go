package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"saml_sso/internal/models"
	"saml_sso/internal/utils"
)

// GetTenantByID retrieves a tenant by ID
func GetTenantByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var tenant models.Tenant
	result := db.First(&tenant, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Tenant not found")
		} else {
			utils.InternalServerError(c, "Error retrieving tenant")
		}
		return
	}

	utils.OK(c, tenant)
}

// GetStytchOrgId retrieves the Stytch organization ID by email
func GetStytchOrgId(c *gin.Context, db *gorm.DB) {
	email := c.Param("email")

	var member models.Member
	result := db.Preload("Tenant").First(&member, "email = ?", email)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Tenant not found")
		} else {
			utils.InternalServerError(c, "Error retrieving tenant")
		}
		return
	}

	utils.OK(c, gin.H{"organization_id": member.Tenant.StytchOrganizationId})
}

// GetTenantByName retrieves a tenant by company name
func GetTenantByName(c *gin.Context, db *gorm.DB) {
	company_name := c.Param("company_name")

	var tenant models.Tenant
	result := db.First(&tenant, "company_name = ?", company_name)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Tenant not found")
		} else {
			utils.InternalServerError(c, "Error retrieving tenant")
		}
		return
	}

	utils.OK(c, gin.H{"connection_id": tenant.ConnectionID})
}

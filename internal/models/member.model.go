package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `gorm:"uniqueIndex" json:"email"`
	TenantID       uint   `json:"tenant_id"`
	StytchMemberID string `json:"stytch_member_id"`
	Tenant         Tenant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tenant"`
}

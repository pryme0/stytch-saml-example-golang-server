package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	ConnectionID         string   `json:"connection_id"`
	StytchOrganizationId string   `json:"stytch_organization_id"`
	IdpSignOnUrl         string   `json:"idp_sign_on_url"`
	IdpIssuerUrl         string   `json:"idp_issuer_url"`
	StytchAudienceUrl    string   `json:"stytch_audience_url"`
	StytchAcsUrl         string   `json:"stytch_acs_url"`
	CompanyName          string   `json:"company_name"`
	Members              []Member `gorm:"foreignKey:TenantID" json:"members"`
}

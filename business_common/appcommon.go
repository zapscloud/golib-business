package business_common

import (
	"log"
)

// Business module tables =================================
const (
	DbAppAccess   = "app_access"
	DbAppContacts = "app_contacts"

	DbAppSites       = "app_sites"
	DbAppTerritories = "app_territories"

	DbBusinessProfiles = "business_profiles"
	DbBusinessRoles    = "business_roles"
	DbBusinessUsers    = "business_users"
)

// Business module table fields
const (
	FLD_BUSINESS_ID           = "business_id"
	FLD_BUSINESS_NAME         = "bussiness_name"
	FLD_BUSINESS_COMM_EMAILID = "bussiness_comm_email_id" // Communication EMailId

	FLD_USER_ID    = "user_id"
	FLD_USER_ROLES = "user_roles"

	FLD_ROLE_ID   = "role_id"
	FLD_ROLE_NAME = "role_name"
	FLD_ROLD_DESC = "role_description"

	FLD_APP_ACCESS_ID  = "app_access_id"
	FLD_APP_CONTACT_ID = "app_contact_id"

	FLD_APP_SITE_ID      = "app_site_id"
	FLD_APP_TERRITORY_ID = "app_territory_id"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

}

func GetServiceModuleCode() string {
	return "S4"
}

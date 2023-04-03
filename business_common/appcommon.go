package business_common

import (
	"log"

	"github.com/zapscloud/golib-dbutils/db_common"
)

// Business module tables =================================
const (
	DbPrefix      = db_common.DB_COLLECTION_PREFIX
	DbAppAccess   = DbPrefix + "app_access"
	DbAppContacts = DbPrefix + "app_contacts"

	DbAppSites       = DbPrefix + "app_sites"
	DbAppTerritories = DbPrefix + "app_territories"

	DbBusinessProfiles = DbPrefix + "business_profiles"
	DbBusinessRoles    = DbPrefix + "business_roles"
	DbBusinessUsers    = DbPrefix + "business_users"
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

package business_common

import (
	"log"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_common"
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
	FLD_BUSINESS_ID      = platform_common.FLD_BUSINESS_ID
	FLD_BUSINESS_NAME    = platform_common.FLD_BUSINESS_NAME
	FLD_BUSINESS_EMAILID = platform_common.FLD_BUSINESS_EMAILID

	FLD_USER_ID    = platform_common.FLD_APP_USER_ID
	FLD_USER_ROLES = "user_roles"

	FLD_ROLE_ID   = platform_common.FLD_APP_ROLE_ID
	FLD_ROLE_NAME = platform_common.FLD_APP_ROLE_NAME
	FLD_ROLD_DESC = platform_common.FLD_APP_ROLE_DESC

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

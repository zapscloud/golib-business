package main

import (
	"fmt"
	"os"

	"github.com/zapscloud/golib-business/business_services"
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-utils/utils"
)

func GetZapsDBCreds() utils.Map {
	dbtype := db_common.DATABASE_TYPE_ZAPSDB
	dbkey := os.Getenv("ZAPS_APP_KEY")
	dbsecret := os.Getenv("ZAPS_APP_SECRET")
	dbapp := os.Getenv("ZAPS_APP")

	dbCreds := utils.Map{
		db_common.DB_TYPE:   dbtype,
		db_common.DB_APP:    dbapp,
		db_common.DB_KEY:    dbkey,
		db_common.DB_SECRET: dbsecret}

	return dbCreds
}

func ZapsDBMain(businessid string) (business_services.UserService, business_services.BusinessService, business_services.RoleService) {

	props := AddBusinessIdToMap(GetZapsDBCreds(), businessid)

	usersrv, err := business_services.NewUserService(props)
	fmt.Println("User Mongo Service Error ", err)
	bizsrv, err := business_services.NewBusinessService(props)
	fmt.Println("User Mongo Service Error ", err)
	rolesrv, err := business_services.NewRoleService(props)
	fmt.Println("User Mongo Service Error ", err)

	return usersrv, bizsrv, rolesrv
}

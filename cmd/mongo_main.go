package main

import (
	"fmt"
	"os"

	"github.com/zapscloud/golib-business/business_services"
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-utils/utils"
)

func GetMongoDBCreds() utils.Map {
	dbtype := db_common.DATABASE_TYPE_MONGODB
	dbuser := os.Getenv("MONGO_DB_USER")
	dbsecret := os.Getenv("MONGO_DB_SECRET")
	dbserver := os.Getenv("MONGO_DB_SERVER")
	dbname := os.Getenv("MONGO_DB_NAME")

	dbCreds := utils.Map{
		db_common.DB_TYPE:   dbtype,
		db_common.DB_SERVER: dbserver,
		db_common.DB_NAME:   dbname,
		db_common.DB_USER:   dbuser,
		db_common.DB_SECRET: dbsecret}

	return dbCreds
}

func MongoDBMain(businessid string) (business_services.UserService, business_services.BusinessService, business_services.RoleService) {

	props := AddBusinessIdToMap(GetMongoDBCreds(), businessid)
	usersrv, err := business_services.NewUserService(props)
	fmt.Println("User Mongo Service Error ", err)
	bizsrv, err := business_services.NewBusinessService(props)
	fmt.Println("User Mongo Service Error ", err)
	rolesrv, err := business_services.NewRoleService(props)
	fmt.Println("User Mongo Service Error ", err)
	return usersrv, bizsrv, rolesrv
}

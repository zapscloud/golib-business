package main

import (
	"fmt"
	"log"

	"github.com/kr/pretty"
	"github.com/zapscloud/golib-business/business_common"
	"github.com/zapscloud/golib-business/business_services"
	"github.com/zapscloud/golib-utils/utils"
)

func AddBusinessIdToMap(dbCred utils.Map, businessId string) utils.Map {
	dbCred[business_common.FLD_BUSINESS_ID] = businessId
	return dbCred
}

func main() {

	businessid := "business003"
	usersrv, bizsrv, rolesrv := MongoDBMain(businessid)
	// usersrv, bizsrv, rolesrv := ZapsDBMain(businessid)

	EmptyBusiness(bizsrv)
	// DeleteBusiness(bizsrv)
	// CreateBusiness(bizsrv)
	// GetBusiness(bizsrv)

	EmptyBusinessUser(usersrv)
	// UpdateUser(usersrv)
	// CreateUser(usersrv)
	// DeleteUser(usersrv)
	// ListUsers(usersrv)
	// GetUser(usersrv)
	// FindUser(usersrv)

	EmptyBusinessRole(rolesrv)
	// DeleteRole(rolesrv)
	// CreateRole(rolesrv)
	// UpdateRole(rolesrv)
	ListRoles(rolesrv)
	// GetRole(rolesrv)
	// FindRole(rolesrv)

}

func EmptyBusinessUser(srv business_services.UserService) {
	fmt.Println("Business User Service ")
}

func EmptyBusiness(srv business_services.BusinessService) {
	fmt.Println("Business Service ")
}

func EmptyBusinessRole(srv business_services.RoleService) {
	fmt.Println("Role Service ")
}

func ListUsers(srv business_services.UserService) {

	response, err := srv.List("", "", 0, 2)

	fmt.Println("List success ", response, err)

	for idx, value := range response {
		fmt.Println("Value ", idx)
		pretty.Print(value)
	}
}

func GetUser(srv business_services.UserService) {

	resdata, err := srv.Get("user003")

	fmt.Println("getUser success ", err)
	fmt.Println("getUser Value ")
	pretty.Print(resdata)
}

func FindUser(srv business_services.UserService) {

	filter := fmt.Sprintf(`{"%s":%v}`, "is_active", true)
	res, err := srv.Find(filter)
	fmt.Println("Get Role", err)
	pretty.Println(res)

}

func UpdateUser(srv business_services.UserService) {

	indata := utils.Map{}

	active := true
	indata["is_active"] = &active

	srv.BeginTransaction()
	resdata, err := srv.Update("user003", indata)
	fmt.Println("updateUser success ", err)
	fmt.Println("updateUser Value ")

	// srv.RollbackTransaction()
	srv.CommitTransaction()

	pretty.Print(resdata)
}

func CreateUser(srv business_services.UserService) {

	userid := "user004"
	// businessid := "business003"

	srv.BeginTransaction()
	resdata, err := srv.Create(utils.Map{"userid": userid})
	fmt.Println("createUser success ", err)
	fmt.Println("createUser Value ")

	if err != nil {
		srv.RollbackTransaction()
	} else {
		srv.CommitTransaction()
	}

	pretty.Print(resdata)
}

func DeleteUser(srv business_services.UserService) {

	srv.BeginTransaction()
	err := srv.Delete("user004", false)
	fmt.Println("updateUser success ", err)
	fmt.Println("updateUser Value ")

	if err != nil {
		srv.RollbackTransaction()
	} else {
		srv.CommitTransaction()
	}
}

func TxnUser(srv business_services.UserService) {

	indata := utils.Map{}

	indata[business_common.FLD_USER_ID] = "user001"
	indata[business_common.FLD_BUSINESS_ID] = "c6tdohn1584s70j0i76g"
	indata["user_name"] = "Sample Insert 01"

	userprofileid := "user001"
	// businessid := "c6tdohn1584s70j0i76g"

	var err error

	srv.BeginTransaction()
	defer func() {
		log.Println("Defer exit ", err)
		if err != nil {
			srv.RollbackTransaction()
		} else {
			srv.CommitTransaction()
		}
	}()

	resdata, err := srv.Create(utils.Map{"userid": userprofileid})
	fmt.Println("insertUser success ", err)
	fmt.Println("insertUser Value ", resdata)
	if err != nil {
		return
	}

	updata := utils.Map{}
	updata["user_name"] = "Sample Updated 02"
	resdata, err = srv.Update("user001", updata)
	fmt.Println("updateUser success ", err)
	fmt.Println("updateUser Value ", resdata)
	if err != nil {
		return
	}

	err = srv.Delete("user001", false)
	fmt.Println("deleteUser success ", err)
	fmt.Println("deleteUser Value ")
	if err != nil {
		return
	}
}

func CreateBusiness(srv business_services.BusinessService) {

	srv.BeginTransaction()
	resdata, err := srv.Create(utils.Map{})
	fmt.Println("createBusiness success ", err)
	fmt.Println("createBusiness Value ")

	if err != nil {
		srv.RollbackTransaction()
	} else {
		srv.CommitTransaction()
	}

	pretty.Print(resdata)
}

func GetBusiness(srv business_services.BusinessService) {

	resdata, err := srv.GetDetails()

	fmt.Println("getUser success ", err)
	fmt.Println("getUser Value ")
	pretty.Print(resdata)
}

func DeleteBusiness(srv business_services.BusinessService) {

	err := srv.Delete()

	fmt.Println("deleteBusiness success ", err)
	fmt.Println("deleteBusiness Value ")
}

func CreateRole(srv business_services.RoleService) {

	indata := utils.Map{
		business_common.FLD_ROLE_ID: "role003",
		"role_name":                 "Demo Role 003",
		"role_scope":                "admin",
	}

	res, err := srv.Create(indata)
	fmt.Println("Create Role", err)
	pretty.Println(res)

}

func GetRole(srv business_services.RoleService) {
	res, err := srv.Get("role001")
	fmt.Println("Get Role", err)
	pretty.Println(res)

}

func FindRole(srv business_services.RoleService) {

	filter := fmt.Sprintf(`{"%s":"%s"}`, "role_scope", "admin")
	res, err := srv.Find(filter)
	fmt.Println("Get Role", err)
	pretty.Println(res)

}

func UpdateRole(srv business_services.RoleService) {

	indata := utils.Map{
		business_common.FLD_ROLE_ID: "role001",
		"role_name":                 "Demo Role 001 Updated",
		"is_active":                 true,
	}

	res, err := srv.Update("role001", indata)
	fmt.Println("Update Role", err)
	pretty.Println(res)

}

func DeleteRole(srv business_services.RoleService) {

	srv.BeginTransaction()
	err := srv.Delete("role001", false)
	fmt.Println("DeleteRole success ", err)
	fmt.Println("DeleteRole Value ")

	if err != nil {
		srv.RollbackTransaction()
	} else {
		srv.CommitTransaction()
	}
}

func ListRoles(srv business_services.RoleService) {

	filter := "" //fmt.Sprintf(`{"%s":"%s"}`, "role_scope", "admin")

	sort := `{ "role_scope":1, business_common.FLD_ROLE_ID:1}`

	res, err := srv.List(filter, sort, 0, 0)
	fmt.Println("List User success ", err)
	fmt.Println("List User summary ", res)
	pretty.Print(res)
}

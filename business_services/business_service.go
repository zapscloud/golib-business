package business_services

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib-business/business_common"
	"github.com/zapscloud/golib-business/business_repository"
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-dbutils/db_utils"
	"github.com/zapscloud/golib-platform/platform_common"
	"github.com/zapscloud/golib-platform/platform_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// LoyaltyCardService - Users Service structure
type BusinessService interface {
	GetDetails() (utils.Map, error)
	Create(indata utils.Map) (utils.Map, error)
	Update(indata utils.Map) (utils.Map, error)
	Find(filter string) (utils.Map, error)
	Delete() error

	BeginTransaction()
	CommitTransaction()
	RollbackTransaction()

	initializeService()
	EndService()
}

// LoyaltyCardService - Users Service structure
type businessBaseService struct {
	db_utils.DatabaseService
	daoBusiness         business_repository.BusinessDao
	daoUser             business_repository.UserDao
	daoContact          business_repository.ContactDao
	daoSysUser          platform_repository.SysUserDao
	daoPlatformBusiness platform_repository.BusinessDao
	child               BusinessService
	businessID          string
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

// NewBusinessService - Construct Business Service
func NewBusinessService(props utils.Map) (BusinessService, error) {
	funcode := business_common.GetServiceModuleCode() + "M" + "01"

	p := businessBaseService{}

	err := p.OpenDatabaseService(props)
	if err != nil {
		// log.Fatal(err)
		log.Println("NewBusinessMongoService App Connection Error ", err)
		return nil, err
	}
	log.Printf("BusinessService ")
	// Verify whether the business id data passed
	businessId, err := utils.IsMemberExist(props, business_common.FLD_BUSINESS_ID)
	if err != nil {
		return nil, err
	}

	// Assign the BusinessId
	p.businessID = businessId

	// Initialize other Service
	p.initializeAppService()

	dataBusiness, err := p.daoPlatformBusiness.GetDetails(businessId)
	if err != nil {
		err := &utils.AppError{ErrorCode: funcode + "01", ErrorMsg: "Invalid business_id", ErrorDetail: "Given app_business_id is not exist"}
		return nil, err
	}

	businessRegion := dataBusiness[platform_common.FLD_BUSINESS_REGION_ID].(string)

	daoRegion := platform_repository.NewRegionDao(p.GetClient())
	dataRegion, err := daoRegion.GetDetails(businessRegion)
	if err != nil {
		log.Println("NewBusinessService Get Region details Error ", err)
		return nil, err
	}

	if _, dataok := dataRegion[platform_common.FLD_REGION_MONGODB_SERVER]; !dataok {
		err := &utils.AppError{ErrorCode: funcode + "02", ErrorMsg: "Missing MongoDB Values", ErrorDetail: "Missing MongoDB Values for the given region details"}
		return nil, err
	}
	if _, dataok := dataRegion[platform_common.FLD_REGION_MONGODB_NAME]; !dataok {
		err := &utils.AppError{ErrorCode: funcode + "03", ErrorMsg: "Missing MongoDB Values", ErrorDetail: "Missing MongoDB Values for the given region details"}
		return nil, err
	}
	if _, dataok := dataRegion[platform_common.FLD_REGION_MONGODB_USER]; !dataok {
		err := &utils.AppError{ErrorCode: funcode + "04", ErrorMsg: "Missing MongoDB Values", ErrorDetail: "Missing MongoDB Values for the given region details"}
		return nil, err
	}
	if _, dataok := dataRegion[platform_common.FLD_REGION_MONGODB_SECRET]; !dataok {
		err := &utils.AppError{ErrorCode: funcode + "05", ErrorMsg: "Missing MongoDB Values", ErrorDetail: "Missing MongoDB Values for the given region details"}
		return nil, err
	}

	dbtype := props[db_common.DB_TYPE].(db_common.DatabaseType)
	dbserver := dataRegion[platform_common.FLD_REGION_MONGODB_SERVER].(string)
	dbname := dataRegion[platform_common.FLD_REGION_MONGODB_NAME].(string)
	dbuser := dataRegion[platform_common.FLD_REGION_MONGODB_USER].(string)
	dbsecret := dataRegion[platform_common.FLD_REGION_MONGODB_SECRET].(string)

	if tenantdata, tenantok := dataBusiness[platform_common.FLD_BUSINESS_IS_TENANT_DB]; tenantok && tenantdata.(bool) {
		dbname = dataRegion[platform_common.FLD_REGION_MONGODB_NAME].(string) + "-" + businessId
	}

	// Prepare DBCredentials from the new Region Information
	props = utils.Map{
		db_common.DB_TYPE:   dbtype,
		db_common.DB_SERVER: dbserver,
		db_common.DB_NAME:   dbname,
		db_common.DB_USER:   dbuser,
		db_common.DB_SECRET: dbsecret}

	// Close Previously Opened driver
	p.CloseDatabaseService()

	// Open
	err = p.OpenDatabaseService(props)
	if err != nil {
		// log.Fatal(err)
		log.Println("NewBusinessService Connection Error ", err)
		return nil, err
	}

	p.initializeService()
	p.child = &p

	return &p, err
}

// EndLoyaltyCardService - Close all the services
func (p *businessBaseService) EndService() {
	log.Printf("EndBusinessMongoService ")
	p.CloseDatabaseService()
}

func (p *businessBaseService) initializeService() {
	log.Printf("BusinessMongoService:: GetBusinessDao ")
	p.daoBusiness = business_repository.NewBusinessDao(p.GetClient(), p.businessID)
	p.daoUser = business_repository.NewUserDao(p.GetClient(), p.businessID)
	p.daoContact = business_repository.NewContactDao(p.GetClient(), p.businessID)
}

func (p *businessBaseService) initializeAppService() {
	log.Printf("BusinessMongoService:: GetBusinessDao ")
	p.daoSysUser = platform_repository.NewSysUserDao(p.GetClient())
	p.daoPlatformBusiness = platform_repository.NewBusinessDao(p.GetClient())
}

// Create - Create Service
func (p *businessBaseService) Create(indata utils.Map) (utils.Map, error) {

	funcode := business_common.GetServiceModuleCode() + "01"

	log.Println("BusinessService::Create - Begin")

	// Add Business Id
	indata[business_common.FLD_BUSINESS_ID] = p.businessID

	// Create Business
	dataBusiness, err := p.daoBusiness.Create(indata)
	if err != nil {
		log.Println("Business Create Error  ", err)
		err := &utils.AppError{ErrorCode: funcode + "02", ErrorMsg: "Business Create Error", ErrorDetail: "Error while creating business tenant"}
		return nil, err
	}

	log.Println("Business create  ", dataBusiness)

	log.Println("BusinessService::Create - End ")
	return dataBusiness, nil
}

// Update - Update Service
func (p *businessBaseService) Update(indata utils.Map) (utils.Map, error) {

	log.Println("BusinessService::Update - Begin")

	data, err := p.daoBusiness.Get(p.businessID)
	if err != nil {
		return data, err
	}

	// Delete Business Id if exist
	delete(indata, business_common.FLD_BUSINESS_ID)

	data, err = p.daoBusiness.Update(indata)
	log.Println("BusinessService::Update - End ")
	return data, err
}

// FindByCode - Find By Code
func (p *businessBaseService) GetDetails() (utils.Map, error) {
	funcode := business_common.GetServiceModuleCode() + "02"
	log.Printf("BusinessService::GetDetails::  Begin %v", p.businessID)

	data, err := p.daoBusiness.Get(p.businessID)
	if err != nil {
		err := &utils.AppError{ErrorCode: funcode + "02", ErrorMsg: "Invalid app_user_id", ErrorDetail: "Given app_user_id is not exist"}
		return nil, err
	}

	log.Println("BusinessService::GetDetails:: End ", err)
	return data, err
}

func (p *businessBaseService) Find(filter string) (utils.Map, error) {
	fmt.Println("BusinessService::FindByCode::  Begin ", filter)

	data, err := p.daoContact.Find(filter)
	log.Println("BusinessService::FindByCode:: End ", data, err)
	return data, err
}

// Delete - Delete Service
func (p *businessBaseService) Delete() error {

	log.Println("BusinessService::Delete - Begin", p.businessID)

	result, err := p.daoContact.DeleteAll()
	if err != nil {
		return err
	}
	log.Printf("BusinessService::DeleteAll - Contact  %v", result)

	result, err = p.daoBusiness.Delete(p.businessID)
	if err != nil {
		return err
	}
	log.Printf("BusinessService::Delete - End %v", result)
	return nil
}

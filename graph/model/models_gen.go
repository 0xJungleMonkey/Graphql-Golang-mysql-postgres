// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Address struct {
	ID               int    `json:"id"`
	AddressType      string `json:"address_type"`
	Status           string `json:"status"`
	Entity           string `json:"entity"`
	NumberAndStreet  string `json:"number_and_street"`
	SuiteOrApartment string `json:"suite_or_apartment"`
	City             string `json:"city"`
	PostalCode       string `json:"postal_code"`
	Country          string `json:"country"`
	Notes            string `json:"notes"`
}

type Building struct {
	CustomerID                       int                 `json:"customer_id"`
	AddressID                        int                 `json:"address_id"`
	Address                          *Address            `json:"address"`
	Customer                         *Customer           `json:"customer"`
	ID                               int                 `json:"id"`
	Interventions                    []*FactIntervention `json:"interventions"`
	FullNameOfBuildingAdmin          int                 `json:"FullNameOfBuildingAdmin"`
	EmailOfAdminOfBuilding           string              `json:"EmailOfAdminOfBuilding"`
	PhoneNumOfBuildingAdmin          string              `json:"PhoneNumOfBuildingAdmin"`
	FullNameOfTechContactForBuilding string              `json:"FullNameOfTechContactForBuilding"`
	TechContactEmailForBuilding      string              `json:"TechContactEmailForBuilding"`
	TechContactPhoneForBuilding      string              `json:"TechContactPhoneForBuilding"`
	BuildingDetails                  *BuildingDetails    `json:"building_details"`
}

type BuildingCustomerFactIntervention struct {
	Building            *Building           `json:"building"`
	Interventionhistory []*FactIntervention `json:"interventionhistory"`
}

type BuildingDetails struct {
	BuildingID     int    `json:"building_id"`
	ID             int    `json:"id"`
	InformationKey string `json:"InformationKey"`
	Value          string `json:"Value"`
}

type BuildingFactIntervention struct {
	EmployeeID       int               `json:"employee_id"`
	Building         *Building         `json:"building"`
	FactIntervention *FactIntervention `json:"factIntervention"`
}

type Customer struct {
	AddressID                int    `json:"address_id"`
	UserID                   int    `json:"user_id"`
	ID                       int    `json:"id"`
	CustomerCreationDate     string `json:"CustomerCreationDate"`
	CompanyName              string `json:"CompanyName"`
	CompanyHQAdress          string `json:"CompanyHQAdress"`
	FullNameOfCompanyContact string `json:"FullNameOfCompanyContact"`
	CompanyContactPhone      string `json:"CompanyContactPhone"`
	CompanyContactEMail      string `json:"CompanyContactEMail"`
	CompanyDesc              string `json:"CompanyDesc"`
	FullNameServiceTechAuth  string `json:"FullNameServiceTechAuth"`
	TechAuthPhoneService     string `json:"TechAuthPhoneService"`
	TechManagerEmailService  string `json:"TechManagerEmailService"`
}

type Employee struct {
	UserID            int                 `json:"user_id"`
	ID                int                 `json:"id"`
	FirstName         string              `json:"first_name"`
	LastName          string              `json:"last_name"`
	Title             string              `json:"title"`
	Email             string              `json:"email"`
	FactInterventions []*FactIntervention `json:"factInterventions"`
}

type EmployeeBuildingandDetails struct {
	Employee          *Employee                   `json:"employee"`
	InterventionsList []*BuildingFactIntervention `json:"interventions_list"`
	Interventions     []*FactIntervention         `json:"interventions"`
	Buildings         []*Building                 `json:"buildings"`
}

type FactIntervention struct {
	ID            int       `json:"id"`
	EmployeeID    int       `json:"employee_id"`
	BuildingID    int       `json:"building_id"`
	BatteryID     int       `json:"battery_id"`
	ColumnID      int       `json:"column_id"`
	ElevatorID    int       `json:"elevator_id"`
	StartDatetime string    `json:"start_datetime"`
	EndDatetime   string    `json:"end_datetime"`
	Result        string    `json:"result"`
	Report        string    `json:"report"`
	Status        string    `json:"status"`
	Building      *Building `json:"building"`
}

type RadEmployee struct {
	UserID    int    `json:"user_id"`
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
	Email     string `json:"email"`
}

type RadEmployeeBuildingandDetails struct {
	Employee      *Employee           `json:"employee"`
	Interventions []*FactIntervention `json:"interventions"`
}

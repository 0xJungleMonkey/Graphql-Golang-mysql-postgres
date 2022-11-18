package graph

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/OxJungleMonkey/go-graphql/graph/model"
	"github.com/jinzhu/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB1                      *gorm.DB
	DB2                      *gorm.DB
	newFactInterventionStore []*model.FactIntervention
	// newBuildingStore         []*model.Building
	// newAddressStore          []*model.Address
	// newEmployeeStore         []*model.Employee
	FactInterventionStore map[int]model.FactIntervention
	BuildingStore         map[int]model.Building
	AddressStore          map[int]model.Address
	EmployeeStore         map[int]model.Employee
	// radEmployeeBuildingandDetails map[int]model.RadEmployeeBuildingandDetails

	// tiredStore               map[int]model.BuildingFactIntervention
	// mydb                  gorm.DB
}

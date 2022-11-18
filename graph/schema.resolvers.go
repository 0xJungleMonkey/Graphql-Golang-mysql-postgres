package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OxJungleMonkey/go-graphql/graph/generated"
	"github.com/OxJungleMonkey/go-graphql/graph/model"
)

// Addresses is the resolver for the addresses field.
func (r *queryResolver) Addresses(ctx context.Context) ([]*model.Address, error) {
	panic(fmt.Errorf("not implemented: Addresses - addresses"))
}

// Buildings is the resolver for the buildings field.
func (r *queryResolver) Buildings(ctx context.Context) ([]*model.Building, error) {
	panic(fmt.Errorf("not implemented: Buildings - buildings"))
}

// Facts is the resolver for the facts field.
func (r *queryResolver) Facts(ctx context.Context) ([]*model.FactIntervention, error) {
	panic(fmt.Errorf("not implemented: Facts - facts"))
}

// Fact is the resolver for the fact field.
func (r *queryResolver) Fact(ctx context.Context, buildingID int) (*model.FactIntervention, error) {
	panic(fmt.Errorf("not implemented: Fact - fact"))
}

// Building is the resolver for the building field.
func (r *queryResolver) Building(ctx context.Context, id int) (*model.Building, error) {
	panic(fmt.Errorf("not implemented: Building - building"))
}

// GetAddressofInterventionByInterventionID is the resolver for the GetAddressofInterventionByInterventionId field.
func (r *queryResolver) GetAddressofInterventionByInterventionID(ctx context.Context, id int) (*model.BuildingFactIntervention, error) {
	var m model.BuildingFactIntervention
	var factintervention model.FactIntervention = r.FactInterventionStore[id]
	var b model.Building = r.BuildingStore[factintervention.BuildingID]
	m.Building = &b
	m.FactIntervention = &factintervention
	// r.DB1.Set("gorm:auto_preload", true).Find(&m)
	r.DB2.Set("gorm:auto_preload", true).First(&factintervention, id)
	r.DB1.Set("gorm:auto_preload", true).First(&b, factintervention.BuildingID)
	return &m, nil
}

// v1
func (r *queryResolver) GetCustomerInfoandBuildingDetailsByBuildingID(ctx context.Context, id int) (*model.BuildingCustomerFactIntervention, error) {
	var q2 model.BuildingCustomerFactIntervention
	var building model.Building = r.BuildingStore[id]
	q2.Building = &building
	for _, tempFact := range r.newFactInterventionStore {
		if tempFact.BuildingID == id {
			q2.Building.Interventions = append(q2.Building.Interventions, tempFact)
		}
	}
	r.DB1.Set("gorm:auto_preload", true).First(&building, id)
	r.DB2.Where("building_id = ?", id).Find(&q2.Building.Interventions)
	return &q2, nil
}

// GetBuildingandDetailsByEmployeeID is the resolver for the GetBuildingandDetailsByEmployeeId field.
func (r *queryResolver) GetBuildingandDetailsByEmployeeID(ctx context.Context, id int) (*model.EmployeeBuildingandDetails, error) {
	var xEmployeeBuildingandDetails model.EmployeeBuildingandDetails
	xEmployeeBuildingandDetails.Employee = new(model.Employee)
	fmt.Println("before")
	r.DB1.Set("gorm:auto_preload", false).First(&xEmployeeBuildingandDetails.Employee, id)
	r.DB2.Set("gorm:auto_preload", false).Where("employee_id = ?", &xEmployeeBuildingandDetails.Employee.ID).Find(&xEmployeeBuildingandDetails.Employee.FactInterventions)

	for _, myIntervention := range xEmployeeBuildingandDetails.Employee.FactInterventions {
		myIntervention.Building = new(model.Building)
		r.DB1.Set("gorm:auto_preload", false).First(&myIntervention.Building, myIntervention.BuildingID)
		myIntervention.Building.BuildingDetails = new(model.BuildingDetails)
		r.DB1.First(&myIntervention.Building.BuildingDetails, myIntervention.BuildingID)
	}
	return &xEmployeeBuildingandDetails, nil
}

// RadGetBuildingandDetailsByEmployeeID is the resolver for the radGetBuildingandDetailsByEmployeeId field.
func (r *queryResolver) RadGetBuildingandDetailsByEmployeeID(ctx context.Context, id int) (*model.RadEmployeeBuildingandDetails, error) {
	var modRadEmployeeBuildingandDetails model.RadEmployeeBuildingandDetails
	modRadEmployeeBuildingandDetails.Employee = new(model.Employee)
	fmt.Println("before")
	r.DB1.Set("gorm:auto_preload", false).First(&modRadEmployeeBuildingandDetails.Employee, id)
	r.DB2.Set("gorm:auto_preload", false).Where("employee_id = ?", &modRadEmployeeBuildingandDetails.Employee.ID).Find(&modRadEmployeeBuildingandDetails.Employee.FactInterventions)

	for _, myIntervention := range modRadEmployeeBuildingandDetails.Employee.FactInterventions {
		myIntervention.Building = new(model.Building)
		r.DB1.Set("gorm:auto_preload", false).First(&myIntervention.Building, myIntervention.BuildingID)
		myIntervention.Building.BuildingDetails = new(model.BuildingDetails)
		r.DB1.First(&myIntervention.Building.BuildingDetails, myIntervention.BuildingID)
	}
	return &modRadEmployeeBuildingandDetails, nil
}

// Q3 is the resolver for the q3 field.
func (r *queryResolver) Q3(ctx context.Context, employeeID int) (*model.EmployeeBuildingandDetails, error) {
	panic(fmt.Errorf("not implemented: Q3 - q3"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

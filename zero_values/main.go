package zero_values

import (
	"fmt"

	"github.com/go-logr/logr"
)

type Department interface {
	GetName() string
}

type NullDepartment struct{}

func (d NullDepartment) GetName() string {
	return "Unassigned"
}

type Company struct {
	PrimaryDepartment Department
}

func (c Company) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	var company Company
	company.DisplayInfo()
}

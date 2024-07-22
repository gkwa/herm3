package custom_getters

import (
	"fmt"

	"github.com/go-logr/logr"
)

type Department interface {
	GetName() string
}

type Company struct {
	PrimaryDepartment Department
}

func (c *Company) GetDepartmentName() string {
	if c.PrimaryDepartment == nil {
		return "No Department Assigned"
	}
	return c.PrimaryDepartment.GetName()
}

func (c *Company) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.GetDepartmentName())
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	company := &Company{}
	company.DisplayInfo()
}

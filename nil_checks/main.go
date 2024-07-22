package nil_checks

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

func (c *Company) DisplayInfo() error {
	if c.PrimaryDepartment == nil {
		return fmt.Errorf("primary department is nil")
	}
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
	return nil
}

type StandardDepartment struct {
	name string
}

func (d *StandardDepartment) GetName() string {
	return d.name
}

func NewStandardDepartment(name string) *StandardDepartment {
	return &StandardDepartment{name: name}
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")

	companyWithDept := &Company{PrimaryDepartment: NewStandardDepartment("Engineering")}
	err := companyWithDept.DisplayInfo()
	if err != nil {
		logger.Error(err, "Failed to display company info")
	}

	companyWithoutDept := &Company{}
	err = companyWithoutDept.DisplayInfo()
	if err != nil {
		logger.Error(err, "Failed to display company info")
	}
}

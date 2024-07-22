package null_object

import (
	"fmt"

	"github.com/go-logr/logr"
)

type Department interface {
	GetName() string
}

type NullDepartment struct{}

func (d NullDepartment) GetName() string {
	return "No Department"
}

type Company struct {
	PrimaryDepartment Department
}

func NewCompany() *Company {
	return &Company{PrimaryDepartment: NullDepartment{}}
}

func (c *Company) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	company := NewCompany()
	company.DisplayInfo()
}

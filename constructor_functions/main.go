package constructor_functions

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

func (c *Company) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
}

func NewCompany(dept Department) *Company {
	return &Company{
		PrimaryDepartment: dept,
	}
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
	dept := NewStandardDepartment("Engineering")
	myCompany := NewCompany(dept)
	myCompany.DisplayInfo()
}

package default_values

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

// NewCompany is a constructor function that ensures Company is properly initialized
func NewCompany(dept Department) *Company {
	if dept == nil {
		dept = &DefaultDepartment{}
	}
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

type DefaultDepartment struct{}

func (d *DefaultDepartment) GetName() string {
	return "Default Department"
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	myCompany := NewCompany(nil)
	myCompany.DisplayInfo()
}

package composition

import (
	"fmt"

	"github.com/go-logr/logr"
)

type Department struct {
	name string
}

func (d *Department) GetName() string {
	return d.name
}

type Company struct {
	primaryDepartment *Department
}

func (c *Company) SetPrimaryDepartment(dept *Department) {
	c.primaryDepartment = dept
}

func (c *Company) GetPrimaryDepartment() *Department {
	return c.primaryDepartment
}

func (c *Company) DisplayInfo() {
	if c.primaryDepartment != nil {
		fmt.Printf("Company's primary department: %s\n", c.primaryDepartment.GetName())
	} else {
		fmt.Println("Company has no primary department")
	}
}

func NewCompany() *Company {
	return &Company{}
}

func NewDepartment(name string) *Department {
	return &Department{name: name}
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")

	company := NewCompany()
	company.DisplayInfo()

	engineeringDept := NewDepartment("Engineering")
	company.SetPrimaryDepartment(engineeringDept)
	company.DisplayInfo()
}

package option_pattern

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

type CompanyOption func(*Company)

func WithDepartment(dept Department) CompanyOption {
	return func(c *Company) {
		c.PrimaryDepartment = dept
	}
}

func NewCompany(options ...CompanyOption) *Company {
	c := &Company{}
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *Company) DisplayInfo() {
	if c.PrimaryDepartment != nil {
		fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
	} else {
		fmt.Println("Company has no primary department")
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
	company := NewCompany(WithDepartment(dept))
	company.DisplayInfo()
}

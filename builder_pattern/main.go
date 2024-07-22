package builder_pattern

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
	if c.PrimaryDepartment != nil {
		fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
	} else {
		fmt.Println("Company has no primary department")
	}
}

type CompanyBuilder struct {
	company Company
}

func NewCompanyBuilder() *CompanyBuilder {
	return &CompanyBuilder{}
}

func (cb *CompanyBuilder) SetPrimaryDepartment(dept Department) *CompanyBuilder {
	cb.company.PrimaryDepartment = dept
	return cb
}

func (cb *CompanyBuilder) Build() *Company {
	return &cb.company
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
	company := NewCompanyBuilder().
		SetPrimaryDepartment(dept).
		Build()
	company.DisplayInfo()
}

package lazy_initialization

import (
	"fmt"
	"sync"

	"github.com/go-logr/logr"
)

type Department interface {
	GetName() string
}

type Company struct {
	primaryDepartment Department
	deptInitializer   func() Department
	initOnce          sync.Once
}

func (c *Company) GetPrimaryDepartment() Department {
	c.initOnce.Do(func() {
		if c.primaryDepartment == nil && c.deptInitializer != nil {
			c.primaryDepartment = c.deptInitializer()
		}
	})
	return c.primaryDepartment
}

func (c *Company) DisplayInfo() {
	dept := c.GetPrimaryDepartment()
	if dept != nil {
		fmt.Printf("Company's primary department: %s\n", dept.GetName())
	} else {
		fmt.Println("Company has no primary department")
	}
}

func NewCompany(initializer func() Department) *Company {
	return &Company{deptInitializer: initializer}
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
	company := NewCompany(func() Department {
		return NewStandardDepartment("Engineering")
	})
	company.DisplayInfo()
}

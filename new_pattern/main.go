package new_pattern

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

func New(dept Department) (*Company, error) {
	if dept == nil {
		return nil, fmt.Errorf("department cannot be nil")
	}
	return &Company{PrimaryDepartment: dept}, nil
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
	myCompany, err := New(dept)
	if err != nil {
		logger.Error(err, "Failed to create company")
		return
	}
	myCompany.DisplayInfo()
}

package generics

import (
	"fmt"

	"github.com/go-logr/logr"
)

type Department interface {
	GetName() string
}

type Company[T Department] struct {
	PrimaryDepartment T
}

func NewCompany[T Department](dept T) *Company[T] {
	return &Company[T]{PrimaryDepartment: dept}
}

func (c *Company[T]) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.GetName())
}

type EngineeringDepartment struct{}

func (d EngineeringDepartment) GetName() string {
	return "Engineering"
}

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	company := NewCompany(EngineeringDepartment{})
	company.DisplayInfo()
}

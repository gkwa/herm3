package linters

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

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	var company Company
	company.DisplayInfo()
}

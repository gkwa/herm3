package initial

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

func Main(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Main function")
	myCompany := &Company{}
	myCompany.DisplayInfo()
}

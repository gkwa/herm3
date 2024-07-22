
### Use constructor functions
Create a function that initializes a Company with a properly initialized Department implementation.

This solution fixes the nil interface problem in several ways:

1. Guaranteed Initialization: The constructor function `NewCompany` requires a `Department` as an argument, ensuring that a `Company` is always created with a non-nil `PrimaryDepartment`.

2. Encapsulation: The creation logic is encapsulated in the constructor, reducing the chance of incorrect initialization elsewhere in the code.

3. Clear Intent: Using a constructor function makes it clear to other developers that proper initialization is required and how it should be done.

4. Compile-time Safety: If a developer tries to create a `Company` without providing a `Department`, they'll get a compile-time error rather than a runtime panic.

5. Flexibility: The constructor can be easily extended to include additional initialization logic or validation if needed in the future.

By using this pattern, we shift the responsibility of providing a valid `Department` to the caller of `NewCompany`, making the code more robust and less prone to runtime errors. This approach aligns with Go's philosophy of making the zero value useful, as a `Company` created with the constructor is immediately usable without risk of nil pointer dereferences.


```go

package core

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
```

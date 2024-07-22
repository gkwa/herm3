### Preamble




https://github.com/gkwa/braveside/blob/f95791bd369d097b2d772ba9253204b1a5a66b1c/core/markdown_processor.go#L39-L41

![](https://i.imgur.com/qXRQJrp.png)






This is disturbing.  It requires the client know whether its components are null or else you'll get panic.  This seems like bad design.







[[braveside - client needs conditional#Use constructor functions]]




### Summary

We're exploring a common design issue in Go programming where struct initialization with interface fields can lead to unexpected runtime panics.

Our scenario involves two types: Company (which is our struct) and Department (which is now an interface).

The Company struct has a method that depends on an initialized Department interface.

However, if we're not careful with our initialization, we may encounter a panic when calling this method due to a nil interface value.

This exercise will help us understand the importance of proper struct initialization with interface fields and how to avoid such runtime errors in real-world scenarios like company management systems.

Now, let's create a new code example that demonstrates this problem:

```go
package main

import (
	"fmt"
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

func main() {
	// Create a Company without initializing its PrimaryDepartment
	myCompany := &Company{}

	// This will panic because myCompany.PrimaryDepartment is nil
	myCompany.DisplayInfo()
}
```

When you run this program, you'll see a panic with an error message like: "panic: runtime error: invalid memory address or nil pointer dereference".

This demonstrates the problem of not properly initializing structs with interface fields in a real-world scenario of a company management system.

### Use constructor functions

See use_constructor_functions.md

Create a function that initializes a Company with a properly initialized Department implementation.

### Implement default values
Provide a default Department implementation when creating a Company if one isn't specified.

### Employ the New() pattern
Use a New() function to create and return properly initialized Company structs with a valid Department implementation.

### Add nil checks
Include nil checks in methods before accessing the Department interface.

### Use composition over embedding
Restructure the relationship between Company and Department to avoid interface fields.

### Implement the Option Pattern
Use functional options to configure the Company struct during initialization, including setting a Department implementation.

### Lazy initialization
Initialize the Department implementation only when it's first accessed.

### Employ the Builder Pattern
Create a separate builder struct to construct a fully initialized Company with a valid Department implementation.

### Utilize Go's zero values
Design the Company struct so its zero value is usable without further initialization, possibly using a NullDepartment implementation.

### Use linters and static analysis tools
Catch potential nil interface dereferences before runtime.

### Implement custom getter methods
Create methods that safely return Department information, handling nil interface cases internally.

### Use the Null Object Pattern
Provide a "null" Department implementation instead of nil.

### Employ Go 1.18+ generics
Create generic constructor functions that ensure proper initialization across different struct types with interface fields.
```

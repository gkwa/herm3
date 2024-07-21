
### Summary



https://github.com/gkwa/braveside/blob/f95791bd369d097b2d772ba9253204b1a5a66b1c/core/markdown_processor.go#L39-L41

![](https://i.imgur.com/qXRQJrp.png)






This is disturbing.  It requires the client know whether its components are null or else you'll get panic.  This seems like bad design.







[[braveside - client needs conditional#Use constructor functions]]







How about we use a "Company" and "Department" scenario? This is a common real-world hierarchical structure that fits our use case well. I'll rewrite the problem definition and create a new code example to illustrate this.

Problem Definition:

We're exploring a common design issue in Go programming where nested struct initialization can lead to unexpected runtime panics.

Our scenario involves two structs: Company (which will be our outer struct) and Department (which will be our inner struct).

The Company struct has a method that depends on an initialized Department struct.

However, if we're not careful with our initialization, we may encounter a panic when calling this method due to a nil pointer dereference.

This exercise will help us understand the importance of proper struct initialization and how to avoid such runtime errors in real-world scenarios like company management systems.

Now, let's create a new code example that demonstrates this problem:

```go
package main

import (
	"fmt"
)

// Department represents a department within a company
type Department struct {
	Name string
}

// Company represents a company, which has a primary Department
type Company struct {
	PrimaryDepartment *Department
}

// DisplayInfo is a method of Company that uses the Department's name
func (c *Company) DisplayInfo() {
	fmt.Printf("Company's primary department: %s\n", c.PrimaryDepartment.Name)
}

func main() {
	// Create a Company without initializing its PrimaryDepartment
	myCompany := &Company{}

	// This will panic because myCompany.PrimaryDepartment is nil
	myCompany.DisplayInfo()
}

```


We define two structs: `Department` and `Company`.

The `Company` struct has a pointer to a `Department` called `PrimaryDepartment`.

The `Company` struct has a `DisplayInfo()` method that tries to access the `Name` field of its `PrimaryDepartment`.

In the `main()` function, we create a `Company` instance without initializing its `PrimaryDepartment` field.

When we call `myCompany.DisplayInfo()`, the program will panic because `myCompany.PrimaryDepartment` is nil, and we're trying to access the `Name` field of a nil pointer.

When you run this program, you'll see a panic with an error message like: "panic: runtime error: invalid memory address or nil pointer dereference".

This demonstrates the problem of not properly initializing nested structs in a real-world scenario of a company management system.

In the next steps, we can explore proper design decisions and solutions to avoid this issue, such as ensuring proper initialization of all fields, using constructor functions to create fully initialized structs, adding nil checks in methods that access potentially nil fields, and using value types instead of pointers where appropriate.



### Use constructor functions

Create a function that initializes a Company with a properly initialized Department.


[[braveside - client needs conditional#Summary|top]]
### Implement default values
Provide a default Department when creating a Company if one isn't specified.

[[braveside - client needs conditional#Summary|top]]

### Employ the New() pattern
Use a New() function to create and return properly initialized structs.

[[braveside - client needs conditional#Summary|top]]

### Utilize value types instead of pointers
Store Department as a value rather than a pointer in the Company struct.

[[braveside - client needs conditional#Summary|top]]

### Add nil checks
Include nil checks in methods before accessing potentially nil fields.

[[braveside - client needs conditional#Summary|top]]

### Use composition over embedding
Restructure the relationship between Company and Department to avoid nested structs.

[[braveside - client needs conditional#Summary|top]]

### Implement the Option Pattern
Use functional options to configure the Company struct during initialization.

[[braveside - client needs conditional#Summary|top]]

### Lazy initialization
Initialize the Department only when it's first accessed.

[[braveside - client needs conditional#Summary|top]]

### Use interfaces
Define an interface for Department-like structures, allowing for more flexible implementations.

[[braveside - client needs conditional#Summary|top]]

### Employ the Builder Pattern
Create a separate builder struct to construct a fully initialized Company.

[[braveside - client needs conditional#Summary|top]]

### Utilize Go's zero values
Design structs so their zero values are usable without further initialization.

[[braveside - client needs conditional#Summary|top]]

### Use linters and static analysis tools
Catch potential nil pointer dereferences before runtime.

[[braveside - client needs conditional#Summary|top]]

### Implement custom getter methods
Create methods that safely return Department information, handling nil cases internally.

[[braveside - client needs conditional#Summary|top]]

### Use the Null Object Pattern
Provide a "null" Department object instead of nil.

[[braveside - client needs conditional#Summary|top]]

### Employ Go 1.18+ generics
Create generic constructor functions that ensure proper initialization across different struct types.

This format will indeed allow for easy linking to specific sections, which will be very helpful as the document grows. Each of these headers can serve as an anchor point for further elaboration on the respective solution. Would you like me to expand on any of these solutions in particular, or should we move on to implementing one or more of them?

[[braveside - client needs conditional#Summary|top]]

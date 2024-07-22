### Utilize Go's zero values
Design the Company struct so its zero value is usable without further initialization, possibly using a NullDepartment implementation.

This solution addresses the nil interface problem by:

1. Safe Default State: Ensures that a Company is always in a valid state, even when not explicitly initialized.
2. Simplicity: Reduces the need for complex initialization logic.
3. Consistency: Provides consistent behavior across all instances of Company.
4. Robustness: Prevents nil pointer dereferences by always having a valid Department.
5. Idiomatic Go: Aligns with Go's philosophy of making zero values useful.


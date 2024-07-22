### Use the Null Object Pattern
Provide a "null" Department implementation instead of nil.

This solution addresses the nil interface problem by:

1. Avoiding Nil Checks: Eliminates the need for nil checks in client code.
2. Consistent Behavior: Provides default behavior for "empty" or "null" cases.
3. Simplification: Simplifies client code by removing special cases for null.
4. Polymorphism: Allows null objects to be used interchangeably with regular objects.
5. Robustness: Prevents nil pointer dereferences by always having a valid object.


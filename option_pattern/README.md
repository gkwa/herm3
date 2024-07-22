### Implement the Option Pattern
Use functional options to configure the Company struct during initialization, including setting a Department implementation.

This solution addresses the nil interface problem by:

1. Flexible Initialization: Allows for optional configuration of the Company struct, including its Department.
2. Default Values: Can provide default values for fields not explicitly set.
3. Clear API: Provides a clear and extensible API for object creation.
4. Encapsulation: Keeps the internal structure of Company private while allowing controlled modification.
5. Type Safety: Ensures type safety at compile-time for options.


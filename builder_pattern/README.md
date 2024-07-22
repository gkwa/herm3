### Employ the Builder Pattern
Create a separate builder struct to construct a fully initialized Company with a valid Department implementation.

This solution addresses the nil interface problem by:

1. Step-by-Step Construction: Allows for clear, step-by-step construction of complex objects.
2. Validation: Can perform validation at each step or before final object creation.
3. Immutable Objects: Can create immutable objects once construction is complete.
4. Flexible Creation Process: Allows for a flexible object creation process with optional components.
5. Separation of Concerns: Separates the construction logic from the object's behavior.


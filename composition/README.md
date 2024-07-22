### Use composition over embedding
Restructure the relationship between Company and Department to avoid interface fields.

This solution addresses the nil interface problem by:

1. Explicit Relationships: Composition makes the relationship between Company and Department more explicit and controllable.

2. Avoiding Interface Fields: By not using interface fields directly, we eliminate the possibility of nil interface panics.

3. Increased Flexibility: Composition allows for more flexible structures and easier testing and mocking.

4. Improved Encapsulation: The internal details of how Company manages its Department are hidden from outside code.

5. Better Separation of Concerns: Each struct (Company and Department) can focus on its own responsibilities without tight coupling.

By using composition, we create a more robust and flexible design that naturally avoids the pitfalls of nil interfaces while promoting better code organization and maintainability.


### Employ the New() pattern
Use a New() function to create and return properly initialized Company structs with a valid Department implementation.

This solution addresses the nil interface problem by:

1. Standardized Initialization: The New() function provides a clear, standard way to create properly initialized Company instances.

2. Encapsulation of Creation Logic: All the logic for creating a valid Company is contained within the New() function.

3. Flexibility: The New() function can handle different initialization scenarios, such as providing default values or validating inputs.

4. Clear API: It provides a clear and consistent API for object creation, improving code readability and maintainability.

5. Error Handling: The New() function can return an error along with the created object, allowing for more robust error handling during initialization.

By employing the New() pattern, we ensure that Company objects are always created in a valid state, preventing nil interface issues while providing a flexible and clear API for object creation.


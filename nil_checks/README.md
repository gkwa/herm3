### Add nil checks
Include nil checks in methods before accessing the Department interface.

This solution addresses the nil interface problem by:

1. Defensive Programming: Adding nil checks prevents the code from attempting to use a nil interface.

2. Graceful Error Handling: Instead of panicking, the code can handle nil cases more gracefully, such as by returning an error or using a default value.

3. Improved Robustness: The code becomes more resilient to unexpected nil values, reducing the likelihood of runtime panics.

4. Clear Error Messages: Custom error messages can be provided when a nil interface is encountered, aiding in debugging and error reporting.

5. Flexibility: Nil checks allow for different behaviors to be implemented when a nil interface is encountered, depending on the specific requirements of the application.

By adding nil checks, we create a more robust and fault-tolerant system that can handle nil interfaces gracefully, preventing panics and providing clearer error information.


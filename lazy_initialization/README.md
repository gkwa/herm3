### Lazy initialization
Initialize the Department implementation only when it's first accessed.

This solution addresses the nil interface problem by:

1. On-Demand Initialization: Creates the Department only when needed, saving resources.
2. Avoiding Nil Checks: Ensures a Department is always available when accessed.
3. Simplified API: Allows for a simpler initial object creation.
4. Performance Optimization: Can improve performance by delaying expensive initializations.
5. Thread Safety: Can be implemented in a thread-safe manner for concurrent access.


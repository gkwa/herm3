### Implement custom getter methods
Create methods that safely return Department information, handling nil interface cases internally.

This solution addresses the nil interface problem by:

1. Encapsulation: Hides the complexity of nil checks from the caller.
2. Consistent Behavior: Provides a consistent way to access Department information.
3. Centralized Logic: Centralizes nil-handling logic in one place.
4. Improved API: Offers a more user-friendly API that doesn't expose nil interfaces.
5. Flexibility: Allows for custom behavior when Department is nil.


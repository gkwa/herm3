### Implement default values
Provide a default Department implementation when creating a Company if one isn't specified.

This solution addresses the nil interface problem by:

1. Fallback Mechanism: The `NewCompany` function now checks if the provided `Department` is nil and uses a default implementation if so.

2. Always Valid State: This ensures that a `Company` always has a valid `Department`, even if none is explicitly provided.

3. Flexibility: Callers can still provide custom `Department` implementations when needed, but aren't required to.

4. Fail-Safe Behavior: Instead of panicking on a nil `Department`, the code now has a predictable default behavior.

5. Ease of Use: This approach simplifies the API for cases where a specific `Department` isn't important.

By implementing default values, we provide a safety net that prevents nil interface panics while still allowing for customization. This approach maintains the robustness of the constructor pattern while adding an extra layer of convenience and safety.

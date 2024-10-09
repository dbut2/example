# Contributing to example

Thank you for your interest in contributing to the `example` package! This document outlines the specific guidelines and considerations for contributing to this project.

## Understanding the Project

Before contributing, please ensure you understand the core purpose of the `example` package:

- It's designed to make complex logic accessible to non-technical users.
- It uses markdown tables for input and output.
- The goal is to simplify testing and demonstration of functions that operate on tabular data.

## Areas for Contribution

1. **New Example Implementations**: Add new examples in the `/examples` directory that demonstrate different use cases for the package. Focus on scenarios that would be valuable for non-technical users to interact with.
2. **Improving Type Support**: Enhance the package's ability to handle different data types in markdown tables. Consider common data types that might be useful in business logic or data processing scenarios.
3. **Input Parsing Enhancements**: Improve the parsing of markdown tables to handle more complex input scenarios, such as nested data structures or multi-line cell content.
4. **Output Formatting**: Develop new ways to format the output for better readability, especially for complex data structures or large datasets.
5. **Error Handling and User Feedback**: Improve error messages and feedback to be more user-friendly and informative for non-technical users.

## Development Guidelines

When contributing to the `example` package, please keep the following in mind:

1. **User-Centric Design**: Always consider the non-technical end-user. Any new feature or change should be accessible and understandable to users without programming experience.
2. **Markdown Table Compatibility**: Ensure that your contributions maintain or enhance the package's ability to work with markdown tables as the primary interface.
3. **Type Safety and Flexibility**: While adding new features or improving existing ones, maintain a balance between type safety and the flexibility needed to handle various input scenarios.
4. **Performance Considerations**: Given that the package uses reflection, be mindful of performance implications, especially when handling large datasets.
5. **Testing and Documentation via Examples**: Leverage Go's example tests to provide both documentation and testing. Follow the pattern of existing examples in the project.

## Testing and Documentation

The `example` package benefits from Go's built-in testing features in a unique way:

1. **Example Tests**: When you add new examples to the `/examples` directory, they serve a dual purpose:
    - They provide clear documentation of expected behavior for developers implementing the package.
    - They automatically become part of the test suite when running `go test ./...`.
2. **Ensuring Compatibility**: By following the established pattern for examples, you're not only documenting the expected behavior but also ensuring that the package remains compatible with these examples over time.
3. **Writing Effective Examples**: When creating examples:
    - Use realistic and meaningful input data that demonstrates practical use cases.
    - Ensure the example accurately reflects what the package will produce.
    - Consider edge cases and include examples that demonstrate how the package handles them.

Remember, these examples are likely to be the first point of reference for developers implementing the package, so make them clear, concise, and informative.

## Documentation for Examples

For new examples, include clear comments or a separate README in the example directory explaining:

- The purpose of the example
- The expected input format
- What the output represents
- Any specific considerations for implementation

## Questions?

If you have any questions about contributing or need clarification on the project's goals and design philosophy, please open an issue for discussion.

Thank you for helping make complex logic more accessible with `example`!

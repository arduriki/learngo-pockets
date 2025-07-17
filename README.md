# Learn Go with pocket projects

## Hello Directory - Concepts Learned

The `hello` directory demonstrates several fundamental Go concepts through a polyglot "Hello World" application:

### 1. Command-Line Flags
- Using the `flag` package to parse command-line arguments
- Defining flags with `flag.StringVar()` to accept user input
- Setting default values for optional parameters
- Parsing flags with `flag.Parse()`

### 2. Custom Types
- Creating custom types with `type language string`
- Using type definitions to make code more expressive and type-safe

### 3. Maps and Data Structures
- Using maps to store key-value pairs (`phrasebook` map)
- Storing multilingual greetings with language codes as keys
- Supporting multiple languages: Greek, English, French, Hebrew, Urdu, Vietnamese, and Catalan

### 4. Error Handling
- Using the comma-ok idiom to check map key existence
- Gracefully handling unsupported languages with error messages
- Returning formatted error messages using `fmt.Sprintf()`

### 5. Testing
- Writing unit tests with the `testing` package
- Using table-driven tests with `map[string]testCase` for multiple scenarios
- Implementing subtests with `t.Run()` for better test organization
- Testing both happy paths and error cases
- Using Example tests to demonstrate expected output

### Usage Example
```bash
go run main.go -lang=fr  # Output: Bonjour le monde
go run main.go -lang=el  # Output: Χαίρετε Κόσμε
go run main.go           # Output: Hello world (default)
```

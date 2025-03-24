# HTTP Codes Library for Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JuniorVieira99/jr_httpcodes)
![GoDoc](https://pkg.go.dev/badge/github.com/JuniorVieira99/jr_httpcodes)
![Go Report Card](https://goreportcard.com/badge/github.com/JuniorVieira99/jr_httpcodes)
[![jr_httpcodes_tests](https://github.com/JuniorVieira99/jr_httpcodes/actions/workflows/tests_workflow.yaml/badge.svg)](https://github.com/JuniorVieira99/jr_httpcodes/actions/workflows/tests_workflow.yaml)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

A lightweight, type-safe Go library for working with HTTP status codes and HTTP methods. This library provides constants, validation utilities, and human-readable descriptions for all standard HTTP status codes and methods.

## Index

- [Installation](#installation)
- [Key Features](#key-features)
- [Quick Start](#quick-start)
- [Usage Examples](#usage-examples)
  - [Working with Status Codes](#working-with-status-codes)
  - [Working with HTTP Methods](#working-with-http-methods)
- [Documentation](#documentation)
- [API Reference](#api-reference)
  - [Status Code Functions](#status-code-functions)
  - [Method Functions](#method-functions)
- [Available Constants](#available-constants)
- [Thread Safety](#thread-safety)
- [Overall Use](#overall-use)
- [License](#license)

## Installation

```bash
go get github.com/JuniorVieira99/jr_httpcodes
```

## Key Features

- Type-safe HTTP status codes and methods
- Human-readable descriptions for all standard HTTP status codes
- Validation utilities for status codes and methods
- Thread-safe registration of custom codes and methods
- Utility functions for debugging and logging

## Quick Start

```go
import (
    "fmt"
    "github.com/JuniorVieira99/jr_httpcodes"
)

func main() {
    fmt.Println(codes.GetStatusInfo(codes.OK)) // Output: "Request succeeded and response contains requested data"
}
```

## Usage Examples

### Working with Status Codes

```go
import (
    "fmt"
    "github.com/JuniorVieira99/jr_httpcodes"
)

func main() {
    
    // Check status code category
    if codes.IsSuccess(codes.OK) {
        fmt.Println("This is a success status code")
    }
    
    // Get description
    fmt.Println(codes.GetStatusInfo(codes.NotFound))
    // Output: "Requested resource could not be found"
    
    // Validate a status code
    err := codes.ValidateStatusCode(codes.StatusCode(999))
    if err != nil {
        fmt.Println("Invalid status code")
    }
    
    // String representation
    fmt.Println(codes.InternalServerError.String())
    // Output: "500 -> Server encountered unexpected condition"
    
    // Register a custom status code
    codes.RegisterStatusCode(codes.StatusCode(599), codes.Description("My Custom Error"))
}
```

### Working with HTTP Methods

```go
import (
    "fmt"
    "github.com/JuniorVieira99/jr_httpcodes"
)

func main() {
    // Get method description
    fmt.Println(codes.GetMethodDescription(codes.POST))
    // Output: "Send data to server for processing"
    
    // Validate a method
    err := codes.ValidateMethod(codes.Method("INVALID"))
    if err != nil {
        fmt.Println("Invalid method")
    }
    
    // String representation
    fmt.Println(codes.GET.String())
    // Output: "GET -> Retrieve data from server"
    
    // Register a custom method
    codes.RegisterMethod(codes.Method("CUSTOM"), codes.Description("My Custom Method"))
}
```

## Documentation

For local documentation check the `docs`folder.
For online documentation check the [GoDoc](https://pkg.go.dev/github.com/JuniorVieira99/jr_httpcodes).

## API Reference

### Status Code Functions

| Function | Description |
|----------|-------------|
| `IsValidStatusCode(code StatusCode) bool` | Checks if a status code is valid (100-599) |
| `IsInformational(code StatusCode) bool` | Checks if a code is informational (1xx) |
| `IsSuccess(code StatusCode) bool` | Checks if a code indicates success (2xx) |
| `IsRedirection(code StatusCode) bool` | Checks if a code indicates redirection (3xx) |
| `IsClientError(code StatusCode) bool` | Checks if a code indicates client error (4xx) |
| `IsServerError(code StatusCode) bool` | Checks if a code indicates server error (5xx) |
| `ValidateStatusCode(code StatusCode) error` | Returns error for invalid status codes |
| `GetStatusInfo(code StatusCode) string` | Returns human-readable description |
| `RegisterStatusCode(code StatusCode, desc Description)` | Registers a custom status code |
| `String() string` | Returns human-readable representation |
| `Print() string` | Prints the status code to the console |
| `CallMap() map[StatusCode]Description` | A map of status codes to status descriptions |
| `StringStatusCodeMap() string` | Returns a string representation of the status code map |
| `PrintStatusCodeMap()` | Prints the status code map to the console |

### Method Functions

| Function | Description |
|----------|-------------|
| `ValidateMethod(method Method) error` | Returns error for invalid methods |
| `GetMethodDescription(method Method) string` | Returns human-readable description |
| `RegisterMethod(method Method, desc Description)` | Registers a custom method |
| `String() string` | Returns human-readable representation |
| `Print() string` | Prints the method to the console |
| `CallMap() map[Method]Description` | A map of method names to method functions |
| `StringMethodMap() string` | Returns a string representation of the method map |
| `PrintMethodMap()` | Prints the method map to the console |

## Available Constants

The library includes constants for all standard HTTP status codes (100-511) and methods (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, CONNECT, TRACE).

**Note**: Check docs for detail information.

## Thread Safety

All registration functions are thread-safe and can be called from multiple goroutines.

## Tests

To run tests, use the following command:

```bash
go test ./tests
```

## Overall Use

```go
import (
    "fmt"
    "github.com/JuniorVieira99/jr_httpcodes"
)
func main()
{
    // STATUS CODE
    // ----------------

    // Check status code category
    if codes.IsSuccess(codes.OK) {
        fmt.Println("This is a success status code")
    }

    // Get description
    fmt.Println(codes.GetStatusInfo(codes.NotFound))
    // Output: "Requested resource could not be found"
    
    // Validate a status code
    err := codes.ValidateStatusCode(codes.StatusCode(999))
    if err != nil {
        fmt.Println("Invalid status code")
    }
    
    // String representation
    fmt.Println(codes.InternalServerError.String())
    // Output: "500 -> Server encountered unexpected condition"
    
    // Print status code
    codes.InternalServerError.Print()
    // Output: "500 -> Server encountered unexpected condition"

    // Register a custom status code
    codes.RegisterStatusCode(codes.StatusCode(599), codes.Description("My Custom Error"))

    // METHOD
    // ----------------
    
    // Get method description
    fmt.Println(codes.GetMethodDescription(codes.POST))
    // Output: "Send data to server for processing"
    
    // Validate a method
    err := codes.ValidateMethod(codes.Method("INVALID"))
    if err != nil {
        fmt.Println("Invalid method")
    }
    
    // String representation
    fmt.Println(codes.GET.String())
    // Output: "GET -> Retrieve data from server"
    
    // Register a custom method
    codes.RegisterMethod(codes.Method("CUSTOM"), codes.Description("My Custom Method")) 
}
```

## License

MIT License

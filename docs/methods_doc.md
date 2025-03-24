# HTTP Methods

The `httpcodes` package provides a set of constants and functions for working with HTTP Methods.

## Index

- [HTTP Methods Type](#http-methods-type)
- [HTTP Methods Description Type](#http-methods-description-type)
- [HTTP Methods Constants](#http-methods-constants)
- [HTTP Methods Descriptions](#http-methods-descriptions)
- [HTTP Map](#http-map)
- [HTTP Methods Functions and Methods](#http-methods-functions-and-methods)
  - [ValidateMethod](#validatemethod)
  - [String](#string)
  - [Print](#print)
  - [CallMap](#callmap)
  - [GetMethodDescription](#getmethoddescription)
  - [RegisterMethod](#registermethod)
  - [DeleteRegisteredMethod](#deleteregisteredmethod)

## HTTP Methods Type

A type safe representation of HTTP methods.

```go
type Method string
```

## HTTP Methods Description Type

A type safe representation of HTTP method descriptions.

```go
type Description string
```

## HTTP Methods Constants

```go
const (
    GET     Method = "GET"
    POST    Method = "POST"
    PUT     Method = "PUT"
    DELETE  Method = "DELETE"
    HEAD    Method = "HEAD"
    OPTIONS Method = "OPTIONS"
    TRACE   Method = "TRACE"
    CONNECT Method = "CONNECT"
)
```

## HTTP Methods Descriptions

```go
const (
    GETDesc     Description = "Retrieve data from server"
    POSTDesc    Description = "Send data to server for processing"
    PUTDesc     Description = "Update data on server"
    DELETEDesc  Description = "Delete data from server"
    PATCHDesc   Description = "Partially update data on server"
    HEADDesc    Description = "Retrieve metadata from server"
    OPTIONSDesc Description = "Retrieve options from server"
    CONNECTDesc Description = "Connect to server"
    TRACEDesc   Description = "Trace route to server"
)
```

## HTTP Map

Use it for mapping methods to their descriptions.

```go
var MethodDescriptionMap = map[Method]Description{
    GET:     GETDesc,
    POST:    POSTDesc,
    PUT:     PUTDesc,
    DELETE:  DELETEDesc,
    PATCH:   PATCHDesc,
    HEAD:    HEADDesc,
    OPTIONS: OPTIONSDesc,
    CONNECT: CONNECTDesc,
    TRACE:   TRACEDesc,
}
```

## HTTP Methods Functions and Methods

### ValidateMethod

`ValidateMethod` validates the method and returns an error if it's invalid.

**Signature**:

```go
func ValidateMethod(method Method) error {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `method` | `Method` | The HTTP method to validate |

**Returns**:

| Type | Description |
|------|-------------|
| `error` | An error if the method is invalid |

**Example**:

```go
err := codes.ValidateMethod(codes.Method("INVALID"))
if err != nil {
    fmt.Println(err.Error())
}
```

### String

`String` returns a human-readable representation of the method.

**Signature**:

```go
func (method Method) String() string {...}
```

**Returns**:

| Type | Description |
|------|-------------|
| `string` | A human-readable representation of the method |

**Example**:

```go
fmt.Println(codes.GET.String())
```

**Ouput**:

```shell
GET -> Retrieve data from server
```

### Print

`Print` prints a string representation of the method to the console.

**Signature**:

```go
func (method Method) Print() {...}
```

**Example**:

```go
codes.GET.Print()
```

**Output**:

```shell
GET -> Retrieve data from server
```

### CallMap

`CallMap` returns a map of methods to their descriptions.

**Signature**:

```go
func (method Method) CallMap() map[Method]Description {...}
```

**Returns**:

| Type | Description |
|------|-------------|
| `map[Method]Description` | A map of all HTTP methods and their descriptions |

**Example**:

```go
methodMap := codes.GET.CallMap()
fmt.Println(methodMap[codes.POST])
```

**Output**:

```shell
Send data to server for processing
```

### GetMethodDescription

`GetMethodDescription` returns a human-readable description of the HTTP method.

**Signature**:

```go
func GetMethodDescription(method Method) string {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `method` | `Method` | The HTTP method to get description for |

**Returns**:

| Type | Description |
|------|-------------|
| `string` | Human-readable description of the HTTP method |

**Example**:

```go
desc := codes.GetMethodDescription(codes.DELETE)
fmt.Println(desc)
```

**Output**:

```shell
Delete data from server
```

### RegisterMethod

`RegisterMethod` adds a custom HTTP method to the package's map of methods.

**Signature**:

```go
func RegisterMethod(method Method, description Description) {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `method` | `Method` | The custom HTTP method to register |
| `description` | `Description` | Human-readable description of the method |

**Example**:

```go
codes.RegisterMethod(codes.Method("CUSTOM"), codes.Description("My Custom Method"))
fmt.Println(codes.GetMethodDescription(codes.Method("CUSTOM")))
```

**Output**:

```shell
My Custom Method
```

### DeleteRegisteredMethod

`DeleteRegisteredMethod` removes a custom HTTP method from the package's map of methods.

**Signature**:

```go
func DeleteRegisteredMethod(method Method) {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `method` | `Method` | The custom HTTP method to remove |

**Example**:

```go
codes.RegisterMethod(codes.Method("CUSTOM"), codes.Description("My Custom Method"))
fmt.Println(codes.GetMethodDescription(codes.Method("CUSTOM")))
```

# HTTP Status Codes

The `httpcodes` package provides a set of constants and functions for working with HTTP status codes.

## Index

- [HTTP Status Codes](#http-status-codes)
  - [Index](#index)
    - [StatusCode Type](#statuscode-type)
    - [Description Type](#description-type)
    - [HTTP Status Codes Constants](#http-status-codes-constants)
      - [Descriptions](#descriptions)
    - [Description Map](#description-map)
    - [StatusCode Functions and Methods](#statuscode-functions-and-methods)
      - [IsValidStatusCode](#isvalidstatuscode)
      - [IsInformational](#isinformational)
      - [IsSuccess](#issuccess)
      - [IsRedirection](#isredirection)
      - [IsClientError](#isclienterror)
      - [IsServerError](#isservererror)
      - [ValidateStatusCode](#validatestatuscode)
      - [GetStatusInfo](#getstatusinfo)
      - [String](#string)
      - [Print](#print)
      - [CallMap](#callmap)
      - [RegisterStatusCode](#registerstatuscode)
      - [DeleteStatusCode](#deletestatuscode)

## StatusCode Type

A type safe representation of HTTP status codes.

```go
type StatusCode int
```

## Description Type

A type safe representation of HTTP status code descriptions.

```go
type Description string
```

## HTTP Status Codes Constants

```go
// HTTP Status Codes
const (

// Informational 1xx

Continue           StatusCode = 100
SwitchingProtocols StatusCode = 101
Processing         StatusCode = 102

// Success 2xx

OK                   StatusCode = 200
Created              StatusCode = 201
Accepted             StatusCode = 202
NonAuthoritativeInfo StatusCode = 203
NoContent            StatusCode = 204
ResetContent         StatusCode = 205
PartialContent       StatusCode = 206

// Redirection 3xx

MultipleChoices   StatusCode = 300
MovedPermanently  StatusCode = 301
Found             StatusCode = 302
SeeOther          StatusCode = 303
NotModified       StatusCode = 304
UseProxy          StatusCode = 305
TemporaryRedirect StatusCode = 307
PermanentRedirect StatusCode = 308

// Bad Request 4xx

BadRequest                  StatusCode = 400
Unauthorized                StatusCode = 401
PaymentRequired             StatusCode = 402
Forbidden                   StatusCode = 403
NotFound                    StatusCode = 404
MethodNotAllowed            StatusCode = 405
NotAcceptable               StatusCode = 406
ProxyAuthRequired           StatusCode = 407
RequestTimeout              StatusCode = 408
Conflict                    StatusCode = 409
Gone                        StatusCode = 410
LengthRequired              StatusCode = 411
PreconditionFailed          StatusCode = 412
PayloadTooLarge             StatusCode = 413
URITooLong                  StatusCode = 414
UnsupportedMediaType        StatusCode = 415
RangeNotSatisfiable         StatusCode = 416
ExpectationFailed           StatusCode = 417
Teapot                      StatusCode = 418
UnprocessableEntity         StatusCode = 422
TooEarly                    StatusCode = 425
UpgradeRequired             StatusCode = 426
PreconditionRequired        StatusCode = 428
TooManyRequests             StatusCode = 429
RequestHeaderFieldsTooLarge StatusCode = 431
UnavailableForLegalReasons  StatusCode = 451

// Server Error 5xx

InternalServerError           StatusCode = 500
NotImplemented                StatusCode = 501
BadGateway                    StatusCode = 502
ServiceUnavailable            StatusCode = 503
GatewayTimeout                StatusCode = 504
HTTPVersionNotSupported       StatusCode = 505
VariantAlsoNegotiates         StatusCode = 506
InsufficientStorage           StatusCode = 507
LoopDetected                  StatusCode = 508
NotExtended                   StatusCode = 510
NetworkAuthenticationRequired StatusCode = 511
)
```

### Descriptions

```go
const (

// Informational 1xx
ContinueDesc           Description = "Request received, processing continues"
SwitchingProtocolsDesc Description = "Server is switching protocols"
ProcessingDesc         Description = "Server is processing the request"

// Success 2xx

OKDesc                   Description = "Request succeeded and response contains requested data"
CreatedDesc              Description = "Resource created successfully and location provided"
AcceptedDesc             Description = "Request accepted for processing but processing not completed"
NonAuthoritativeInfoDesc Description = "Response contains non-authoritative information"
NoContentDesc            Description = "Request succeeded but no content returned"
ResetContentDesc         Description = "Request succeeded, client should reset document view"
PartialContentDesc       Description = "Partial content delivered as per range request"

// Redirection 3xx

MultipleChoicesDesc   Description = "Multiple options for resource available"
MovedPermanentlyDesc  Description = "Resource moved permanently to new location"
FoundDesc             Description = "Resource temporarily found at different location"
SeeOtherDesc          Description = "Client should get resource from different URI"
NotModifiedDesc       Description = "Resource not modified since last request"
UseProxyDesc          Description = "Requested resource must be accessed through proxy"
TemporaryRedirectDesc Description = "Resource temporarily moved to different location"
PermanentRedirectDesc Description = "Resource permanently moved to different location"

// Bad Request 4xx

BadRequestDesc                  Description = "Server cannot process request due to client error"
UnauthorizedDesc                Description = "Authentication required for resource access"
PaymentRequiredDesc             Description = "Payment required before processing request"
ForbiddenDesc                   Description = "Server refuses to fulfill request despite authentication"
NotFoundDesc                    Description = "Requested resource could not be found"
MethodNotAllowedDesc            Description = "Request method not supported for this resource"
NotAcceptableDesc               Description = "Resource cannot generate acceptable response"
ProxyAuthRequiredDesc           Description = "Authentication with proxy required"
RequestTimeoutDesc              Description = "Server timed out waiting for request"
ConflictDesc                    Description = "Request conflicts with current state of resource"
GoneDesc                        Description = "Resource permanently removed with no forwarding address"
LengthRequiredDesc              Description = "Content-Length header required for request"
PreconditionFailedDesc          Description = "Server precondition check failed"
PayloadTooLargeDesc             Description = "Request payload larger than server willing to process"
URITooLongDesc                  Description = "Request URI too long for server to process"
UnsupportedMediaTypeDesc        Description = "Media format not supported by server"
RangeNotSatisfiableDesc         Description = "Requested range cannot be satisfied"
ExpectationFailedDesc           Description = "Server cannot meet client expectation"
TeapotDesc                      Description = "I'm a teapot - RFC 2324 April Fools' joke"
UnprocessableEntityDesc         Description = "Request well-formed but semantically invalid"
TooEarlyDesc                    Description = "Server unwilling to risk processing due to replay attack"
UpgradeRequiredDesc             Description = "Client must switch to different protocol"
PreconditionRequiredDesc        Description = "Resource access requires conditional request"
TooManyRequestsDesc             Description = "Too many requests in given time period"
RequestHeaderFieldsTooLargeDesc Description = "Header fields too large for server to process"
UnavailableForLegalReasonsDesc  Description = "Resource access denied for legal reasons"

// Server Error 5xx

InternalServerErrorDesc           Description = "Server encountered unexpected condition"
NotImplementedDesc                Description = "Server does not support functionality required"
BadGatewayDesc                    Description = "Invalid response received from upstream server"
ServiceUnavailableDesc            Description = "Server temporarily unavailable"
GatewayTimeoutDesc                Description = "Upstream server failed to respond in time"
HTTPVersionNotSupportedDesc       Description = "HTTP version not supported by server"
VariantAlsoNegotiatesDesc         Description = "Server configuration error with transparent content negotiation"
InsufficientStorageDesc           Description = "Server unable to store resource to complete request"
LoopDetectedDesc                  Description = "Server detected infinite loop while processing request"
NotExtendedDesc                   Description = "Further extensions required to fulfill request"
NetworkAuthenticationRequiredDesc Description = "Client must authenticate to gain network access"
)
```

## Description Map

Use it for mapping status codes to their descriptions.

```go
var StatusDescriptionMap = map[StatusCode]Description{
// 1xx Informational
Continue:           ContinueDesc,
SwitchingProtocols: SwitchingProtocolsDesc,
Processing:         ProcessingDesc,

// 2xx Success
OK:                   OKDesc,
Created:              CreatedDesc,
Accepted:             AcceptedDesc,
NonAuthoritativeInfo: NonAuthoritativeInfoDesc,
NoContent:            NoContentDesc,
ResetContent:         ResetContentDesc,
PartialContent:       PartialContentDesc,

// 3xx Redirection
MultipleChoices:   MultipleChoicesDesc,
MovedPermanently:  MovedPermanentlyDesc,
Found:             FoundDesc,
SeeOther:          SeeOtherDesc,
NotModified:       NotModifiedDesc,
UseProxy:          UseProxyDesc,
TemporaryRedirect: TemporaryRedirectDesc,
PermanentRedirect: PermanentRedirectDesc,

// 4xx Client Errors
BadRequest:                  BadRequestDesc,
Unauthorized:                UnauthorizedDesc,
PaymentRequired:             PaymentRequiredDesc,
Forbidden:                   ForbiddenDesc,
NotFound:                    NotFoundDesc,
MethodNotAllowed:            MethodNotAllowedDesc,
NotAcceptable:               NotAcceptableDesc,
ProxyAuthRequired:           ProxyAuthRequiredDesc,
RequestTimeout:              RequestTimeoutDesc,
Conflict:                    ConflictDesc,
Gone:                        GoneDesc,
LengthRequired:              LengthRequiredDesc,
PreconditionFailed:          PreconditionFailedDesc,
PayloadTooLarge:             PayloadTooLargeDesc,
URITooLong:                  URITooLongDesc,
UnsupportedMediaType:        UnsupportedMediaTypeDesc,
RangeNotSatisfiable:         RangeNotSatisfiableDesc,
ExpectationFailed:           ExpectationFailedDesc,
Teapot:                      TeapotDesc,
UnprocessableEntity:         UnprocessableEntityDesc,
TooEarly:                    TooEarlyDesc,
UpgradeRequired:             UpgradeRequiredDesc,
PreconditionRequired:        PreconditionRequiredDesc,
TooManyRequests:             TooManyRequestsDesc,
RequestHeaderFieldsTooLarge: RequestHeaderFieldsTooLargeDesc,
UnavailableForLegalReasons:  UnavailableForLegalReasonsDesc,

// 5xx Server Errors
InternalServerError:           InternalServerErrorDesc,
NotImplemented:                NotImplementedDesc,
BadGateway:                    BadGatewayDesc,
ServiceUnavailable:            ServiceUnavailableDesc,
GatewayTimeout:                GatewayTimeoutDesc,
HTTPVersionNotSupported:       HTTPVersionNotSupportedDesc,
VariantAlsoNegotiates:         VariantAlsoNegotiatesDesc,
InsufficientStorage:           InsufficientStorageDesc,
LoopDetected:                  LoopDetectedDesc,
NotExtended:                   NotExtendedDesc,
NetworkAuthenticationRequired: NetworkAuthenticationRequiredDesc,
}
```

## StatusCode Functions and Methods

### IsValidStatusCode

`IsValidStatusCode` checks if the provided status code is within valid HTTP status code range.

**Signature**:

```go
func IsValidStatusCode(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to validate |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code is valid (100-599), `false` otherwise |

**Example**:

```go
if codes.IsValidStatusCode(codes.StatusCode(600)) {
    fmt.Println("Valid status code")
} else {
    fmt.Println("Invalid status code")
}
```

**Output**:

```shell
Invalid status code
```

### IsInformational

`IsInformational` checks if the status code indicates an informational response (1xx).

**Signature**:

```go
func IsInformational(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to check |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code is informational (100-199), `false` otherwise |

**Example**:

```go
if codes.IsInformational(codes.Continue) {
    fmt.Println("This is an informational status code")
}
```

### IsSuccess

`IsSuccess` checks if the status code indicates a successful request (2xx).

**Signature**:

```go
func IsSuccess(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to check |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code indicates success (200-299), `false` otherwise |

**Example**:

```go
if codes.IsSuccess(codes.OK) {
    fmt.Println("Request was successful")
}
```

### IsRedirection

`IsRedirection` checks if the status code indicates a redirection (3xx).

**Signature**:

```go
func IsRedirection(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to check |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code indicates redirection (300-399), `false` otherwise |

**Example**:

```go
if codes.IsRedirection(codes.MovedPermanently) {
    fmt.Println("Resource has been redirected")
}
```

### IsClientError

`IsClientError` checks if the status code indicates a client error (4xx).

**Signature**:

```go
func IsClientError(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to check |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code indicates client error (400-499), `false` otherwise |

**Example**:

```go
if codes.IsClientError(codes.NotFound) {
    fmt.Println("Client error occurred")
}
```

### IsServerError

`IsServerError` checks if the status code indicates a server error (5xx).

**Signature**:

```go
func IsServerError(code StatusCode) bool {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to check |

**Returns**:

| Type | Description |
|------|-------------|
| `bool` | `true` if the code indicates server error (500-599), `false` otherwise |

**Example**:

```go
if codes.IsServerError(codes.InternalServerError) {
    fmt.Println("Server error occurred")
}
```

### ValidateStatusCode

`ValidateStatusCode` validates the status code and returns an error if it's invalid.

**Signature**:

```go
func ValidateStatusCode(code StatusCode) error {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The HTTP status code to validate |

**Returns**:

| Type | Description |
|------|-------------|
| `error` | An error if the status code is invalid, nil otherwise |

**Example**:

```go
err := codes.ValidateStatusCode(codes.StatusCode(999))
if err != nil {
    fmt.Println(err.Error())
}
```

### GetStatusInfo

`GetStatusInfo` returns a human-readable description of the status code.

**Signature**:

```go
func GetStatusInfo(sc StatusCode) string {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `sc` | `StatusCode` | The HTTP status code to get description for |

**Returns**:

| Type | Description |
|------|-------------|
| `string` | Human-readable description of the status code |

**Example**:

```go
info := codes.GetStatusInfo(codes.NotFound)
fmt.Println(info)
```

**Output**:

```shell
Requested resource could not be found
```

### String

`String` returns a string representation of the status code.

**Signature**:

```go
func (sc StatusCode) String() string {...}
```

**Returns**:

| Type | Description |
|------|-------------|
| `string` | String representation in the format "code -> description" |

**Example**:

```go
fmt.Println(codes.NotFound.String())
```

**Output**:

```shell
404 -> Requested resource could not be found
```

### Print

`Print` prints a string representation of the status code to the console.

**Signature**:

```go
func (sc StatusCode) Print() {...}
```

**Example**:

```go
codes.NotFound.Print()
```

**Output**:

```shell
404 -> Requested resource could not be found
```

### CallMap

`CallMap` returns a map of status codes to their descriptions.

**Signature**:

```go
func (sc StatusCode) CallMap() map[StatusCode]Description {...}
```

**Returns**:

| Type | Description |
|------|-------------|
| `map[StatusCode]Description` | A map of all HTTP status codes and their descriptions |

**Example**:

```go
statusMap := codes.OK.CallMap()
fmt.Println(statusMap[codes.NotFound])
```

**Output**:

```shell
Requested resource could not be found
```

### RegisterStatusCode

`RegisterStatusCode` adds a custom status code to the package's map of status codes.

**Signature**:

```go
func RegisterStatusCode(code StatusCode, desc Description) {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The custom status code to register |
| `desc` | `Description` | Human-readable description of the status code |

**Example**:

```go
codes.RegisterStatusCode(codes.StatusCode(599), codes.Description("My Custom Error"))
fmt.Println(codes.GetStatusInfo(codes.StatusCode(599)))
```

**Output**:

```shell
My Custom Error
```

### DeleteStatusCode

`DeleteStatusCode` removes a custom status code from the package's map of status codes.

**Signature**:

```go
func DeleteStatusCode(code StatusCode) {...}
```

**Arguments**:

| Name | Type | Description |
|------|------|-------------|
| `code` | `StatusCode` | The custom status code to remove |

**Example**:

```go
codes.DeleteStatusCode(codes.StatusCode(599))
fmt.Println(codes.GetStatusInfo(codes.StatusCode(599)))
```

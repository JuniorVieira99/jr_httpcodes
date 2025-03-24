// Package codes provides HTTP status codes, HTTP methods and utility functions
// for working with HTTP responses and requests.
//
// This package defines common HTTP status codes as constants with their
// human-readable descriptions and provides helper functions to check the category
// of a status code (informational, success, redirection, client error, server error).
//
// It also includes common HTTP methods (GET, POST, PUT, DELETE, ...) with their descriptions
// and validation utilities.
//
// Status Code Example:
//
//	if codes.IsSuccess(response.StatusCode) {
//	    // Process successful response
//	} else if codes.IsClientError(response.StatusCode) {
//	    fmt.Printf("Client error: %s\n", codes.GetStatusInfo(response.StatusCode))
//	}
//
// Method Example:
//
//	method := codes.POST
//	fmt.Println(codes.GetMethodDescription(method)) // Output: "Send data to server for processing"
//	if err := codes.ValidateMethod(method); err != nil {
//	    log.Fatalf("Invalid HTTP method: %v", err)
//	}
//
// The package is designed to provide type safety and descriptive constants
// for HTTP operations while maintaining compatibility with standard Go HTTP libraries.
package codes

import (
	"fmt"
	"strings"
	"sync"
)

// Package Mutex
// --------------------------------------------------------------------

var mu sync.RWMutex

// Types
// --------------------------------------------------------------------

// StatusCode represents an HTTP status code.
type StatusCode int

// Method represents an HTTP request method.
type Method string

// Description represents a human-readable description of an HTTP status code.
type Description string

// Status Codes Constants
// --------------------------------------------------------------------

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

// StatusCode Descriptions
// --------------------------------------------------------------------

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

// StatusDescriptionMap maps status codes to their descriptions.
//
// Example:
//
//	desc := DescriptionMap[OK]
//	fmt.Println(desc) // Output: "Request succeeded and response contains requested data"
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

// RegisterStatusCode adds a custom status code to the package's map of status codes.
// It takes a StatusCode and a Description as parameters and adds the code to the map.
// The function is intended for use by other packages that want to add status codes
// that are not part of the standard HTTP/1.1 specification.
func RegisterStatusCode(code StatusCode, desc Description) {
	mu.Lock()
	if _, exists := StatusDescriptionMap[code]; exists {
		mu.Unlock()
		return
	}
	StatusDescriptionMap[code] = desc
	mu.Unlock()
}

// DeleteStatusCode removes a custom status code from the package's map of status codes.
// It takes a StatusCode as a parameter and deletes the code from the map if it exists.
// The function is intended for use by other packages that want to remove status codes
// that are not part of the standard HTTP/1.1 specification.
func DeleteStatusCode(code StatusCode) {
	mu.Lock()
	if _, exists := StatusDescriptionMap[code]; !exists {
		mu.Unlock()
		return
	}
	delete(StatusDescriptionMap, code)
	mu.Unlock()
}

// Description String func
func (c Description) String() string {
	return string(c)
}

// StatusCode Funcs
// --------------------------------------------------------------------

// IsValidStatusCode checks if the provided status code is a valid HTTP status code.
// Valid status codes are within the standard ranges (100-599).
func IsValidStatusCode(code StatusCode) bool {
	return code >= 100 && code < 600
}

// IsInformational checks if the status code indicates an informational response (1xx).
func IsInformational(code StatusCode) bool {
	return code >= 100 && code < 200
}

// IsSuccess checks if the status code indicates a successful request (2xx).
func IsSuccess(code StatusCode) bool {
	return code >= 200 && code < 300
}

// IsRedirection checks if the status code indicates a redirection (3xx).
func IsRedirection(code StatusCode) bool {
	return code >= 300 && code < 400
}

// IsClientError checks if the status code indicates a client error (4xx).
func IsClientError(code StatusCode) bool {
	return code >= 400 && code < 500
}

// IsServerError checks if the status code indicates a server error (5xx).
func IsServerError(code StatusCode) bool {
	return code >= 500 && code < 600
}

// ValidateStatusCode validates the status code and returns an error if it's invalid.
func ValidateStatusCode(code StatusCode) error {
	if !IsValidStatusCode(code) {
		return fmt.Errorf("invalid status code: %d", code)
	}
	return nil
}

// GetStatusInfo returns a human-readable description of the status code.
func GetStatusInfo(sc StatusCode) string {
	if desc, exists := StatusDescriptionMap[sc]; exists {
		return string(desc)
	}
	return "Unknown Status Code"
}

// String returns a string representation of the status code.
func (sc StatusCode) String() string {
	return fmt.Sprintf("%d -> %s", int(sc), GetStatusInfo(sc))
}

// Print prints a string representation of the status code to the console.
func (sc StatusCode) Print() {
	fmt.Println(sc.String())
}

// CallMap returns a map of status codes to their descriptions.
func (sc StatusCode) CallMap() map[StatusCode]Description {
	return StatusDescriptionMap
}

// Method Constants
// --------------------------------------------------------------------

// Method is a type for HTTP request methods.
const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	CONNECT Method = "CONNECT"
	TRACE   Method = "TRACE"
)

// Method Descriptions
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

// MethodDescriptionMap maps HTTP methods to their descriptions.
//
// Map:
//   - GET: "Retrieve data from server"
//   - POST: "Send data to server for processing"
//   - PUT: "Update data on server"
//   - DELETE: "Delete data from server"
//   - PATCH: "Partially update data on server"
//   - HEAD: "Retrieve metadata from server"
//   - OPTIONS: "Retrieve options from server"
//   - CONNECT: "Connect to server"
//   - TRACE: "Trace route to server"
//
// Example:
//
//	desc := MethodDescriptionMap[GET]
//	fmt.Println(desc) // Output: "Retrieve data from server"
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

// RegisterMethod adds a custom HTTP method to the package's map of methods.
// It takes a Method and a Description as parameters and adds the method to the map.
// The function is intended for use by other packages that want to add HTTP methods
// that are not part of the standard HTTP/1.1 specification.
func RegisterMethod(method Method, description Description) {
	mu.Lock()
	if _, exists := MethodDescriptionMap[method]; exists {
		mu.Unlock()
		return
	}
	MethodDescriptionMap[method] = description
	mu.Unlock()
}

// DeleteMethod removes an HTTP method from the package's map of methods.
// It takes a Method as a parameter and deletes the method from the map
// if it exists. The function is thread-safe and can be called from multiple
// goroutines.
func DeleteMethod(method Method) {
	mu.Lock()
	if _, exists := MethodDescriptionMap[method]; !exists {
		mu.Unlock()
		return
	}
	delete(MethodDescriptionMap, method)
	mu.Unlock()
}

// GetMethodDescription returns a human-readable description of the HTTP method.
func GetMethodDescription(method Method) string {
	if desc, exists := MethodDescriptionMap[method]; exists {
		return string(desc)
	}
	return "Unknown Method"
}

// Method Funcs
// --------------------------------------------------------------------

// ValidateMethod validates the method and returns an error if it's invalid.
func ValidateMethod(method Method) error {
	_, ok := MethodDescriptionMap[method]
	if !ok {
		return fmt.Errorf("invalid method: %s", method)
	}
	return nil
}

// String returns a string representation of the method.
func (m Method) String() string {
	return fmt.Sprintf("%v -> %s", string(m), GetMethodDescription(m))
}

// Print prints a string representation of the method to the console.
func (m Method) Print() {
	fmt.Println(m.String())
}

// CallMap returns a map of methods to their descriptions.
func (m Method) CallMap() map[Method]Description {
	return MethodDescriptionMap
}

// Utils
// --------------------------------------------------------------------

// StringMap takes a map of StatusCode to Description and returns a string
// representation of it, with each key-value pair separated by a line break.
// Each key-value pair is formatted as "code -> description".
func StringStatusCodeMap(m map[StatusCode]Description) string {
	var sb strings.Builder
	sb.Grow(len(m) * 20)

	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%d -> %s\n", k, v))
	}
	return sb.String()
}

// PrintMap prints the map of status codes to their descriptions to the console.
//
// It takes a map of status codes to their descriptions and prints it to the console.
// The output is formatted as "status code -> description" with one entry per line.
//
// Example:
//
//	m := codes.StatusDescriptionMap
//	codes.PrintMap(m) // Output: "100 -> Continue ..."
func PrintStatusCodeMap(m map[StatusCode]Description) {
	fmt.Println(StringStatusCodeMap(m))
}

// StringMethodMap takes a map of Method to Description and returns a string
// representation of it, with each key-value pair separated by a line break.
// Each key-value pair is formatted as "method -> description".
func StringMethodMap(m map[Method]Description) string {
	var sb strings.Builder
	sb.Grow(len(m) * 20)

	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%s -> %s\n", k, v))
	}
	return sb.String()
}

// PrintMethodMap prints the map of methods to their descriptions to the console.
//
// It takes a map of methods to their descriptions and prints it to the console.
// The output is formatted as "method -> description" with one entry per line.
//
// Example:
//
//	m := codes.MethodDescriptionMap
//	codes.PrintMethodMap(m) // Output: "GET -> Retrieve data from server ..."
func PrintMethodMap(m map[Method]Description) {
	fmt.Println(StringMethodMap(m))
}

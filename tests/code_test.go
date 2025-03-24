package code_test

import (
	"jr_httpcodes/codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCodeValidation(t *testing.T) {
	tests := []struct {
		name      string
		code      codes.StatusCode
		isValid   bool
		category  string
		errExpect bool
	}{
		{"Valid OK", codes.OK, true, "success", false},
		{"Valid NotFound", codes.NotFound, true, "client_error", false},
		{"Valid InternalServerError", codes.InternalServerError, true, "server_error", false},
		{"Valid Continue", codes.Continue, true, "informational", false},
		{"Valid Found", codes.Found, true, "redirection", false},
		{"Invalid Low", codes.StatusCode(99), false, "invalid", true},
		{"Invalid High", codes.StatusCode(600), false, "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isValid, codes.IsValidStatusCode(tt.code))

			err := codes.ValidateStatusCode(tt.code)
			if tt.errExpect {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			switch tt.category {
			case "informational":
				assert.True(t, codes.IsInformational(tt.code))
			case "success":
				assert.True(t, codes.IsSuccess(tt.code))
			case "redirection":
				assert.True(t, codes.IsRedirection(tt.code))
			case "client_error":
				assert.True(t, codes.IsClientError(tt.code))
			case "server_error":
				assert.True(t, codes.IsServerError(tt.code))
			}
		})
	}
}

func TestStatusCodeDescriptions(t *testing.T) {
	tests := []struct {
		name        string
		code        codes.StatusCode
		description string
	}{
		{"OK Description", codes.OK, "Request succeeded and response contains requested data"},
		{"NotFound Description", codes.NotFound, "Requested resource could not be found"},
		{"Unknown Code", codes.StatusCode(999), "Unknown Status Code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.description, codes.GetStatusInfo(tt.code))
		})
	}
}

func TestStatusCodeMethods(t *testing.T) {
	// Test String method
	okString := codes.OK.String()
	assert.Contains(t, okString, "200")
	assert.Contains(t, okString, "Request succeeded")

	// Test CallMap method
	okMap := codes.OK.CallMap()
	assert.Equal(t, codes.StatusDescriptionMap, okMap)
}

func TestMethodValidation(t *testing.T) {
	tests := []struct {
		name      string
		method    codes.Method
		isValid   bool
		errExpect bool
	}{
		{"Valid GET", codes.GET, true, false},
		{"Valid POST", codes.POST, true, false},
		{"Valid PUT", codes.PUT, true, false},
		{"Valid DELETE", codes.DELETE, true, false},
		{"Invalid Method", codes.Method("INVALID"), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := codes.ValidateMethod(tt.method)
			if tt.errExpect {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMethodDescriptions(t *testing.T) {
	tests := []struct {
		name        string
		method      codes.Method
		description string
	}{
		{"GET Description", codes.GET, "Retrieve data from server"},
		{"POST Description", codes.POST, "Send data to server for processing"},
		{"Unknown Method", codes.Method("INVALID"), "Unknown Method"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.description, codes.GetMethodDescription(tt.method))
		})
	}
}

func TestMethodMethods(t *testing.T) {
	// Test String method
	getString := codes.GET.String()
	assert.Contains(t, getString, "GET")
	assert.Contains(t, getString, "Retrieve data")

	// Test CallMap method
	getMap := codes.GET.CallMap()
	assert.Equal(t, codes.MethodDescriptionMap, getMap)
}

func TestRegistrationFunctions(t *testing.T) {
	// Test RegisterStatusCode
	customCode := codes.StatusCode(599)
	customDesc := codes.Description("Custom status code")
	codes.RegisterStatusCode(customCode, customDesc)
	assert.Equal(t, string(customDesc), codes.GetStatusInfo(customCode))

	// Test RegisterMethod
	customMethod := codes.Method("CUSTOM")
	customMethodDesc := codes.Description("Custom method")
	codes.RegisterMethod(customMethod, customMethodDesc)
	assert.Equal(t, string(customMethodDesc), codes.GetMethodDescription(customMethod))
}

func TestUtilityFunctions(t *testing.T) {
	// Create a small test map
	testMap := map[codes.StatusCode]codes.Description{
		codes.OK:       codes.OKDesc,
		codes.NotFound: codes.NotFoundDesc,
	}

	// Test StringStatusCodeMap function
	mapStr := codes.StringStatusCodeMap(testMap)
	assert.Contains(t, mapStr, "200 ->")
	assert.Contains(t, mapStr, "404 ->")

	// StringStatusCodeMap is tested only for coverage as it prints to console
	codes.PrintStatusCodeMap(testMap)
}

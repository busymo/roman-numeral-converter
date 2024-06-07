package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConvertRange(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Routes
	r := gin.Default()
	r.GET("/convert", convertRange)

	// Create a test request to pass to our handler.
	req, err := http.NewRequest(http.MethodGet, "/convert?range=1-5", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Create a router to serve request
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `["I","II","III","IV","V"]`
	assert.Equal(t, expected, rr.Body.String())
}

func TestConvertRangeErrors(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Routes
	r := gin.Default()
	r.GET("/convert", convertRange)

	// Test case: Invalid range format
	req, _ := http.NewRequest(http.MethodGet, "/convert?range=1to5", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid input! Ensure that the input is in the format 'from-to'.")

	// Test case: 'from' is greater than 'to'
	req, _ = http.NewRequest(http.MethodGet, "/convert?range=5-1", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid input. Ensure that numbers are integers and between 1 to 3999. Also, 'from' is less than or equal to 'to'.")

	// Test case: 'from' or 'to' is not a number
	req, _ = http.NewRequest(http.MethodGet, "/convert?range=a-5", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid input. Ensure that numbers are integers and between 1 to 3999. Also, 'from' is less than or equal to 'to'.")
}

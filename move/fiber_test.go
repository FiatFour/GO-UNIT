package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoutes(t *testing.T) {
	app := setupFiber()

	// Define test cases
	tests := []struct {
		description  string
		requestBody  UserFiber
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  UserFiber{"jane.doe@example.com", "Jane Doe", 30},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid email",
			requestBody:  UserFiber{"invalid-email", "Jane Doe", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid fullname",
			requestBody:  UserFiber{"jane.doe@example.com", "12345", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid age",
			requestBody:  UserFiber{"jane.doe@example.com", "Jane Doe", -5},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
}

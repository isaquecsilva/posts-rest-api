package rest_errors

import (
	"net/http"
	"testing"
)

func TestNewBadRequestRestError(test *testing.T) {
	test.Run("Should create correct bad request rest error object", func(t *testing.T) {
		bad_request := NewBadRequestRestError("Nickname cannot be empty")

		if bad_request.Code != http.StatusBadRequest ||
			bad_request.Message != "Nickname cannot be empty" {
			t.FailNow()
		}
	})
}

func TestNewInternalServerRestError(test *testing.T) {
	test.Run("should create correct internal server rest error object", func(t *testing.T) {
		internal_error := NewInternalServerRestError("Database connection timeout")

		if internal_error.Code != http.StatusInternalServerError ||
			internal_error.Message != "Database connection timeout" {
			t.FailNow()
		}
	})
}

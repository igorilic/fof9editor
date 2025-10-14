// ABOUTME: Tests for core validation types and functions
// ABOUTME: Verifies ValidationResult and field validation behavior

package validation

import "testing"

func TestNewValidationResult(t *testing.T) {
	result := NewValidationResult()

	if !result.Valid {
		t.Error("Expected new ValidationResult to be valid")
	}

	if len(result.Errors) != 0 {
		t.Error("Expected new ValidationResult to have no errors")
	}
}

func TestAddError(t *testing.T) {
	result := NewValidationResult()

	result.AddError("TestField", "test error message")

	if result.Valid {
		t.Error("Expected ValidationResult to be invalid after adding error")
	}

	if len(result.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.Errors))
	}

	if result.Errors[0].Field != "TestField" {
		t.Errorf("Expected field 'TestField', got '%s'", result.Errors[0].Field)
	}

	if result.Errors[0].Message != "test error message" {
		t.Errorf("Expected message 'test error message', got '%s'", result.Errors[0].Message)
	}
}

func TestHasError(t *testing.T) {
	result := NewValidationResult()

	if result.HasError("TestField") {
		t.Error("Expected no error for 'TestField'")
	}

	result.AddError("TestField", "test error")

	if !result.HasError("TestField") {
		t.Error("Expected error for 'TestField'")
	}

	if result.HasError("OtherField") {
		t.Error("Expected no error for 'OtherField'")
	}
}

func TestGetError(t *testing.T) {
	result := NewValidationResult()

	msg := result.GetError("TestField")
	if msg != "" {
		t.Errorf("Expected empty message, got '%s'", msg)
	}

	result.AddError("TestField", "test error message")

	msg = result.GetError("TestField")
	if msg != "test error message" {
		t.Errorf("Expected 'test error message', got '%s'", msg)
	}
}

func TestMerge(t *testing.T) {
	result1 := NewValidationResult()
	result2 := NewValidationResult()

	result1.AddError("Field1", "error 1")
	result2.AddError("Field2", "error 2")

	result1.Merge(result2)

	if result1.Valid {
		t.Error("Expected merged result to be invalid")
	}

	if len(result1.Errors) != 2 {
		t.Errorf("Expected 2 errors after merge, got %d", len(result1.Errors))
	}
}

func TestValidationError(t *testing.T) {
	err := ValidationError{
		Field:   "TestField",
		Message: "test message",
	}

	expected := "TestField: test message"
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

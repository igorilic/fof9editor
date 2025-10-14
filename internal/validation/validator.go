// ABOUTME: Core validation types and interfaces for FOF9 Editor
// ABOUTME: Provides field validation with error messages and validation rules

package validation

import "fmt"

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationResult holds the result of a validation operation
type ValidationResult struct {
	Valid  bool
	Errors []ValidationError
}

// NewValidationResult creates a new empty validation result
func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		Valid:  true,
		Errors: make([]ValidationError, 0),
	}
}

// AddError adds a validation error and marks the result as invalid
func (r *ValidationResult) AddError(field, message string) {
	r.Valid = false
	r.Errors = append(r.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// HasError checks if a specific field has an error
func (r *ValidationResult) HasError(field string) bool {
	for _, err := range r.Errors {
		if err.Field == field {
			return true
		}
	}
	return false
}

// GetError returns the error message for a specific field
func (r *ValidationResult) GetError(field string) string {
	for _, err := range r.Errors {
		if err.Field == field {
			return err.Message
		}
	}
	return ""
}

// Merge combines two validation results
func (r *ValidationResult) Merge(other *ValidationResult) {
	if !other.Valid {
		r.Valid = false
		r.Errors = append(r.Errors, other.Errors...)
	}
}

// FieldValidator is a function that validates a field value
type FieldValidator func(value interface{}) error

// ValidateField validates a single field with the given validators
func ValidateField(field string, value interface{}, validators ...FieldValidator) *ValidationResult {
	result := NewValidationResult()

	for _, validator := range validators {
		if err := validator(value); err != nil {
			result.AddError(field, err.Error())
		}
	}

	return result
}

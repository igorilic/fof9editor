// ABOUTME: Common validation rules for field values
// ABOUTME: Provides reusable validators for strings, integers, and ranges

package validation

import (
	"fmt"
	"strings"
)

// Required validates that a string is not empty
func Required(message string) FieldValidator {
	return func(value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid type for required validation")
		}
		if strings.TrimSpace(str) == "" {
			if message == "" {
				return fmt.Errorf("this field is required")
			}
			return fmt.Errorf("%s", message)
		}
		return nil
	}
}

// MinLength validates minimum string length
func MinLength(min int) FieldValidator {
	return func(value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid type for min length validation")
		}
		if len(strings.TrimSpace(str)) < min {
			return fmt.Errorf("must be at least %d characters", min)
		}
		return nil
	}
}

// MaxLength validates maximum string length
func MaxLength(max int) FieldValidator {
	return func(value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid type for max length validation")
		}
		if len(str) > max {
			return fmt.Errorf("must be at most %d characters", max)
		}
		return nil
	}
}

// IntRange validates that an integer is within a range (inclusive)
func IntRange(min, max int) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for range validation")
		}
		if num < min || num > max {
			return fmt.Errorf("must be between %d and %d", min, max)
		}
		return nil
	}
}

// IntMin validates minimum integer value
func IntMin(min int) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for min validation")
		}
		if num < min {
			return fmt.Errorf("must be at least %d", min)
		}
		return nil
	}
}

// IntMax validates maximum integer value
func IntMax(max int) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for max validation")
		}
		if num > max {
			return fmt.Errorf("must be at most %d", max)
		}
		return nil
	}
}

// IntPositive validates that an integer is positive (> 0)
func IntPositive() FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for positive validation")
		}
		if num <= 0 {
			return fmt.Errorf("must be a positive number")
		}
		return nil
	}
}

// IntNonNegative validates that an integer is non-negative (>= 0)
func IntNonNegative() FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for non-negative validation")
		}
		if num < 0 {
			return fmt.Errorf("must be zero or greater")
		}
		return nil
	}
}

// YearRange validates that a year is reasonable for the game
func YearRange(minYear, maxYear int) FieldValidator {
	return func(value interface{}) error {
		year, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for year validation")
		}
		if year < minYear || year > maxYear {
			return fmt.Errorf("year must be between %d and %d", minYear, maxYear)
		}
		return nil
	}
}

// MonthRange validates month (1-12)
func MonthRange() FieldValidator {
	return IntRange(1, 12)
}

// DayRange validates day of month (1-31)
func DayRange() FieldValidator {
	return IntRange(1, 31)
}

// RGBValue validates RGB color value (0-255)
func RGBValue() FieldValidator {
	return IntRange(0, 255)
}

// OneOf validates that an integer is one of the allowed values
func OneOf(allowed ...int) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return fmt.Errorf("invalid type for one-of validation")
		}
		for _, a := range allowed {
			if num == a {
				return nil
			}
		}
		return fmt.Errorf("must be one of: %v", allowed)
	}
}

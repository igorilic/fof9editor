// ABOUTME: Tests for common validation rules
// ABOUTME: Verifies string, integer, and range validators

package validation

import "testing"

func TestRequired(t *testing.T) {
	validator := Required("Field is required")

	// Test empty string
	err := validator("")
	if err == nil {
		t.Error("Expected error for empty string")
	}

	// Test whitespace
	err = validator("   ")
	if err == nil {
		t.Error("Expected error for whitespace string")
	}

	// Test valid string
	err = validator("valid")
	if err != nil {
		t.Errorf("Expected no error for valid string, got %v", err)
	}
}

func TestMinLength(t *testing.T) {
	validator := MinLength(3)

	// Test too short
	err := validator("ab")
	if err == nil {
		t.Error("Expected error for string too short")
	}

	// Test exact length
	err = validator("abc")
	if err != nil {
		t.Errorf("Expected no error for exact length, got %v", err)
	}

	// Test longer
	err = validator("abcd")
	if err != nil {
		t.Errorf("Expected no error for longer string, got %v", err)
	}
}

func TestMaxLength(t *testing.T) {
	validator := MaxLength(5)

	// Test within limit
	err := validator("abc")
	if err != nil {
		t.Errorf("Expected no error for string within limit, got %v", err)
	}

	// Test exact limit
	err = validator("abcde")
	if err != nil {
		t.Errorf("Expected no error for exact length, got %v", err)
	}

	// Test too long
	err = validator("abcdef")
	if err == nil {
		t.Error("Expected error for string too long")
	}
}

func TestIntRange(t *testing.T) {
	validator := IntRange(10, 20)

	// Test below range
	err := validator(9)
	if err == nil {
		t.Error("Expected error for value below range")
	}

	// Test minimum
	err = validator(10)
	if err != nil {
		t.Errorf("Expected no error for minimum value, got %v", err)
	}

	// Test within range
	err = validator(15)
	if err != nil {
		t.Errorf("Expected no error for value within range, got %v", err)
	}

	// Test maximum
	err = validator(20)
	if err != nil {
		t.Errorf("Expected no error for maximum value, got %v", err)
	}

	// Test above range
	err = validator(21)
	if err == nil {
		t.Error("Expected error for value above range")
	}
}

func TestIntMin(t *testing.T) {
	validator := IntMin(10)

	// Test below minimum
	err := validator(9)
	if err == nil {
		t.Error("Expected error for value below minimum")
	}

	// Test minimum
	err = validator(10)
	if err != nil {
		t.Errorf("Expected no error for minimum value, got %v", err)
	}

	// Test above minimum
	err = validator(11)
	if err != nil {
		t.Errorf("Expected no error for value above minimum, got %v", err)
	}
}

func TestIntMax(t *testing.T) {
	validator := IntMax(10)

	// Test below maximum
	err := validator(9)
	if err != nil {
		t.Errorf("Expected no error for value below maximum, got %v", err)
	}

	// Test maximum
	err = validator(10)
	if err != nil {
		t.Errorf("Expected no error for maximum value, got %v", err)
	}

	// Test above maximum
	err = validator(11)
	if err == nil {
		t.Error("Expected error for value above maximum")
	}
}

func TestIntPositive(t *testing.T) {
	validator := IntPositive()

	// Test negative
	err := validator(-1)
	if err == nil {
		t.Error("Expected error for negative value")
	}

	// Test zero
	err = validator(0)
	if err == nil {
		t.Error("Expected error for zero")
	}

	// Test positive
	err = validator(1)
	if err != nil {
		t.Errorf("Expected no error for positive value, got %v", err)
	}
}

func TestIntNonNegative(t *testing.T) {
	validator := IntNonNegative()

	// Test negative
	err := validator(-1)
	if err == nil {
		t.Error("Expected error for negative value")
	}

	// Test zero
	err = validator(0)
	if err != nil {
		t.Errorf("Expected no error for zero, got %v", err)
	}

	// Test positive
	err = validator(1)
	if err != nil {
		t.Errorf("Expected no error for positive value, got %v", err)
	}
}

func TestYearRange(t *testing.T) {
	validator := YearRange(1920, 2100)

	// Test before range
	err := validator(1919)
	if err == nil {
		t.Error("Expected error for year before range")
	}

	// Test minimum
	err = validator(1920)
	if err != nil {
		t.Errorf("Expected no error for minimum year, got %v", err)
	}

	// Test within range
	err = validator(2000)
	if err != nil {
		t.Errorf("Expected no error for year within range, got %v", err)
	}

	// Test maximum
	err = validator(2100)
	if err != nil {
		t.Errorf("Expected no error for maximum year, got %v", err)
	}

	// Test after range
	err = validator(2101)
	if err == nil {
		t.Error("Expected error for year after range")
	}
}

func TestMonthRange(t *testing.T) {
	validator := MonthRange()

	// Test invalid month
	err := validator(0)
	if err == nil {
		t.Error("Expected error for month 0")
	}

	err = validator(13)
	if err == nil {
		t.Error("Expected error for month 13")
	}

	// Test valid months
	for month := 1; month <= 12; month++ {
		err = validator(month)
		if err != nil {
			t.Errorf("Expected no error for month %d, got %v", month, err)
		}
	}
}

func TestDayRange(t *testing.T) {
	validator := DayRange()

	// Test invalid day
	err := validator(0)
	if err == nil {
		t.Error("Expected error for day 0")
	}

	err = validator(32)
	if err == nil {
		t.Error("Expected error for day 32")
	}

	// Test valid days
	for day := 1; day <= 31; day++ {
		err = validator(day)
		if err != nil {
			t.Errorf("Expected no error for day %d, got %v", day, err)
		}
	}
}

func TestRGBValue(t *testing.T) {
	validator := RGBValue()

	// Test below range
	err := validator(-1)
	if err == nil {
		t.Error("Expected error for RGB value -1")
	}

	// Test minimum
	err = validator(0)
	if err != nil {
		t.Errorf("Expected no error for RGB 0, got %v", err)
	}

	// Test mid-range
	err = validator(128)
	if err != nil {
		t.Errorf("Expected no error for RGB 128, got %v", err)
	}

	// Test maximum
	err = validator(255)
	if err != nil {
		t.Errorf("Expected no error for RGB 255, got %v", err)
	}

	// Test above range
	err = validator(256)
	if err == nil {
		t.Error("Expected error for RGB value 256")
	}
}

func TestOneOf(t *testing.T) {
	validator := OneOf(1, 3, 5, 7)

	// Test invalid values
	err := validator(0)
	if err == nil {
		t.Error("Expected error for value not in list")
	}

	err = validator(2)
	if err == nil {
		t.Error("Expected error for value not in list")
	}

	// Test valid values
	for _, val := range []int{1, 3, 5, 7} {
		err = validator(val)
		if err != nil {
			t.Errorf("Expected no error for value %d, got %v", val, err)
		}
	}
}

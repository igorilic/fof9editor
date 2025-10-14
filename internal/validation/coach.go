// ABOUTME: Validation rules specific to Coach data
// ABOUTME: Validates coach fields according to FOF9 game constraints

package validation

import "github.com/igorilic/fof9editor/internal/models"

// ValidateCoach validates all fields of a coach
func ValidateCoach(coach *models.Coach) *ValidationResult {
	result := NewValidationResult()

	// Name validation
	result.Merge(ValidateField("FirstName", coach.FirstName,
		Required("First name is required"),
		MinLength(1),
		MaxLength(50),
	))

	result.Merge(ValidateField("LastName", coach.LastName,
		Required("Last name is required"),
		MinLength(1),
		MaxLength(50),
	))

	// Team validation (0-31 for 32 teams)
	result.Merge(ValidateField("Team", coach.Team,
		IntRange(0, 31),
	))

	// Position validation (0=HC, 1=OC, 2=DC, 3=ST, 4=S&C)
	result.Merge(ValidateField("Position", coach.Position,
		IntRange(0, 4),
	))

	// Position group validation (varies by position)
	result.Merge(ValidateField("PositionGroup", coach.PositionGroup,
		IntNonNegative(),
	))

	// Birth date validation
	if coach.BirthMonth != 0 {
		result.Merge(ValidateField("BirthMonth", coach.BirthMonth,
			MonthRange(),
		))
	}

	if coach.BirthDay != 0 {
		result.Merge(ValidateField("BirthDay", coach.BirthDay,
			DayRange(),
		))
	}

	if coach.BirthYear != 0 {
		result.Merge(ValidateField("BirthYear", coach.BirthYear,
			YearRange(1920, 2020),
		))
	}

	// Birth city (optional)
	if coach.BirthCity != "" {
		result.Merge(ValidateField("BirthCity", coach.BirthCity,
			MaxLength(50),
		))
	}

	// Birth city ID validation
	if coach.BirthCityID != 0 {
		result.Merge(ValidateField("BirthCityID", coach.BirthCityID,
			IntNonNegative(),
		))
	}

	// College (optional)
	if coach.College != "" {
		result.Merge(ValidateField("College", coach.College,
			MaxLength(50),
		))
	}

	// College ID validation
	if coach.CollegeID != 0 {
		result.Merge(ValidateField("CollegeID", coach.CollegeID,
			IntNonNegative(),
		))
	}

	// Coaching styles
	// Offensive style (0-6)
	result.Merge(ValidateField("OffensiveStyle", coach.OffensiveStyle,
		IntRange(0, 6),
	))

	// Defensive style (0-4)
	result.Merge(ValidateField("DefensiveStyle", coach.DefensiveStyle,
		IntRange(0, 4),
	))

	// Pay scale validation (0-9999 representing $0 to $99.99M in $10K increments)
	result.Merge(ValidateField("PayScale", coach.PayScale,
		IntRange(0, 9999),
	))

	return result
}

// ValidateCoachField validates a single coach field
func ValidateCoachField(fieldName string, value interface{}) *ValidationResult {
	result := NewValidationResult()

	switch fieldName {
	case "FirstName", "LastName":
		if str, ok := value.(string); ok {
			result.Merge(ValidateField(fieldName, str,
				Required(fieldName+" is required"),
				MinLength(1),
				MaxLength(50),
			))
		}

	case "Team":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 31)))
		}

	case "Position":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 4)))
		}

	case "PositionGroup":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntNonNegative()))
		}

	case "BirthMonth":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, MonthRange()))
		}

	case "BirthDay":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, DayRange()))
		}

	case "BirthYear":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, YearRange(1920, 2020)))
		}

	case "BirthCity", "College":
		if str, ok := value.(string); ok && str != "" {
			result.Merge(ValidateField(fieldName, str, MaxLength(50)))
		}

	case "BirthCityID", "CollegeID":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, IntNonNegative()))
		}

	case "OffensiveStyle":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 6)))
		}

	case "DefensiveStyle":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 4)))
		}

	case "PayScale":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 9999)))
		}
	}

	return result
}

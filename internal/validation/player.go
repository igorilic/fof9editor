// ABOUTME: Validation rules specific to Player data
// ABOUTME: Validates player fields according to FOF9 game constraints

package validation

import "github.com/igorilic/fof9editor/internal/models"

// ValidatePlayer validates all fields of a player
func ValidatePlayer(player *models.Player) *ValidationResult {
	result := NewValidationResult()

	// Name validation
	result.Merge(ValidateField("FirstName", player.FirstName,
		Required("First name is required"),
		MinLength(1),
		MaxLength(50),
	))

	result.Merge(ValidateField("LastName", player.LastName,
		Required("Last name is required"),
		MinLength(1),
		MaxLength(50),
	))

	// Team validation (0-31 for 32 teams)
	result.Merge(ValidateField("Team", player.Team,
		IntRange(0, 31),
	))

	// Position validation (varies by game, typically 0-21)
	result.Merge(ValidateField("Position", player.PositionKey,
		IntRange(0, 21),
	))

	// Uniform number (0-99)
	result.Merge(ValidateField("Uniform", player.Uniform,
		IntRange(0, 99),
	))

	// Overall rating (0-99)
	result.Merge(ValidateField("OverallRating", player.OverallRating,
		IntRange(0, 99),
	))

	// Physical attributes
	// Height in inches (60-90 inches = 5'0" to 7'6")
	result.Merge(ValidateField("Height", player.Height,
		IntRange(60, 90),
	))

	// Weight in pounds (150-400 lbs)
	result.Merge(ValidateField("Weight", player.Weight,
		IntRange(150, 400),
	))

	// Hand size (7-12 inches)
	if player.HandSize != 0 { // Optional field
		result.Merge(ValidateField("HandSize", player.HandSize,
			IntRange(7, 12),
		))
	}

	// Arm length (28-38 inches)
	if player.ArmLength != 0 { // Optional field
		result.Merge(ValidateField("ArmLength", player.ArmLength,
			IntRange(28, 38),
		))
	}

	// Experience (0-25 years)
	result.Merge(ValidateField("Experience", player.Experience,
		IntRange(0, 25),
	))

	// College (optional, max 50 chars)
	if player.College != "" {
		result.Merge(ValidateField("College", player.College,
			MaxLength(50),
		))
	}

	// Year entered league (1920-2100)
	if player.YearEntry != 0 {
		result.Merge(ValidateField("YearEntry", player.YearEntry,
			YearRange(1920, 2100),
		))
	}

	// Draft round (1-7 for NFL, 0 for undrafted)
	result.Merge(ValidateField("RoundDrafted", player.RoundDrafted,
		IntRange(0, 7),
	))

	// Draft selection (1-300)
	if player.SelectionDrafted != 0 {
		result.Merge(ValidateField("SelectionDrafted", player.SelectionDrafted,
			IntRange(1, 300),
		))
	}

	return result
}

// ValidatePlayerField validates a single player field
func ValidatePlayerField(fieldName string, value interface{}) *ValidationResult {
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

	case "Position", "PositionKey":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 21)))
		}

	case "Uniform":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 99)))
		}

	case "OverallRating":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 99)))
		}

	case "Height":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(60, 90)))
		}

	case "Weight":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(150, 400)))
		}

	case "HandSize":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, IntRange(7, 12)))
		}

	case "ArmLength":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, IntRange(28, 38)))
		}

	case "Experience":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 25)))
		}

	case "College":
		if str, ok := value.(string); ok && str != "" {
			result.Merge(ValidateField(fieldName, str, MaxLength(50)))
		}

	case "YearEntry":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, YearRange(1920, 2100)))
		}

	case "RoundDrafted":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 7)))
		}

	case "SelectionDrafted":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, IntRange(1, 300)))
		}
	}

	return result
}

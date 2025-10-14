// ABOUTME: Validation rules specific to Team data
// ABOUTME: Validates team fields according to FOF9 game constraints

package validation

import "github.com/igorilic/fof9editor/internal/models"

// ValidateTeam validates all fields of a team
func ValidateTeam(team *models.Team) *ValidationResult {
	result := NewValidationResult()

	// Team identity
	result.Merge(ValidateField("TeamName", team.TeamName,
		Required("Team name is required"),
		MinLength(1),
		MaxLength(50),
	))

	result.Merge(ValidateField("NickName", team.NickName,
		Required("Nickname is required"),
		MinLength(1),
		MaxLength(50),
	))

	result.Merge(ValidateField("Abbreviation", team.Abbreviation,
		Required("Abbreviation is required"),
		MinLength(2),
		MaxLength(5),
	))

	// Year validation
	result.Merge(ValidateField("Year", team.Year,
		YearRange(1920, 2100),
	))

	// Team ID (0-31 for 32 teams)
	result.Merge(ValidateField("TeamID", team.TeamID,
		IntRange(0, 31),
	))

	// League structure
	// Conference (0-1 for AFC/NFC)
	result.Merge(ValidateField("Conference", team.Conference,
		IntRange(0, 1),
	))

	// Division (0-3 for North/South/East/West)
	result.Merge(ValidateField("Division", team.Division,
		IntRange(0, 3),
	))

	// City ID validation
	result.Merge(ValidateField("City", team.City,
		IntNonNegative(),
	))

	// Team colors (RGB 0-255)
	result.Merge(ValidateField("PrimaryRed", team.PrimaryRed, RGBValue()))
	result.Merge(ValidateField("PrimaryGreen", team.PrimaryGreen, RGBValue()))
	result.Merge(ValidateField("PrimaryBlue", team.PrimaryBlue, RGBValue()))
	result.Merge(ValidateField("SecondaryRed", team.SecondaryRed, RGBValue()))
	result.Merge(ValidateField("SecondaryGreen", team.SecondaryGreen, RGBValue()))
	result.Merge(ValidateField("SecondaryBlue", team.SecondaryBlue, RGBValue()))

	// Stadium info
	// Roof type (0=outdoor, 1=dome, 2=retractable)
	result.Merge(ValidateField("Roof", team.Roof,
		IntRange(0, 2),
	))

	// Turf type (0=grass, 1=artificial, 2=hybrid)
	result.Merge(ValidateField("Turf", team.Turf,
		IntRange(0, 2),
	))

	// Year built (1900-2100)
	if team.Built != 0 {
		result.Merge(ValidateField("Built", team.Built,
			YearRange(1900, 2100),
		))
	}

	// Stadium capacity (1,000 - 200,000)
	result.Merge(ValidateField("Capacity", team.Capacity,
		IntRange(1000, 200000),
	))

	// Luxury boxes (0-500)
	result.Merge(ValidateField("Luxury", team.Luxury,
		IntRange(0, 500),
	))

	// Condition (1-10 scale)
	result.Merge(ValidateField("Condition", team.Condition,
		IntRange(1, 10),
	))

	// Financial data
	// Attendance (0 - capacity)
	if team.Attendance > team.Capacity {
		result.AddError("Attendance", "cannot exceed stadium capacity")
	} else {
		result.Merge(ValidateField("Attendance", team.Attendance,
			IntNonNegative(),
		))
	}

	// Support level (0-100)
	result.Merge(ValidateField("Support", team.Support,
		IntRange(0, 100),
	))

	// Future stadium validation (if plan is active)
	if team.Plan != 0 {
		if team.FutureName != "" {
			result.Merge(ValidateField("FutureName", team.FutureName,
				MaxLength(50),
			))
		}

		if team.FutureAbbr != "" {
			result.Merge(ValidateField("FutureAbbr", team.FutureAbbr,
				MaxLength(5),
			))
		}

		result.Merge(ValidateField("FutureRoof", team.FutureRoof,
			IntRange(0, 2),
		))

		result.Merge(ValidateField("FutureTurf", team.FutureTurf,
			IntRange(0, 2),
		))

		if team.FutureCap != 0 {
			result.Merge(ValidateField("FutureCap", team.FutureCap,
				IntRange(1000, 200000),
			))
		}

		result.Merge(ValidateField("FutureLuxury", team.FutureLuxury,
			IntRange(0, 500),
		))
	}

	return result
}

// ValidateTeamField validates a single team field
func ValidateTeamField(fieldName string, value interface{}) *ValidationResult {
	result := NewValidationResult()

	switch fieldName {
	case "TeamName", "NickName":
		if str, ok := value.(string); ok {
			result.Merge(ValidateField(fieldName, str,
				Required(fieldName+" is required"),
				MinLength(1),
				MaxLength(50),
			))
		}

	case "Abbreviation":
		if str, ok := value.(string); ok {
			result.Merge(ValidateField(fieldName, str,
				Required("Abbreviation is required"),
				MinLength(2),
				MaxLength(5),
			))
		}

	case "Year":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, YearRange(1920, 2100)))
		}

	case "TeamID":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 31)))
		}

	case "Conference":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 1)))
		}

	case "Division":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 3)))
		}

	case "City":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntNonNegative()))
		}

	case "PrimaryRed", "PrimaryGreen", "PrimaryBlue",
		"SecondaryRed", "SecondaryGreen", "SecondaryBlue":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, RGBValue()))
		}

	case "Roof", "FutureRoof":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 2)))
		}

	case "Turf", "FutureTurf":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 2)))
		}

	case "Built":
		if num, ok := value.(int); ok && num != 0 {
			result.Merge(ValidateField(fieldName, num, YearRange(1900, 2100)))
		}

	case "Capacity", "FutureCap":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(1000, 200000)))
		}

	case "Luxury", "FutureLuxury":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 500)))
		}

	case "Condition":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(1, 10)))
		}

	case "Attendance":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntNonNegative()))
		}

	case "Support":
		if num, ok := value.(int); ok {
			result.Merge(ValidateField(fieldName, num, IntRange(0, 100)))
		}

	case "FutureName":
		if str, ok := value.(string); ok && str != "" {
			result.Merge(ValidateField(fieldName, str, MaxLength(50)))
		}

	case "FutureAbbr":
		if str, ok := value.(string); ok && str != "" {
			result.Merge(ValidateField(fieldName, str, MaxLength(5)))
		}
	}

	return result
}

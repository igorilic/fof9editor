// ABOUTME: Coach and Team CSV loading functionality for FOF9 Editor
// ABOUTME: Maps CSV records to Coach and Team structs with proper type conversions

package data

import (
	"fmt"
	"reflect"

	"github.com/igorilic/fof9editor/internal/models"
)

// LoadCoaches reads a coach CSV file and returns a slice of Coach structs
func LoadCoaches(filepath string) ([]models.Coach, error) {
	reader := NewCSVReader(filepath)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read coach CSV: %w", err)
	}

	coaches := make([]models.Coach, 0, len(records))
	for i, record := range records {
		coach, err := mapRowToCoach(record)
		if err != nil {
			return nil, fmt.Errorf("error parsing coach at row %d: %w", i+2, err)
		}
		coaches = append(coaches, coach)
	}

	return coaches, nil
}

// mapRowToCoach converts a CSV row (map of column->value) to a Coach struct
func mapRowToCoach(row map[string]string) (models.Coach, error) {
	coach := models.Coach{}
	coachValue := reflect.ValueOf(&coach).Elem()
	coachType := coachValue.Type()

	for i := 0; i < coachType.NumField(); i++ {
		field := coachType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		value, ok := row[csvTag]
		if !ok {
			// Column missing from CSV - use zero value
			continue
		}

		fieldValue := coachValue.Field(i)
		if err := setFieldValue(fieldValue, value, field.Name); err != nil {
			return coach, fmt.Errorf("field %s: %w", field.Name, err)
		}
	}

	return coach, nil
}

// LoadTeams reads a team CSV file and returns a slice of Team structs
func LoadTeams(filepath string) ([]models.Team, error) {
	reader := NewCSVReader(filepath)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read team CSV: %w", err)
	}

	teams := make([]models.Team, 0, len(records))
	for i, record := range records {
		team, err := mapRowToTeam(record)
		if err != nil {
			return nil, fmt.Errorf("error parsing team at row %d: %w", i+2, err)
		}
		teams = append(teams, team)
	}

	return teams, nil
}

// mapRowToTeam converts a CSV row (map of column->value) to a Team struct
func mapRowToTeam(row map[string]string) (models.Team, error) {
	team := models.Team{}
	teamValue := reflect.ValueOf(&team).Elem()
	teamType := teamValue.Type()

	for i := 0; i < teamType.NumField(); i++ {
		field := teamType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		value, ok := row[csvTag]
		if !ok {
			// Column missing from CSV - use zero value
			continue
		}

		fieldValue := teamValue.Field(i)
		if err := setFieldValue(fieldValue, value, field.Name); err != nil {
			return team, fmt.Errorf("field %s: %w", field.Name, err)
		}
	}

	return team, nil
}

// ABOUTME: Coach and Team CSV writing functionality for FOF9 Editor
// ABOUTME: Converts Coach and Team structs to CSV format and writes them to files

package data

import (
	"fmt"
	"reflect"

	"github.com/igorilic/fof9editor/internal/models"
)

// SaveCoaches writes a slice of Coach structs to a CSV file
func SaveCoaches(filepath string, coaches []models.Coach) error {
	if len(coaches) == 0 {
		// Write empty file with headers only
		headers := getCoachHeaders()
		writer := NewCSVWriter(filepath)
		return writer.WriteAll(headers, []map[string]string{})
	}

	// Get headers from coach struct
	headers := getCoachHeaders()

	// Convert coaches to records
	records := make([]map[string]string, 0, len(coaches))
	for i, coach := range coaches {
		record, err := coachToMap(coach)
		if err != nil {
			return fmt.Errorf("error converting coach %d to CSV: %w", i, err)
		}
		records = append(records, record)
	}

	// Write to file
	writer := NewCSVWriter(filepath)
	return writer.WriteAll(headers, records)
}

// getCoachHeaders returns the CSV headers for Coach struct
func getCoachHeaders() []string {
	coach := models.Coach{}
	coachType := reflect.TypeOf(coach)

	headers := make([]string, 0, coachType.NumField())
	for i := 0; i < coachType.NumField(); i++ {
		field := coachType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag != "" {
			headers = append(headers, csvTag)
		}
	}

	return headers
}

// coachToMap converts a Coach struct to a map[string]string for CSV writing
func coachToMap(coach models.Coach) (map[string]string, error) {
	record := make(map[string]string)
	coachValue := reflect.ValueOf(coach)
	coachType := coachValue.Type()

	for i := 0; i < coachType.NumField(); i++ {
		field := coachType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		fieldValue := coachValue.Field(i)
		strValue, err := fieldValueToString(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", field.Name, err)
		}

		record[csvTag] = strValue
	}

	return record, nil
}

// SaveTeams writes a slice of Team structs to a CSV file
func SaveTeams(filepath string, teams []models.Team) error {
	if len(teams) == 0 {
		// Write empty file with headers only
		headers := getTeamHeaders()
		writer := NewCSVWriter(filepath)
		return writer.WriteAll(headers, []map[string]string{})
	}

	// Get headers from team struct
	headers := getTeamHeaders()

	// Convert teams to records
	records := make([]map[string]string, 0, len(teams))
	for i, team := range teams {
		record, err := teamToMap(team)
		if err != nil {
			return fmt.Errorf("error converting team %d to CSV: %w", i, err)
		}
		records = append(records, record)
	}

	// Write to file
	writer := NewCSVWriter(filepath)
	return writer.WriteAll(headers, records)
}

// getTeamHeaders returns the CSV headers for Team struct
func getTeamHeaders() []string {
	team := models.Team{}
	teamType := reflect.TypeOf(team)

	headers := make([]string, 0, teamType.NumField())
	for i := 0; i < teamType.NumField(); i++ {
		field := teamType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag != "" {
			headers = append(headers, csvTag)
		}
	}

	return headers
}

// teamToMap converts a Team struct to a map[string]string for CSV writing
func teamToMap(team models.Team) (map[string]string, error) {
	record := make(map[string]string)
	teamValue := reflect.ValueOf(team)
	teamType := teamValue.Type()

	for i := 0; i < teamType.NumField(); i++ {
		field := teamType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		fieldValue := teamValue.Field(i)
		strValue, err := fieldValueToString(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", field.Name, err)
		}

		record[csvTag] = strValue
	}

	return record, nil
}

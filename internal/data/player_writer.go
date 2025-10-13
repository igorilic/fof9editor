// ABOUTME: Player CSV writing functionality for FOF9 Editor
// ABOUTME: Converts Player structs to CSV format and writes them to files

package data

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/igorilic/fof9editor/internal/models"
)

// SavePlayers writes a slice of Player structs to a CSV file
func SavePlayers(filepath string, players []models.Player) error {
	if len(players) == 0 {
		// Write empty file with headers only
		headers := getPlayerHeaders()
		writer := NewCSVWriter(filepath)
		return writer.WriteAll(headers, []map[string]string{})
	}

	// Get headers from first player struct
	headers := getPlayerHeaders()

	// Convert players to records
	records := make([]map[string]string, 0, len(players))
	for i, player := range players {
		record, err := playerToMap(player)
		if err != nil {
			return fmt.Errorf("error converting player %d to CSV: %w", i, err)
		}
		records = append(records, record)
	}

	// Write to file
	writer := NewCSVWriter(filepath)
	return writer.WriteAll(headers, records)
}

// getPlayerHeaders returns the CSV headers for Player struct
func getPlayerHeaders() []string {
	player := models.Player{}
	playerType := reflect.TypeOf(player)

	headers := make([]string, 0, playerType.NumField())
	for i := 0; i < playerType.NumField(); i++ {
		field := playerType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag != "" {
			headers = append(headers, csvTag)
		}
	}

	return headers
}

// playerToMap converts a Player struct to a map[string]string for CSV writing
func playerToMap(player models.Player) (map[string]string, error) {
	record := make(map[string]string)
	playerValue := reflect.ValueOf(player)
	playerType := playerValue.Type()

	for i := 0; i < playerType.NumField(); i++ {
		field := playerType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		fieldValue := playerValue.Field(i)
		strValue, err := fieldValueToString(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", field.Name, err)
		}

		record[csvTag] = strValue
	}

	return record, nil
}

// fieldValueToString converts a reflect.Value to its string representation
func fieldValueToString(value reflect.Value) (string, error) {
	switch value.Kind() {
	case reflect.Int:
		return strconv.FormatInt(value.Int(), 10), nil

	case reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64), nil

	case reflect.String:
		return value.String(), nil

	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil

	default:
		return "", fmt.Errorf("unsupported field type: %s", value.Kind())
	}
}

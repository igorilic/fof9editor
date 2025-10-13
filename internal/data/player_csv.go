// ABOUTME: Player CSV loading functionality for FOF9 Editor
// ABOUTME: Maps CSV records to Player structs with proper type conversions

package data

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/igorilic/fof9editor/internal/models"
)

// LoadPlayers reads a player CSV file and returns a slice of Player structs
func LoadPlayers(filepath string) ([]models.Player, error) {
	reader := NewCSVReader(filepath)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read player CSV: %w", err)
	}

	players := make([]models.Player, 0, len(records))
	for i, record := range records {
		player, err := mapRowToPlayer(record)
		if err != nil {
			return nil, fmt.Errorf("error parsing player at row %d: %w", i+2, err)
		}
		players = append(players, player)
	}

	return players, nil
}

// mapRowToPlayer converts a CSV row (map of column->value) to a Player struct
func mapRowToPlayer(row map[string]string) (models.Player, error) {
	player := models.Player{}
	playerValue := reflect.ValueOf(&player).Elem()
	playerType := playerValue.Type()

	for i := 0; i < playerType.NumField(); i++ {
		field := playerType.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag == "" {
			continue
		}

		value, ok := row[csvTag]
		if !ok {
			// Column missing from CSV - use zero value
			continue
		}

		fieldValue := playerValue.Field(i)
		if err := setFieldValue(fieldValue, value, field.Name); err != nil {
			return player, fmt.Errorf("field %s: %w", field.Name, err)
		}
	}

	return player, nil
}

// setFieldValue sets a struct field value from a string based on the field's type
func setFieldValue(field reflect.Value, value string, fieldName string) error {
	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", fieldName)
	}

	// Empty string handling
	if value == "" {
		// Leave as zero value
		return nil
	}

	switch field.Kind() {
	case reflect.Int:
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid integer value '%s': %w", value, err)
		}
		field.SetInt(int64(intVal))

	case reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid float value '%s': %w", value, err)
		}
		field.SetFloat(floatVal)

	case reflect.String:
		field.SetString(value)

	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean value '%s': %w", value, err)
		}
		field.SetBool(boolVal)

	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}

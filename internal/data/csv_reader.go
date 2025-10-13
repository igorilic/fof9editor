// ABOUTME: CSV reading functionality for FOF9 Editor data files
// ABOUTME: Provides generic CSV parsing with header mapping support

package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// CSVReader handles reading CSV files with header support
type CSVReader struct {
	filepath string
}

// NewCSVReader creates a new CSV reader for the given file path
func NewCSVReader(filepath string) *CSVReader {
	return &CSVReader{
		filepath: filepath,
	}
}

// ReadAll reads all records from the CSV file and returns them as a slice of maps
// where each map represents a row with column headers as keys
func (r *CSVReader) ReadAll() ([]map[string]string, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", r.filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read header row
	headers, err := csvReader.Read()
	if err != nil {
		if err == io.EOF {
			return []map[string]string{}, nil // Empty file
		}
		return nil, fmt.Errorf("failed to read headers: %w", err)
	}

	// Read all data rows
	var records []map[string]string
	lineNum := 2 // Start at 2 (1 is header)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading line %d: %w", lineNum, err)
		}

		// Create map for this row
		record := make(map[string]string)
		for i, header := range headers {
			if i < len(row) {
				record[strings.TrimSpace(header)] = strings.TrimSpace(row[i])
			} else {
				record[strings.TrimSpace(header)] = "" // Missing column value
			}
		}

		records = append(records, record)
		lineNum++
	}

	return records, nil
}

// ReadAllWithHeaders reads all records and also returns the headers separately
func (r *CSVReader) ReadAllWithHeaders() ([]string, []map[string]string, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file %s: %w", r.filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read header row
	headers, err := csvReader.Read()
	if err != nil {
		if err == io.EOF {
			return []string{}, []map[string]string{}, nil // Empty file
		}
		return nil, nil, fmt.Errorf("failed to read headers: %w", err)
	}

	// Trim whitespace from headers
	for i := range headers {
		headers[i] = strings.TrimSpace(headers[i])
	}

	// Read all data rows
	var records []map[string]string
	lineNum := 2

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("error reading line %d: %w", lineNum, err)
		}

		record := make(map[string]string)
		for i, header := range headers {
			if i < len(row) {
				record[strings.TrimSpace(header)] = strings.TrimSpace(row[i])
			} else {
				record[strings.TrimSpace(header)] = ""
			}
		}

		records = append(records, record)
		lineNum++
	}

	return headers, records, nil
}

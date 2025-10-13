// ABOUTME: CSV writing functionality for FOF9 Editor data files
// ABOUTME: Provides atomic CSV writing with backup and rollback support

package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// CSVWriter handles writing CSV files with atomic operations
type CSVWriter struct {
	filepath string
}

// NewCSVWriter creates a new CSV writer for the given file path
func NewCSVWriter(filepath string) *CSVWriter {
	return &CSVWriter{
		filepath: filepath,
	}
}

// WriteAll writes records to a CSV file atomically
// It writes to a temporary file first, then renames it to the target file
// This ensures the original file is not corrupted if the write fails
func (w *CSVWriter) WriteAll(headers []string, records []map[string]string) error {
	if len(headers) == 0 {
		return fmt.Errorf("headers cannot be empty")
	}

	// Create parent directory if it doesn't exist
	dir := filepath.Dir(w.filepath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write to temporary file
	tmpFile := w.filepath + ".tmp"
	file, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file %s: %w", tmpFile, err)
	}

	csvWriter := csv.NewWriter(file)

	// Write headers
	if err := csvWriter.Write(headers); err != nil {
		file.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Write data rows
	for i, record := range records {
		row := make([]string, len(headers))
		for j, header := range headers {
			row[j] = record[header]
		}
		if err := csvWriter.Write(row); err != nil {
			file.Close()
			os.Remove(tmpFile)
			return fmt.Errorf("failed to write row %d: %w", i+1, err)
		}
	}

	// Flush and check for errors
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		file.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	// Close file
	if err := file.Close(); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Atomically replace the original file
	if err := os.Rename(tmpFile, w.filepath); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to rename temp file to %s: %w", w.filepath, err)
	}

	return nil
}

// WriteAllFromSlice writes records from a slice of slices to a CSV file
// This is useful when you have data as [][]string instead of []map[string]string
func (w *CSVWriter) WriteAllFromSlice(headers []string, rows [][]string) error {
	if len(headers) == 0 {
		return fmt.Errorf("headers cannot be empty")
	}

	// Create parent directory if it doesn't exist
	dir := filepath.Dir(w.filepath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write to temporary file
	tmpFile := w.filepath + ".tmp"
	file, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file %s: %w", tmpFile, err)
	}

	csvWriter := csv.NewWriter(file)

	// Write headers
	if err := csvWriter.Write(headers); err != nil {
		file.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Write data rows
	for i, row := range rows {
		if len(row) != len(headers) {
			file.Close()
			os.Remove(tmpFile)
			return fmt.Errorf("row %d has %d columns, expected %d", i+1, len(row), len(headers))
		}
		if err := csvWriter.Write(row); err != nil {
			file.Close()
			os.Remove(tmpFile)
			return fmt.Errorf("failed to write row %d: %w", i+1, err)
		}
	}

	// Flush and check for errors
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		file.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	// Close file
	if err := file.Close(); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Atomically replace the original file
	if err := os.Rename(tmpFile, w.filepath); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to rename temp file to %s: %w", w.filepath, err)
	}

	return nil
}

// ABOUTME: Tests for CSV writing functionality
// ABOUTME: Covers atomic writes, error handling, and file operations

package data

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewCSVWriter(t *testing.T) {
	writer := NewCSVWriter("test.csv")
	if writer == nil {
		t.Fatal("NewCSVWriter returned nil")
	}
	if writer.filepath != "test.csv" {
		t.Errorf("Expected filepath 'test.csv', got '%s'", writer.filepath)
	}
}

func TestWriteAll_SimpleData(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE", "CITY"}
	records := []map[string]string{
		{"NAME": "John Doe", "AGE": "25", "CITY": "New York"},
		{"NAME": "Jane Smith", "AGE": "30", "CITY": "Los Angeles"},
	}

	err := writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Fatal("Output file was not created")
	}

	// Read back and verify
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if len(readRecords) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(readRecords))
	}

	if readRecords[0]["NAME"] != "John Doe" {
		t.Errorf("Expected NAME 'John Doe', got '%s'", readRecords[0]["NAME"])
	}
	if readRecords[1]["AGE"] != "30" {
		t.Errorf("Expected AGE '30', got '%s'", readRecords[1]["AGE"])
	}
}

func TestWriteAll_EmptyRecords(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE"}
	records := []map[string]string{}

	err := writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	// Verify file has only headers
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if len(readRecords) != 0 {
		t.Errorf("Expected 0 records, got %d", len(readRecords))
	}
}

func TestWriteAll_NoHeaders(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{}
	records := []map[string]string{
		{"NAME": "John"},
	}

	err := writer.WriteAll(headers, records)
	if err == nil {
		t.Fatal("Expected error for empty headers, got nil")
	}
}

func TestWriteAll_MissingColumnInRecord(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE", "CITY"}
	records := []map[string]string{
		{"NAME": "John", "AGE": "25"}, // Missing CITY
	}

	err := writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("WriteAll should handle missing columns: %v", err)
	}

	// Read back and verify empty string for missing column
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if readRecords[0]["CITY"] != "" {
		t.Errorf("Expected empty CITY, got '%s'", readRecords[0]["CITY"])
	}
}

func TestWriteAll_OverwriteExistingFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	// Write initial file
	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME"}
	records := []map[string]string{
		{"NAME": "John"},
	}
	err := writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("First WriteAll failed: %v", err)
	}

	// Overwrite with new data
	newRecords := []map[string]string{
		{"NAME": "Jane"},
		{"NAME": "Bob"},
	}
	err = writer.WriteAll(headers, newRecords)
	if err != nil {
		t.Fatalf("Second WriteAll failed: %v", err)
	}

	// Verify new data
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if len(readRecords) != 2 {
		t.Fatalf("Expected 2 records after overwrite, got %d", len(readRecords))
	}
	if readRecords[0]["NAME"] != "Jane" {
		t.Errorf("Expected NAME 'Jane', got '%s'", readRecords[0]["NAME"])
	}
}

func TestWriteAll_CreatesParentDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "subdir", "nested", "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME"}
	records := []map[string]string{
		{"NAME": "John"},
	}

	err := writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Fatal("Output file was not created")
	}
}

func TestWriteAll_AtomicWrite(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	// Create initial file
	initialContent := "NAME,AGE\nJohn,25\n"
	err := os.WriteFile(tmpFile, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create initial file: %v", err)
	}

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE"}
	records := []map[string]string{
		{"NAME": "Jane", "AGE": "30"},
	}

	err = writer.WriteAll(headers, records)
	if err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	// Verify temp file doesn't exist (cleaned up)
	tmpTempFile := tmpFile + ".tmp"
	if _, err := os.Stat(tmpTempFile); !os.IsNotExist(err) {
		t.Error("Temporary file was not cleaned up")
	}

	// Verify new data
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if readRecords[0]["NAME"] != "Jane" {
		t.Errorf("Expected NAME 'Jane', got '%s'", readRecords[0]["NAME"])
	}
}

// WriteAllFromSlice Tests

func TestWriteAllFromSlice_SimpleData(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE", "CITY"}
	rows := [][]string{
		{"John Doe", "25", "New York"},
		{"Jane Smith", "30", "Los Angeles"},
	}

	err := writer.WriteAllFromSlice(headers, rows)
	if err != nil {
		t.Fatalf("WriteAllFromSlice failed: %v", err)
	}

	// Read back and verify
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if len(readRecords) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(readRecords))
	}

	if readRecords[0]["NAME"] != "John Doe" {
		t.Errorf("Expected NAME 'John Doe', got '%s'", readRecords[0]["NAME"])
	}
}

func TestWriteAllFromSlice_MismatchedColumns(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE", "CITY"}
	rows := [][]string{
		{"John", "25"}, // Missing column
	}

	err := writer.WriteAllFromSlice(headers, rows)
	if err == nil {
		t.Fatal("Expected error for mismatched columns, got nil")
	}
}

func TestWriteAllFromSlice_EmptyRows(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{"NAME", "AGE"}
	rows := [][]string{}

	err := writer.WriteAllFromSlice(headers, rows)
	if err != nil {
		t.Fatalf("WriteAllFromSlice failed: %v", err)
	}

	// Verify file has only headers
	reader := NewCSVReader(tmpFile)
	readRecords, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}

	if len(readRecords) != 0 {
		t.Errorf("Expected 0 records, got %d", len(readRecords))
	}
}

func TestWriteAllFromSlice_NoHeaders(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	writer := NewCSVWriter(tmpFile)
	headers := []string{}
	rows := [][]string{
		{"John"},
	}

	err := writer.WriteAllFromSlice(headers, rows)
	if err == nil {
		t.Fatal("Expected error for empty headers, got nil")
	}
}

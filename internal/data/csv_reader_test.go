// ABOUTME: Tests for CSV reading functionality
// ABOUTME: Covers normal cases, empty files, errors, and edge cases

package data

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewCSVReader(t *testing.T) {
	reader := NewCSVReader("test.csv")
	if reader == nil {
		t.Fatal("NewCSVReader returned nil")
	}
	if reader.filepath != "test.csv" {
		t.Errorf("Expected filepath 'test.csv', got '%s'", reader.filepath)
	}
}

func TestReadAll_SimpleFile(t *testing.T) {
	reader := NewCSVReader("../../testdata/fixtures/csv/simple.csv")
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(records) != 3 {
		t.Fatalf("Expected 3 records, got %d", len(records))
	}

	// Check first record
	if records[0]["NAME"] != "John Doe" {
		t.Errorf("Expected NAME 'John Doe', got '%s'", records[0]["NAME"])
	}
	if records[0]["AGE"] != "25" {
		t.Errorf("Expected AGE '25', got '%s'", records[0]["AGE"])
	}
	if records[0]["CITY"] != "New York" {
		t.Errorf("Expected CITY 'New York', got '%s'", records[0]["CITY"])
	}

	// Check second record
	if records[1]["NAME"] != "Jane Smith" {
		t.Errorf("Expected NAME 'Jane Smith', got '%s'", records[1]["NAME"])
	}
	if records[1]["AGE"] != "30" {
		t.Errorf("Expected AGE '30', got '%s'", records[1]["AGE"])
	}

	// Check third record
	if records[2]["NAME"] != "Bob Johnson" {
		t.Errorf("Expected NAME 'Bob Johnson', got '%s'", records[2]["NAME"])
	}
}

func TestReadAll_EmptyFile(t *testing.T) {
	reader := NewCSVReader("../../testdata/fixtures/csv/empty.csv")
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records for empty file, got %d", len(records))
	}
}

func TestReadAll_MissingColumns(t *testing.T) {
	reader := NewCSVReader("../../testdata/fixtures/csv/missing_columns.csv")
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	// First record missing STATE column
	if records[0]["NAME"] != "John Doe" {
		t.Errorf("Expected NAME 'John Doe', got '%s'", records[0]["NAME"])
	}
	if records[0]["STATE"] != "" {
		t.Errorf("Expected empty STATE, got '%s'", records[0]["STATE"])
	}

	// Second record missing CITY and STATE columns
	if records[1]["NAME"] != "Jane Smith" {
		t.Errorf("Expected NAME 'Jane Smith', got '%s'", records[1]["NAME"])
	}
	if records[1]["CITY"] != "" {
		t.Errorf("Expected empty CITY, got '%s'", records[1]["CITY"])
	}
	if records[1]["STATE"] != "" {
		t.Errorf("Expected empty STATE, got '%s'", records[1]["STATE"])
	}
}

func TestReadAll_NonExistentFile(t *testing.T) {
	reader := NewCSVReader("nonexistent.csv")
	_, err := reader.ReadAll()
	if err == nil {
		t.Fatal("Expected error for non-existent file, got nil")
	}
}

func TestReadAll_InvalidPath(t *testing.T) {
	reader := NewCSVReader("")
	_, err := reader.ReadAll()
	if err == nil {
		t.Fatal("Expected error for empty path, got nil")
	}
}

func TestReadAllWithHeaders_SimpleFile(t *testing.T) {
	reader := NewCSVReader("../../testdata/fixtures/csv/simple.csv")
	headers, records, err := reader.ReadAllWithHeaders()
	if err != nil {
		t.Fatalf("ReadAllWithHeaders failed: %v", err)
	}

	// Check headers
	expectedHeaders := []string{"NAME", "AGE", "CITY"}
	if len(headers) != len(expectedHeaders) {
		t.Fatalf("Expected %d headers, got %d", len(expectedHeaders), len(headers))
	}
	for i, expected := range expectedHeaders {
		if headers[i] != expected {
			t.Errorf("Header %d: expected '%s', got '%s'", i, expected, headers[i])
		}
	}

	// Check records
	if len(records) != 3 {
		t.Fatalf("Expected 3 records, got %d", len(records))
	}

	if records[0]["NAME"] != "John Doe" {
		t.Errorf("Expected NAME 'John Doe', got '%s'", records[0]["NAME"])
	}
}

func TestReadAllWithHeaders_EmptyFile(t *testing.T) {
	reader := NewCSVReader("../../testdata/fixtures/csv/empty.csv")
	headers, records, err := reader.ReadAllWithHeaders()
	if err != nil {
		t.Fatalf("ReadAllWithHeaders failed: %v", err)
	}

	if len(headers) != 3 {
		t.Errorf("Expected 3 headers, got %d", len(headers))
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records, got %d", len(records))
	}
}

func TestReadAll_WithWhitespace(t *testing.T) {
	// Create temporary file with whitespace
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "whitespace.csv")

	content := "NAME,  AGE  ,CITY\n  John  ,  25  ,  New York  \n"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	reader := NewCSVReader(tmpFile)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	// TrimLeadingSpace should handle leading spaces in fields
	if records[0]["NAME"] != "John" {
		t.Errorf("Expected trimmed NAME 'John', got '%s'", records[0]["NAME"])
	}
}

func TestReadAll_EmptyCSVFile(t *testing.T) {
	// Create completely empty file (no headers)
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "completely_empty.csv")

	err := os.WriteFile(tmpFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	reader := NewCSVReader(tmpFile)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records for completely empty file, got %d", len(records))
	}
}

// ABOUTME: Tests for FormView component
// ABOUTME: Verifies form field creation, validation, and callbacks

package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestNewFormView(t *testing.T) {
	fv := NewFormView()

	if fv == nil {
		t.Fatal("NewFormView returned nil")
	}

	if fv.container == nil {
		t.Error("FormView container is nil")
	}

	if fv.fields == nil {
		t.Error("FormView fields map is nil")
	}

	if fv.fieldEntries == nil {
		t.Error("FormView fieldEntries map is nil")
	}

	if fv.fieldSelects == nil {
		t.Error("FormView fieldSelects map is nil")
	}
}

func TestFormView_GetContainer(t *testing.T) {
	fv := NewFormView()
	container := fv.GetContainer()

	if container == nil {
		t.Error("GetContainer returned nil")
	}
}

func TestFormView_SetFields(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()

	fields := []FieldDef{
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: "John"},
		{Name: "age", Label: "Age", Type: FieldTypeNumber, Value: "25"},
		{Name: "position", Label: "Position", Type: FieldTypeSelect, Value: "QB", Options: []string{"QB", "RB", "WR"}},
	}

	fv.SetFields(fields)

	// Verify fields were created
	if len(fv.fields) != 3 {
		t.Errorf("Expected 3 fields, got %d", len(fv.fields))
	}

	// Verify text field
	if _, ok := fv.fieldEntries["firstName"]; !ok {
		t.Error("firstName field not found in fieldEntries")
	}

	// Verify number field
	if _, ok := fv.fieldEntries["age"]; !ok {
		t.Error("age field not found in fieldEntries")
	}

	// Verify select field
	if _, ok := fv.fieldSelects["position"]; !ok {
		t.Error("position field not found in fieldSelects")
	}
}

func TestFormView_GetFieldValue(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()

	fields := []FieldDef{
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: "John"},
		{Name: "position", Label: "Position", Type: FieldTypeSelect, Value: "QB", Options: []string{"QB", "RB", "WR"}},
	}

	fv.SetFields(fields)

	// Test text field value
	if val := fv.GetFieldValue("firstName"); val != "John" {
		t.Errorf("Expected 'John', got '%s'", val)
	}

	// Test select field value
	if val := fv.GetFieldValue("position"); val != "QB" {
		t.Errorf("Expected 'QB', got '%s'", val)
	}

	// Test non-existent field
	if val := fv.GetFieldValue("nonexistent"); val != "" {
		t.Errorf("Expected empty string for non-existent field, got '%s'", val)
	}
}

func TestFormView_AddButtons(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()
	fv.AddButtons()

	if fv.buttonBar == nil {
		t.Error("Button bar is nil after AddButtons")
	}

	// Verify button bar has buttons
	if len(fv.buttonBar.Objects) == 0 {
		t.Error("Button bar has no objects")
	}
}

func TestFormView_SetCallbacks(t *testing.T) {
	fv := NewFormView()

	saveCalled := false
	deleteCalled := false
	nextCalled := false
	prevCalled := false

	fv.SetCallbacks(
		func() { saveCalled = true },
		func() { deleteCalled = true },
		func() { nextCalled = true },
		func() { prevCalled = true },
	)

	// Trigger callbacks
	if fv.onSave != nil {
		fv.onSave()
	}
	if fv.onDelete != nil {
		fv.onDelete()
	}
	if fv.onNext != nil {
		fv.onNext()
	}
	if fv.onPrev != nil {
		fv.onPrev()
	}

	if !saveCalled {
		t.Error("Save callback was not called")
	}
	if !deleteCalled {
		t.Error("Delete callback was not called")
	}
	if !nextCalled {
		t.Error("Next callback was not called")
	}
	if !prevCalled {
		t.Error("Prev callback was not called")
	}
}

func TestFormView_Clear(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()

	fields := []FieldDef{
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: "John"},
	}

	fv.SetFields(fields)
	fv.AddButtons()

	// Verify fields exist
	if len(fv.fields) == 0 {
		t.Error("Fields should exist before Clear")
	}

	fv.Clear()

	// Verify fields are cleared
	if len(fv.fields) != 0 {
		t.Errorf("Expected 0 fields after Clear, got %d", len(fv.fields))
	}

	if len(fv.fieldEntries) != 0 {
		t.Errorf("Expected 0 fieldEntries after Clear, got %d", len(fv.fieldEntries))
	}

	if len(fv.fieldSelects) != 0 {
		t.Errorf("Expected 0 fieldSelects after Clear, got %d", len(fv.fieldSelects))
	}
}

func TestFormView_NumberFieldValidation(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()

	fields := []FieldDef{
		{Name: "age", Label: "Age", Type: FieldTypeNumber, Value: "25"},
	}

	fv.SetFields(fields)

	entry := fv.fieldEntries["age"]
	if entry == nil {
		t.Fatal("Age field entry is nil")
	}

	// Initial value should be "25"
	if entry.Text != "25" {
		t.Errorf("Expected initial value '25', got '%s'", entry.Text)
	}

	// Test that entry accepts numeric input by setting text directly
	entry.SetText("30")
	if entry.Text != "30" {
		t.Errorf("Expected '30' after valid input, got '%s'", entry.Text)
	}
}

func TestFormView_MultipleFieldTypes(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()

	fields := []FieldDef{
		{Name: "name", Label: "Name", Type: FieldTypeText, Value: "Test"},
		{Name: "count", Label: "Count", Type: FieldTypeNumber, Value: "10"},
		{Name: "status", Label: "Status", Type: FieldTypeSelect, Value: "Active", Options: []string{"Active", "Inactive"}},
	}

	fv.SetFields(fields)

	// Verify all field types exist
	if len(fv.fieldEntries) != 2 { // text and number
		t.Errorf("Expected 2 entries, got %d", len(fv.fieldEntries))
	}

	if len(fv.fieldSelects) != 1 {
		t.Errorf("Expected 1 select, got %d", len(fv.fieldSelects))
	}

	// Verify values
	if fv.GetFieldValue("name") != "Test" {
		t.Error("Name value incorrect")
	}
	if fv.GetFieldValue("count") != "10" {
		t.Error("Count value incorrect")
	}
	if fv.GetFieldValue("status") != "Active" {
		t.Error("Status value incorrect")
	}
}

func TestFormView_ButtonCallbacks(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	fv := NewFormView()
	fv.AddButtons()

	callbackCounts := map[string]int{
		"save":   0,
		"delete": 0,
		"next":   0,
		"prev":   0,
	}

	fv.SetCallbacks(
		func() { callbackCounts["save"]++ },
		func() { callbackCounts["delete"]++ },
		func() { callbackCounts["next"]++ },
		func() { callbackCounts["prev"]++ },
	)

	// Find and tap buttons
	for _, obj := range fv.buttonBar.Objects {
		if btn, ok := obj.(*widget.Button); ok {
			switch btn.Text {
			case "< Previous":
				test.Tap(btn)
			case "Next >":
				test.Tap(btn)
			case "Save":
				test.Tap(btn)
			case "Delete":
				test.Tap(btn)
			}
		}
	}

	// Verify all callbacks were triggered
	for name, count := range callbackCounts {
		if count != 1 {
			t.Errorf("Expected %s callback to be called once, got %d", name, count)
		}
	}
}

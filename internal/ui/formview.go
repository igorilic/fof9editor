// ABOUTME: Form view component for FOF9 Editor
// ABOUTME: Provides record editing with fields, validation, and navigation

package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// FieldType represents the type of form field
type FieldType string

const (
	FieldTypeText   FieldType = "text"
	FieldTypeNumber FieldType = "number"
	FieldTypeSelect FieldType = "select"
)

// FieldDef defines a form field
type FieldDef struct {
	Name    string
	Label   string
	Type    FieldType
	Value   string
	Options []string // For select fields
}

// FormView represents a form for editing records
type FormView struct {
	container    *fyne.Container
	fields       map[string]fyne.CanvasObject
	fieldEntries map[string]*widget.Entry  // Store entries for value retrieval
	fieldSelects map[string]*widget.Select // Store selects for value retrieval
	onSave       func()
	onDelete     func()
	onNext       func()
	onPrev       func()
	buttonBar    *fyne.Container
}

// NewFormView creates a new form view
func NewFormView() *FormView {
	fv := &FormView{
		fields:       make(map[string]fyne.CanvasObject),
		fieldEntries: make(map[string]*widget.Entry),
		fieldSelects: make(map[string]*widget.Select),
	}

	fv.container = container.NewVBox()
	return fv
}

// SetFields configures the form with field definitions
func (fv *FormView) SetFields(fieldDefs []FieldDef) {
	// Clear existing fields
	fv.fields = make(map[string]fyne.CanvasObject)
	fv.fieldEntries = make(map[string]*widget.Entry)
	fv.fieldSelects = make(map[string]*widget.Select)

	// Create form content
	formContent := container.NewVBox()

	for _, def := range fieldDefs {
		label := widget.NewLabel(def.Label + ":")
		label.TextStyle = fyne.TextStyle{Bold: true}

		var fieldWidget fyne.CanvasObject

		switch def.Type {
		case FieldTypeText:
			entry := widget.NewEntry()
			entry.SetText(def.Value)
			fieldWidget = entry
			fv.fieldEntries[def.Name] = entry

		case FieldTypeNumber:
			entry := widget.NewEntry()
			entry.SetText(def.Value)
			// Add validation for numeric input
			entry.OnChanged = func(text string) {
				if text == "" {
					return
				}
				if _, err := strconv.Atoi(text); err != nil {
					// Revert to previous valid value if invalid
					if def.Value != "" {
						entry.SetText(def.Value)
					}
				}
			}
			fieldWidget = entry
			fv.fieldEntries[def.Name] = entry

		case FieldTypeSelect:
			sel := widget.NewSelect(def.Options, nil)
			sel.SetSelected(def.Value)
			fieldWidget = sel
			fv.fieldSelects[def.Name] = sel

		default:
			// Default to text entry
			entry := widget.NewEntry()
			entry.SetText(def.Value)
			fieldWidget = entry
			fv.fieldEntries[def.Name] = entry
		}

		fv.fields[def.Name] = fieldWidget

		// Add label and field to form
		fieldRow := container.NewBorder(nil, nil, label, nil, fieldWidget)
		formContent.Add(fieldRow)
	}

	// Rebuild container with form content and button bar
	fv.container.Objects = []fyne.CanvasObject{formContent}
	if fv.buttonBar != nil {
		fv.container.Add(fv.buttonBar)
	}
	fv.container.Refresh()
}

// AddButtons creates the button bar with navigation and action buttons
func (fv *FormView) AddButtons() {
	// Create buttons
	prevButton := widget.NewButton("< Previous", func() {
		if fv.onPrev != nil {
			fv.onPrev()
		}
	})

	nextButton := widget.NewButton("Next >", func() {
		if fv.onNext != nil {
			fv.onNext()
		}
	})

	saveButton := widget.NewButton("Save", func() {
		if fv.onSave != nil {
			fv.onSave()
		}
	})
	saveButton.Importance = widget.HighImportance

	deleteButton := widget.NewButton("Delete", func() {
		if fv.onDelete != nil {
			fv.onDelete()
		}
	})
	deleteButton.Importance = widget.DangerImportance

	// Create button bar layout
	fv.buttonBar = container.NewHBox(
		prevButton,
		nextButton,
		widget.NewSeparator(),
		saveButton,
		deleteButton,
	)

	// Add button bar to container if it doesn't already exist
	if len(fv.container.Objects) > 0 {
		fv.container.Add(fv.buttonBar)
	}
}

// GetContainer returns the form container
func (fv *FormView) GetContainer() *fyne.Container {
	return fv.container
}

// GetFieldValue returns the current value of a field
func (fv *FormView) GetFieldValue(fieldName string) string {
	// Check entries first
	if entry, ok := fv.fieldEntries[fieldName]; ok {
		return entry.Text
	}
	// Check selects
	if sel, ok := fv.fieldSelects[fieldName]; ok {
		return sel.Selected
	}
	return ""
}

// SetCallbacks sets the form callbacks
func (fv *FormView) SetCallbacks(onSave, onDelete, onNext, onPrev func()) {
	fv.onSave = onSave
	fv.onDelete = onDelete
	fv.onNext = onNext
	fv.onPrev = onPrev
}

// Clear removes all fields from the form
func (fv *FormView) Clear() {
	fv.fields = make(map[string]fyne.CanvasObject)
	fv.fieldEntries = make(map[string]*widget.Entry)
	fv.fieldSelects = make(map[string]*widget.Select)
	fv.container.Objects = []fyne.CanvasObject{}
	if fv.buttonBar != nil {
		fv.container.Add(fv.buttonBar)
	}
	fv.container.Refresh()
}

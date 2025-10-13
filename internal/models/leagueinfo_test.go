package models

import (
	"testing"
)

func TestNewDefaultLeagueInfo(t *testing.T) {
	baseYear := 2024
	info := NewDefaultLeagueInfo(baseYear)

	if info.BaseYear != baseYear {
		t.Errorf("Expected BaseYear %d, got %d", baseYear, info.BaseYear)
	}

	if info.ScheduleID != "32_8_17" {
		t.Errorf("Expected ScheduleID 32_8_17, got %s", info.ScheduleID)
	}

	if info.SalaryCap != 2000 {
		t.Errorf("Expected SalaryCap 2000, got %d", info.SalaryCap)
	}

	if info.Minimum != 70 {
		t.Errorf("Expected Minimum 70, got %d", info.Minimum)
	}

	if info.Salary10 != 180 {
		t.Errorf("Expected Salary10 180, got %d", info.Salary10)
	}
}

func TestGetSalaryMinimum(t *testing.T) {
	info := NewDefaultLeagueInfo(2024)

	tests := []struct {
		experience int
		expected   int
	}{
		{0, 70},   // Rookie
		{1, 85},   // 1 year
		{2, 100},  // 2 years
		{3, 115},  // 3 years
		{4, 130},  // 4 years
		{5, 130},  // 5 years
		{6, 70},   // 6 years (falls through to minimum)
		{7, 150},  // 7 years
		{8, 150},  // 8 years
		{9, 150},  // 9 years
		{10, 180}, // 10 years
		{15, 180}, // 15 years
		{20, 180}, // 20 years
	}

	for _, tt := range tests {
		actual := info.GetSalaryMinimum(tt.experience)
		if actual != tt.expected {
			t.Errorf("Experience %d: expected %d, got %d", tt.experience, tt.expected, actual)
		}
	}
}

func TestValidateScheduleID(t *testing.T) {
	tests := []struct {
		scheduleID string
		expected   bool
	}{
		{"32_8_17", true},
		{"16_4_16", true},
		{"24_6_16", true},
		{"invalid", false},
		{"32_8", false},
		{"32_8_17_extra", false},
		{"", false},
	}

	for _, tt := range tests {
		info := &LeagueInfo{ScheduleID: tt.scheduleID}
		actual := info.ValidateScheduleID()
		if actual != tt.expected {
			t.Errorf("ScheduleID %s: expected %v, got %v", tt.scheduleID, tt.expected, actual)
		}
	}
}

func TestLeagueInfoStructFields(t *testing.T) {
	info := &LeagueInfo{
		ScheduleID: "32_8_17",
		BaseYear:   2024,
		SalaryCap:  2500,
		Minimum:    80,
		Salary1:    90,
		Salary2:    110,
		Salary3:    125,
		Salary45:   140,
		Salary789:  160,
		Salary10:   190,
	}

	if info.ScheduleID != "32_8_17" {
		t.Errorf("Expected ScheduleID 32_8_17, got %s", info.ScheduleID)
	}

	if info.SalaryCap != 2500 {
		t.Errorf("Expected SalaryCap 2500, got %d", info.SalaryCap)
	}

	if !info.ValidateScheduleID() {
		t.Error("Expected ValidateScheduleID to return true")
	}

	// Test salary progression
	if info.Minimum >= info.Salary1 {
		t.Error("Expected Minimum < Salary1")
	}

	if info.Salary789 >= info.Salary10 {
		t.Error("Expected Salary789 < Salary10")
	}
}

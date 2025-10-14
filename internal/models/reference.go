// ABOUTME: Reference data container for FOF9 Editor
// ABOUTME: Holds position, team, and other lookup data for the application

package models

// ReferenceData contains all reference/lookup data for the application
type ReferenceData struct {
	Positions []Position
	Teams     []Team // Teams can serve as reference data for dropdowns
}

// NewReferenceData creates a new ReferenceData instance with default values
func NewReferenceData() *ReferenceData {
	return &ReferenceData{
		Positions: DefaultPositions(),
		Teams:     make([]Team, 0),
	}
}

// GetPositionOptions returns position names for dropdown selections
func (r *ReferenceData) GetPositionOptions() []string {
	options := make([]string, len(r.Positions))
	for i, pos := range r.Positions {
		options[i] = pos.Name
	}
	return options
}

// GetPositionIDByName returns the position ID for a given name
func (r *ReferenceData) GetPositionIDByName(name string) int {
	for _, pos := range r.Positions {
		if pos.Name == name {
			return pos.ID
		}
	}
	return -1
}

// GetTeamOptions returns team names for dropdown selections
func (r *ReferenceData) GetTeamOptions() []string {
	options := make([]string, len(r.Teams))
	for i, team := range r.Teams {
		options[i] = team.GetDisplayName()
	}
	return options
}

// GetTeamIDByName returns the team ID for a given display name
func (r *ReferenceData) GetTeamIDByName(displayName string) int {
	for _, team := range r.Teams {
		if team.GetDisplayName() == displayName {
			return team.TeamID
		}
	}
	return -1
}

// GetTeamNameByID returns the team display name for a given ID
func (r *ReferenceData) GetTeamNameByID(id int) string {
	for _, team := range r.Teams {
		if team.TeamID == id {
			return team.GetDisplayName()
		}
	}
	return "Unknown Team"
}

// GetCoachPositionOptions returns coach position names for dropdown selections
func (r *ReferenceData) GetCoachPositionOptions() []string {
	return []string{
		"Head Coach",
		"Offensive Coordinator",
		"Defensive Coordinator",
		"Special Teams Coordinator",
		"Strength & Conditioning",
	}
}

// GetCoachPositionIDByName returns the coach position ID for a given name
func (r *ReferenceData) GetCoachPositionIDByName(name string) int {
	positions := map[string]int{
		"Head Coach":                    0,
		"Offensive Coordinator":         1,
		"Defensive Coordinator":         2,
		"Special Teams Coordinator":     3,
		"Strength & Conditioning":       4,
	}
	if id, exists := positions[name]; exists {
		return id
	}
	return -1
}

// GetCoachPositionNameByID returns the coach position name for a given ID
func (r *ReferenceData) GetCoachPositionNameByID(id int) string {
	names := []string{
		"Head Coach",
		"Offensive Coordinator",
		"Defensive Coordinator",
		"Special Teams Coordinator",
		"Strength & Conditioning",
	}
	if id >= 0 && id < len(names) {
		return names[id]
	}
	return "Unknown Position"
}

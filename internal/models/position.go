// ABOUTME: Position reference data for FOF9 Editor
// ABOUTME: Defines football positions with IDs, names, and abbreviations

package models

// Position represents a football position
type Position struct {
	ID           int    `csv:"ID"`
	Name         string `csv:"NAME"`
	Abbreviation string `csv:"ABBREVIATION"`
	PositionType string `csv:"TYPE"` // Offense, Defense, Special Teams
}

// Common position constants (typical FOF9 position IDs)
const (
	PositionQB  = 0  // Quarterback
	PositionRB  = 1  // Running Back
	PositionFB  = 2  // Fullback
	PositionWR  = 3  // Wide Receiver
	PositionTE  = 4  // Tight End
	PositionLT  = 5  // Left Tackle
	PositionLG  = 6  // Left Guard
	PositionC   = 7  // Center
	PositionRG  = 8  // Right Guard
	PositionRT  = 9  // Right Tackle
	PositionLE  = 10 // Left End (DE)
	PositionRE  = 11 // Right End (DE)
	PositionDT  = 12 // Defensive Tackle
	PositionLOLB = 13 // Left Outside Linebacker
	PositionMLB = 14 // Middle Linebacker
	PositionROLB = 15 // Right Outside Linebacker
	PositionCB  = 16 // Cornerback
	PositionFS  = 17 // Free Safety
	PositionSS  = 18 // Strong Safety
	PositionK   = 19 // Kicker
	PositionP   = 20 // Punter
	PositionLS  = 21 // Long Snapper
)

// DefaultPositions returns the standard set of football positions
func DefaultPositions() []Position {
	return []Position{
		{ID: PositionQB, Name: "Quarterback", Abbreviation: "QB", PositionType: "Offense"},
		{ID: PositionRB, Name: "Running Back", Abbreviation: "RB", PositionType: "Offense"},
		{ID: PositionFB, Name: "Fullback", Abbreviation: "FB", PositionType: "Offense"},
		{ID: PositionWR, Name: "Wide Receiver", Abbreviation: "WR", PositionType: "Offense"},
		{ID: PositionTE, Name: "Tight End", Abbreviation: "TE", PositionType: "Offense"},
		{ID: PositionLT, Name: "Left Tackle", Abbreviation: "LT", PositionType: "Offense"},
		{ID: PositionLG, Name: "Left Guard", Abbreviation: "LG", PositionType: "Offense"},
		{ID: PositionC, Name: "Center", Abbreviation: "C", PositionType: "Offense"},
		{ID: PositionRG, Name: "Right Guard", Abbreviation: "RG", PositionType: "Offense"},
		{ID: PositionRT, Name: "Right Tackle", Abbreviation: "RT", PositionType: "Offense"},
		{ID: PositionLE, Name: "Left Defensive End", Abbreviation: "LE", PositionType: "Defense"},
		{ID: PositionRE, Name: "Right Defensive End", Abbreviation: "RE", PositionType: "Defense"},
		{ID: PositionDT, Name: "Defensive Tackle", Abbreviation: "DT", PositionType: "Defense"},
		{ID: PositionLOLB, Name: "Left Outside Linebacker", Abbreviation: "LOLB", PositionType: "Defense"},
		{ID: PositionMLB, Name: "Middle Linebacker", Abbreviation: "MLB", PositionType: "Defense"},
		{ID: PositionROLB, Name: "Right Outside Linebacker", Abbreviation: "ROLB", PositionType: "Defense"},
		{ID: PositionCB, Name: "Cornerback", Abbreviation: "CB", PositionType: "Defense"},
		{ID: PositionFS, Name: "Free Safety", Abbreviation: "FS", PositionType: "Defense"},
		{ID: PositionSS, Name: "Strong Safety", Abbreviation: "SS", PositionType: "Defense"},
		{ID: PositionK, Name: "Kicker", Abbreviation: "K", PositionType: "Special Teams"},
		{ID: PositionP, Name: "Punter", Abbreviation: "P", PositionType: "Special Teams"},
		{ID: PositionLS, Name: "Long Snapper", Abbreviation: "LS", PositionType: "Special Teams"},
	}
}

// GetPositionName returns the position name for a given ID
func GetPositionName(id int) string {
	positions := DefaultPositions()
	for _, pos := range positions {
		if pos.ID == id {
			return pos.Name
		}
	}
	return "Unknown"
}

// GetPositionAbbr returns the position abbreviation for a given ID
func GetPositionAbbr(id int) string {
	positions := DefaultPositions()
	for _, pos := range positions {
		if pos.ID == id {
			return pos.Abbreviation
		}
	}
	return "??"
}

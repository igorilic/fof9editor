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

// FOF9 position constants - matches the actual game's POSITION_KEY values
const (
	PositionQB   = 1  // Quarterback
	PositionRB   = 2  // Running Back
	PositionFB   = 3  // Fullback
	PositionTE   = 4  // Tight End
	PositionFL   = 5  // Flanker (WR)
	PositionSE   = 6  // Split End (WR)
	PositionLT   = 7  // Left Tackle
	PositionLG   = 8  // Left Guard
	PositionC    = 9  // Center
	PositionRG   = 10 // Right Guard
	PositionRT   = 11 // Right Tackle
	PositionP    = 12 // Punter
	PositionK    = 13 // Kicker
	PositionDLE  = 14 // Defensive Left End
	PositionDLT  = 15 // Defensive Left Tackle
	PositionDNT  = 16 // Defensive Nose Tackle
	PositionDRT  = 17 // Defensive Right Tackle
	PositionDRE  = 18 // Defensive Right End
	PositionSLB  = 19 // Strong-Side Linebacker
	PositionSILB = 20 // Strong Inside Linebacker
	PositionMLB  = 21 // Middle Linebacker
	PositionWILB = 22 // Weak Inside Linebacker
	PositionWLB  = 23 // Weak-Side Linebacker
	PositionLCB  = 24 // Left Cornerback
	PositionRCB  = 25 // Right Cornerback
	PositionSS   = 26 // Strong Safety
	PositionFS   = 27 // Free Safety
	PositionLS   = 28 // Long Snapper
)

// DefaultPositions returns the standard set of football positions matching FOF9
func DefaultPositions() []Position {
	return []Position{
		{ID: PositionQB, Name: "Quarterback", Abbreviation: "QB", PositionType: "Offense"},
		{ID: PositionRB, Name: "Running Back", Abbreviation: "RB", PositionType: "Offense"},
		{ID: PositionFB, Name: "Fullback", Abbreviation: "FB", PositionType: "Offense"},
		{ID: PositionTE, Name: "Tight End", Abbreviation: "TE", PositionType: "Offense"},
		{ID: PositionFL, Name: "Flanker", Abbreviation: "FL", PositionType: "Offense"},
		{ID: PositionSE, Name: "Split End", Abbreviation: "SE", PositionType: "Offense"},
		{ID: PositionLT, Name: "Left Tackle", Abbreviation: "LT", PositionType: "Offense"},
		{ID: PositionLG, Name: "Left Guard", Abbreviation: "LG", PositionType: "Offense"},
		{ID: PositionC, Name: "Center", Abbreviation: "C", PositionType: "Offense"},
		{ID: PositionRG, Name: "Right Guard", Abbreviation: "RG", PositionType: "Offense"},
		{ID: PositionRT, Name: "Right Tackle", Abbreviation: "RT", PositionType: "Offense"},
		{ID: PositionP, Name: "Punter", Abbreviation: "P", PositionType: "Special Teams"},
		{ID: PositionK, Name: "Kicker", Abbreviation: "K", PositionType: "Special Teams"},
		{ID: PositionDLE, Name: "Defensive Left End", Abbreviation: "DLE", PositionType: "Defense"},
		{ID: PositionDLT, Name: "Defensive Left Tackle", Abbreviation: "DLT", PositionType: "Defense"},
		{ID: PositionDNT, Name: "Defensive Nose Tackle", Abbreviation: "DNT", PositionType: "Defense"},
		{ID: PositionDRT, Name: "Defensive Right Tackle", Abbreviation: "DRT", PositionType: "Defense"},
		{ID: PositionDRE, Name: "Defensive Right End", Abbreviation: "DRE", PositionType: "Defense"},
		{ID: PositionSLB, Name: "Strong-Side Linebacker", Abbreviation: "SLB", PositionType: "Defense"},
		{ID: PositionSILB, Name: "Strong Inside Linebacker", Abbreviation: "SILB", PositionType: "Defense"},
		{ID: PositionMLB, Name: "Middle Linebacker", Abbreviation: "MLB", PositionType: "Defense"},
		{ID: PositionWILB, Name: "Weak Inside Linebacker", Abbreviation: "WILB", PositionType: "Defense"},
		{ID: PositionWLB, Name: "Weak-Side Linebacker", Abbreviation: "WLB", PositionType: "Defense"},
		{ID: PositionLCB, Name: "Left Cornerback", Abbreviation: "LCB", PositionType: "Defense"},
		{ID: PositionRCB, Name: "Right Cornerback", Abbreviation: "RCB", PositionType: "Defense"},
		{ID: PositionSS, Name: "Strong Safety", Abbreviation: "SS", PositionType: "Defense"},
		{ID: PositionFS, Name: "Free Safety", Abbreviation: "FS", PositionType: "Defense"},
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

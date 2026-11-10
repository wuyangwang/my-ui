package main

// aliases are the names an app reaches for that Lucide does not have.
// Each is a name of our own that points at an icon of the set:
//
//	icons.Close()   // the Lucide "x"
//	icons.Trash()   // the Lucide "trash-2"
//
// A name that the set already has is not aliased: the set's own name
// wins, and the generator leaves the alias out.
var aliases = map[string]string{
	"Add":      "plus",
	"Close":    "x",
	"Delete":   "trash-2",
	"Edit":     "pencil",
	"Error":    "circle-alert",
	"Filter":   "funnel",
	"Home":     "house",
	"Loading":  "loader-circle",
	"More":     "ellipsis",
	"Refresh":  "refresh-cw",
	"Remove":   "minus",
	"Success":  "circle-check",
	"Trash":    "trash-2",
	"Warning":  "triangle-alert",
}

package myui

import "github.com/egoist/mygo/ui"

// Variant names what a widget looks like, the shadcn set. The zero
// value, Default, is the widget's own look: the one the theme gives
// it, with nothing painted over it.
type Variant uint8

const (
	// Default is the built-in look of the widget.
	Default Variant = iota
	// Secondary is a muted fill: a face of its own, without a border.
	Secondary
	// Primary fills with the accent, as ui.PrimaryButton does.
	Primary
	// Outline is a bordered face over the background.
	Outline
	// Ghost is bare until hovered.
	Ghost
	// Destructive fills with the theme's Danger, as a Delete button.
	Destructive
	// Link is a text with nothing behind it, in the accent's color.
	Link
	// Success fills with the success color.
	Success
	// Info fills with the info color.
	Info
	// Warning fills with the warning color.
	Warning
)

// Native is the name Default had before the word table changed. It
// keeps the old code working.
const Native = Default

// Size names how roomy a widget is.
type Size uint8

const (
	// SizeMd is the built-in sizing. The zero value.
	SizeMd Size = iota
	// SizeSm is a size smaller than the built-in one.
	SizeSm
	// SizeLg is a size larger than the built-in one.
	SizeLg
	// SizeXs is the smallest size: a compact control.
	SizeXs
)

// Shape names the shape of a widget: its corners, and for a circle its
// roundness. The zero value is the theme's own corner radius.
type Shape uint8

const (
	// ShapeRect keeps the theme's corner radius.
	ShapeRect Shape = iota
	// ShapeRound rounds the corners all the way: a pill.
	ShapeRound
	// ShapeCircle rounds a square face all the way: a circle.
	ShapeCircle
)

// semantic is a role a color plays in a theme: the success of an
// operation, the warning about it, or the information in it.
type semantic uint8

const (
	semanticSuccess semantic = iota
	semanticWarning
	semanticInfo
)

// semantics is the fallback color of every semantic role, light and
// dark. MyGo's Theme carries Success and Warning, which a custom theme
// may set; the fallback is used when a theme leaves a slot empty. Info
// has no slot in MyGo's Theme (see docs/architecture.md), so it always
// comes from here.
var semantics = map[semantic][2]ui.Color{
	semanticSuccess: {ui.Hex("#16a34a"), ui.Hex("#4ade80")},
	semanticWarning: {ui.Hex("#d97706"), ui.Hex("#fbbf24")},
	semanticInfo:    {ui.Hex("#0284c7"), ui.Hex("#38bdf8")},
}

// semanticColor is the color of a semantic role on a theme. Success and
// Warning come from the theme when it sets them, so the colors follow a
// custom theme and the appearance as every other token does; Info falls
// back to the fixed table because MyGo's Theme has no slot for it.
func semanticColor(t *ui.Theme, s semantic) ui.Color {
	var fromTheme ui.Color
	switch s {
	case semanticSuccess:
		fromTheme = t.Success
	case semanticWarning:
		fromTheme = t.Warning
	}
	if fromTheme.A != 0 {
		return fromTheme
	}
	c := semantics[s]
	if t.Dark {
		return c[1]
	}
	return c[0]
}

// semanticOf is the semantic role a variant names, and whether it names
// one at all.
func semanticOf(v Variant) (semantic, bool) {
	switch v {
	case Success:
		return semanticSuccess, true
	case Warning:
		return semanticWarning, true
	case Info:
		return semanticInfo, true
	}
	return 0, false
}

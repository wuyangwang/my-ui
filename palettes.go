package myui

import (
	"github.com/egoist/mygo/ui"
)

// PaletteID names a palette of accent colors. The zero value,
// PaletteNone, keeps the accent the appearance source provides: the
// system's, or the accent the built-in theme ships with.
type PaletteID uint8

const (
	// PaletteNone keeps the accent the appearance source provides.
	PaletteNone PaletteID = iota
	// PaletteBlue is the accent of the built-in themes, given here as a
	// light and a dark triplet.
	PaletteBlue
	PaletteViolet
	PaletteRose
	PaletteOrange
	PaletteEmerald
	PaletteTeal
	PaletteSky
	PalettePink
)

// Palette is three shades of an accent in each appearance, light and
// dark: the accent, the one a control takes while hovered and the one
// while pressed. Their hues are chosen alongside, so text reads on each
// shade, and the leading shade of the dark one is lighter by as much as
// the built-in dark theme's accent is.
type Palette struct {
	Light Accent
	Dark  Accent
}

// Accent is one palette's accent, hover and pressed colors in order, as
// Theme carries them (Accent, AccentHover, AccentPressed).
type Accent [3]ui.Color

// Triplet returns the palette's shades for an appearance, with the text
// color that reads on the accent left to OnAccent.
func (p Palette) Triplet(dark bool) (hover, pressed, accent ui.Color) {
	shades := p.Light
	if dark {
		shades = p.Dark
	}
	return shades[1], shades[2], shades[0]
}

// palettes holds the palettes PaletteID names. The shades come from the
// built-in themes' accents and their neighbors on the hue, light for
// dark appearances and deep for light ones.
var palettes = map[PaletteID]Palette{
	PaletteBlue: {
		Light: Accent{ui.Hex("#2563eb"), ui.Hex("#1d4ed8"), ui.Hex("#1e40af")},
		Dark:  Accent{ui.Hex("#60a5fa"), ui.Hex("#93c5fd"), ui.Hex("#3b82f6")},
	},
	PaletteViolet: {
		Light: Accent{ui.Hex("#7c3aed"), ui.Hex("#6d28d9"), ui.Hex("#5b21b6")},
		Dark:  Accent{ui.Hex("#a78bfa"), ui.Hex("#c4b5fd"), ui.Hex("#8b5cf6")},
	},
	PaletteRose: {
		Light: Accent{ui.Hex("#e11d48"), ui.Hex("#be123c"), ui.Hex("#9f1239")},
		Dark:  Accent{ui.Hex("#fb7185"), ui.Hex("#fda4af"), ui.Hex("#f43f5e")},
	},
	PaletteOrange: {
		Light: Accent{ui.Hex("#ea580c"), ui.Hex("#c2410c"), ui.Hex("#9a3412")},
		Dark:  Accent{ui.Hex("#fb923c"), ui.Hex("#fdba74"), ui.Hex("#f97316")},
	},
	PaletteEmerald: {
		Light: Accent{ui.Hex("#059669"), ui.Hex("#047857"), ui.Hex("#065f46")},
		Dark:  Accent{ui.Hex("#34d399"), ui.Hex("#6ee7b7"), ui.Hex("#10b981")},
	},
	PaletteTeal: {
		Light: Accent{ui.Hex("#0d9488"), ui.Hex("#0f766e"), ui.Hex("#115e59")},
		Dark:  Accent{ui.Hex("#2dd4bf"), ui.Hex("#5eead4"), ui.Hex("#14b8a6")},
	},
	PaletteSky: {
		Light: Accent{ui.Hex("#0284c7"), ui.Hex("#0369a1"), ui.Hex("#075985")},
		Dark:  Accent{ui.Hex("#38bdf8"), ui.Hex("#7dd3fc"), ui.Hex("#0ea5e9")},
	},
	PalettePink: {
		Light: Accent{ui.Hex("#db2777"), ui.Hex("#be185d"), ui.Hex("#9d174d")},
		Dark:  Accent{ui.Hex("#f472b6"), ui.Hex("#f9a8d4"), ui.Hex("#ec4899")},
	},
}

// palette looks a palette id up. An unknown id reads as PaletteNone.
func palette(id PaletteID) (Palette, bool) {
	p, ok := palettes[id]
	return p, ok
}

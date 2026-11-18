package myui

import "github.com/egoist/mygo/ui"

// This file holds my-ui's default theme: shadcn/ui's "new-york" zinc
// theme, expressed in MyGo's Theme fields. It is the base Apply and
// NewTheme build on, so a palette of none means shadcn's own neutral
// scale and its near-black primary, not MyGo's blue one.
//
// The hex values are the sRGB equivalents of shadcn's oklch tokens for
// the zinc base color. MyGo has fewer slots than shadcn has CSS
// variables, so two of them share a purpose:
//
//   - Surface is shadcn's secondary/muted (zinc-100 / zinc-800), the
//     quiet fill of a secondary button, a hovered row or a sidebar.
//   - Background is shadcn's background/card (white / zinc-950); a
//     card on the page paints this and leans on its border and shadow
//     to stand out, as shadcn's bg-card does.

// ShadcnLight returns the shadcn zinc theme for a light appearance.
func ShadcnLight() *ui.Theme {
	return &ui.Theme{
		Background:      ui.Hex("#ffffff"),
		Surface:         ui.Hex("#f4f4f5"),
		SurfaceHover:    ui.Hex("#e4e4e7"),
		SurfacePressed:  ui.Hex("#d4d4d8"),
		Border:          ui.Hex("#e4e4e7"),
		Text:            ui.Hex("#09090b"),
		TextMuted:       ui.Hex("#71717a"),
		Accent:          ui.Hex("#18181b"),
		AccentHover:     ui.Hex("#27272a"),
		AccentPressed:   ui.Hex("#3f3f46"),
		AccentText:      ui.Hex("#fafafa"),
		Danger:          ui.Hex("#dc2626"),
		Warning:         ui.Hex("#d97706"),
		Success:         ui.Hex("#16a34a"),
		Selection:       ui.RGBA(9, 9, 11, 0.10),
		Focus:           ui.RGBA(161, 161, 170, 0.55),
		FocusRingWidth:  2,
		FocusRingOffset: 1,
		Scrollbar:       ui.RGBA(9, 9, 11, 0.25),
		Radius:          8,
		Spacing:         4,
		FontSize:        ui.LightTheme().FontSize,
	}
}

// ShadcnDark returns the shadcn zinc theme for a dark appearance.
func ShadcnDark() *ui.Theme {
	return &ui.Theme{
		Dark:            true,
		Background:      ui.Hex("#09090b"),
		Surface:         ui.Hex("#27272a"),
		SurfaceHover:    ui.Hex("#3f3f46"),
		SurfacePressed:  ui.Hex("#52525b"),
		Border:          ui.Hex("#3f3f46"),
		Text:            ui.Hex("#fafafa"),
		TextMuted:       ui.Hex("#a1a1aa"),
		Accent:          ui.Hex("#fafafa"),
		AccentHover:     ui.Hex("#e4e4e7"),
		AccentPressed:   ui.Hex("#d4d4d8"),
		AccentText:      ui.Hex("#18181b"),
		Danger:          ui.Hex("#ef4444"),
		Warning:         ui.Hex("#f59e0b"),
		Success:         ui.Hex("#22c55e"),
		Selection:       ui.RGBA(250, 250, 250, 0.16),
		Focus:           ui.RGBA(113, 113, 122, 0.70),
		FocusRingWidth:  2,
		FocusRingOffset: 1,
		Scrollbar:       ui.RGBA(250, 250, 250, 0.30),
		Radius:          8,
		Spacing:         4,
		FontSize:        ui.DarkTheme().FontSize,
	}
}

// shadcnBase returns the default theme for the appearance, before a
// palette, a system accent or the desktop's preferences paint over it.
func shadcnBase(dark bool) *ui.Theme {
	if dark {
		return ShadcnDark()
	}
	return ShadcnLight()
}

package controls

import "github.com/egoist/mygo/ui"

// Widgets here are re-exported unchanged only because my-ui has no
// plan to add a look to them. Controls that my-ui styles — buttons,
// toggles, sliders, rating, meter, color pickers — are implemented as
// components in this package instead.
var (
	// NewBackButton is ui.BackButton.
	NewBackButton = ui.BackButton
	// NewForwardButton is ui.ForwardButton.
	NewForwardButton = ui.ForwardButton
	// NewMenuButton is ui.MenuButton.
	NewMenuButton = ui.MenuButton
	// NewPopoverMenuButton is ui.PopoverMenuButton.
	NewPopoverMenuButton = ui.PopoverMenuButton
	// NewToolbar is ui.Toolbar.
	NewToolbar = ui.Toolbar
)

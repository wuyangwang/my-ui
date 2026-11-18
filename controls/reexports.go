package controls

import "github.com/egoist/mygo/ui"

// Only the MyGo widgets my-ui has no counterpart for are re-exported
// here. Button, Icon, Checkbox, Radio, Switch, Slider and Toggle have
// my-ui versions with the library's look, so their MyGo equivalents
// (including the Base variants) are deliberately left out.
var (
	// NewBackButton is ui.BackButton.
	NewBackButton = ui.BackButton
	// NewForwardButton is ui.ForwardButton.
	NewForwardButton = ui.ForwardButton
	// NewRangeSlider is ui.RangeSlider, a slider with two ends.
	NewRangeSlider = ui.RangeSlider
	// NewStepSlider is ui.StepSlider.
	NewStepSlider = ui.StepSlider
	// NewSpinner is ui.Spinner.
	NewSpinner = ui.Spinner
	// NewStepper is ui.Stepper.
	NewStepper = ui.Stepper
	// NewRating is ui.Rating.
	NewRating = ui.Rating
	// NewMeter is ui.Meter.
	NewMeter = ui.Meter
	// NewColorPicker is ui.ColorPicker.
	NewColorPicker = ui.ColorPicker
	// NewColorWell is ui.ColorWell.
	NewColorWell = ui.ColorWell
	// NewMenuButton is ui.MenuButton.
	NewMenuButton = ui.MenuButton
	// NewPopoverMenuButton is ui.PopoverMenuButton.
	NewPopoverMenuButton = ui.PopoverMenuButton
	// NewToolbar is ui.Toolbar.
	NewToolbar = ui.Toolbar
)

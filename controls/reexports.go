package controls

import "github.com/egoist/mygo/ui"

// Only MyGo widgets that my-ui neither provides nor plans to provide
// are re-exported here. Spinner/Spin is on the roadmap, so it is not
// exported; Button, Icon, Checkbox, Radio, Switch, Slider and Toggle
// have my-ui versions, so neither they nor their Base variants are.
var (
	// NewBackButton is ui.BackButton.
	NewBackButton = ui.BackButton
	// NewForwardButton is ui.ForwardButton.
	NewForwardButton = ui.ForwardButton
	// NewRangeSlider is ui.RangeSlider, a slider with two ends.
	NewRangeSlider = ui.RangeSlider
	// NewStepSlider is ui.StepSlider.
	NewStepSlider = ui.StepSlider
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

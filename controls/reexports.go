package controls

import "github.com/egoist/mygo/ui"

// The MyGo widgets in this file need no styling from my-ui: they are
// re-exported unchanged so an app can use one component library for
// its whole view. The Base names are MyGo's bare widgets, for an app
// that paints its own look over them.
var (
	// NewButtonBase is ui.ButtonBase, a button with no look of its own.
	NewButtonBase = ui.ButtonBase
	// NewPrimaryButton is ui.PrimaryButton.
	NewPrimaryButton = ui.PrimaryButton
	// NewBackButton is ui.BackButton.
	NewBackButton = ui.BackButton
	// NewForwardButton is ui.ForwardButton.
	NewForwardButton = ui.ForwardButton
	// NewToggleBase is ui.ToggleBase.
	NewToggleBase = ui.ToggleBase
	// NewSwitchBase is ui.SwitchBase.
	NewSwitchBase = ui.SwitchBase
	// NewSliderBase is ui.SliderBase.
	NewSliderBase = ui.SliderBase
	// NewRangeSlider is ui.RangeSlider, a slider with two ends.
	NewRangeSlider = ui.RangeSlider
	// NewStepSlider is ui.StepSlider.
	NewStepSlider = ui.StepSlider
	// NewCheckboxBase is ui.CheckboxBase.
	NewCheckboxBase = ui.CheckboxBase
	// NewCheckboxGroupBase is ui.CheckboxGroup; my-ui's own group is
	// NewCheckboxGroup.
	NewCheckboxGroupBase = ui.CheckboxGroup
	// NewCheckboxStateNative is ui.CheckboxState.
	NewCheckboxStateNative = ui.CheckboxState
	// NewCheckboxStateBase is ui.CheckboxStateBase.
	NewCheckboxStateBase = ui.CheckboxStateBase
	// NewRadioGroupBase is ui.RadioGroup; my-ui's own group is
	// NewRadioGroup.
	NewRadioGroupBase = ui.RadioGroup
	// NewIconBase is ui.Icon, with no preset of ours.
	NewIconBase = ui.Icon
	// NewSegmentedBase is ui.SegmentedBase.
	NewSegmentedBase = ui.SegmentedBase
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

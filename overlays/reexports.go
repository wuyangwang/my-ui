package overlays

import "github.com/egoist/mygo/ui"

// MyGo's overlay and menu primitives, re-exported unchanged: my-ui's
// Tooltip/Popover/Modal/Drawer build on them when the app wants the
// library's look, and these are there when it does not.
var (
	// NewTooltipBase is ui.TooltipBase.
	NewTooltipBase = ui.TooltipBase
	// NewPopoverBase is ui.PopoverBase.
	NewPopoverBase = ui.PopoverBase
	// NewDialogBase is ui.DialogBase.
	NewDialogBase = ui.DialogBase
	// NewModalBase is ui.Modal, the framework's modal dialog; my-ui's
	// themed one is NewModal.
	NewModalBase = ui.Modal
	// NewToastBase is ui.ToastBase.
	NewToastBase = ui.ToastBase
	// NewToastViewportBase is ui.ToastViewportBase.
	NewToastViewportBase = ui.ToastViewportBase
	// NewMenuBase is ui.MenuBase.
	NewMenuBase = ui.MenuBase
	// NewContextMenuBase is ui.ContextMenuBase.
	NewContextMenuBase = ui.ContextMenuBase
)

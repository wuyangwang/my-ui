package overlays

import "github.com/egoist/mygo/ui"

// Only MyGo's menu primitives, which my-ui has no counterpart for,
// are re-exported. Tooltip, Popover, Modal, Drawer and Toast have my-ui
// versions, so their MyGo Base variants are deliberately left out.
var (
	// NewMenuBase is ui.MenuBase.
	NewMenuBase = ui.MenuBase
	// NewContextMenuBase is ui.ContextMenuBase.
	NewContextMenuBase = ui.ContextMenuBase
)

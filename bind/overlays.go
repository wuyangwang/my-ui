package bind

import (
	"github.com/egoist/mygo/ui"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// TooltipContent builds a tip whose panel is made by build over anchor.
func (u UI) TooltipContent(anchor ui.Element, build func(tip ui.Element)) *overlaypkg.Tooltip {
	return overlaypkg.NewTooltipContent(u.c, anchor, build)
}

// ToastViewport builds the window's toast area; call it once per view.
func (u UI) ToastViewport(opts ...overlaypkg.ViewportOption) *overlaypkg.ToastViewport {
	return overlaypkg.NewToastViewport(u.c, opts...)
}

// Toast builds a toast with title, ready for Description and Show.
func (u UI) Toast(title string) *overlaypkg.Toast {
	return overlaypkg.NewToast(u.c, title)
}

// SplitButton builds a button plus a menu of related actions.
func (u UI) SplitButton(label string, onClick func(), items []overlaypkg.MenuItem) *overlaypkg.SplitButton {
	return overlaypkg.NewSplitButton(u.c, label, onClick, items)
}

// PrimarySplitButton builds the accent-colored split button.
func (u UI) PrimarySplitButton(label string, onClick func(), items []overlaypkg.MenuItem) *overlaypkg.SplitButton {
	return overlaypkg.NewPrimarySplitButton(u.c, label, onClick, items)
}

// Menubar builds a row of pull-down menus.
func (u UI) Menubar(menus []overlaypkg.MenubarMenu) *overlaypkg.Menubar {
	return overlaypkg.NewMenubar(u.c, menus)
}

// HoverCard builds a panel shown over anchor while *open is true.
func (u UI) HoverCard(open *bool, anchor func() ui.Element, build func(panel ui.Element)) *overlaypkg.HoverCard {
	return overlaypkg.NewHoverCard(u.c, open, anchor, build)
}

// DropdownMenu builds a menu shown from trigger while *open is true.
func (u UI) DropdownMenu(open *bool, trigger func() ui.Element, items []overlaypkg.MenuItem) *overlaypkg.DropdownMenu {
	return overlaypkg.NewDropdownMenu(u.c, open, trigger, items)
}

// CommandPalette builds a centered dialog of commands over the window.
func (u UI) CommandPalette(open *bool, query *string, commands []overlaypkg.Command) ui.Element {
	return overlaypkg.NewCommandPalette(u.c, open, query, commands)
}

// ContextMenu builds a menu shown over target while *open is true.
func (u UI) ContextMenu(open *bool, target ui.Element, items []overlaypkg.MenuItem) *overlaypkg.ContextMenu {
	return overlaypkg.NewContextMenu(u.c, open, target, items)
}

// Drawer builds a panel sliding in from side over the window.
func (u UI) Drawer(open *bool, side overlaypkg.DrawerSide, trigger func() ui.Element, build func(panel ui.Element)) *overlaypkg.Drawer {
	return overlaypkg.NewDrawer(u.c, open, side, trigger, build)
}

package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// Modal shows a dialog over a dimmed backdrop, while the app's flag is
// true.
//
//	myui.NewModal(c, &app.open, func() ui.Element {
//		return myui.NewButton(c, "删除").Destructive().Element
//	}, func(panel ui.Element) {
//		ui.Text(c, "确定删除这个项目吗？")
//	})
//
// The trigger is built by the app, and a click on it opens the dialog; a
// click on the backdrop closes the dialog, and Escape closes it too.
//
// The backdrop and the panel are styled by the library, as shadcn's: the
// scrim dims the window, the panel is a card in the middle of it, with a
// close button in its corner.
type Modal struct {
	ui.Element
	open *bool
}

// NewModal shows a dialog over the window while open is true. The flag
// is the app's: it opens and closes the dialog.
//
// The trigger is built where the modal is, so it sits beside the dialog
// it opens; a click on it opens the dialog in the same frame.
//
// The build callback fills the panel: a title, a text and the buttons.
// The panel is the dialog as assistive technology sees it: name it with
// panel.Label("删除项目"), so a screen reader says what the dialog is
// for.
func NewModal(c *ui.Context, open *bool, trigger func() ui.Element, build func(panel ui.Element)) *Modal {
	m := &Modal{open: open}
	if trigger != nil {
		t := trigger()
		// A click on the trigger opens the dialog.
		if t.Clicked() && m.open != nil {
			*m.open = true
		}
	}
	modal := ui.DialogBase(c, open, func(backdrop ui.Element, panel ui.Element) {
		th := c.Theme()
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5))
		panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(6)).Shadow(0, 8, 24, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(128))
		panel.Children(func() {
			// The close button, in the corner of the panel, as shadcn's.
			close := NewButton(c, "").Ghost().Xs()
			close.Icon(icons.X()).Label("关闭")
			close.Element = close.Element.Absolute().Top(th.Space(4)).Right(th.Space(4))
			if close.Clicked() && m.open != nil {
				*m.open = false
			}
			if build != nil {
				build(panel)
			}
		})
	})
	m.Element = modal
	return m
}

// Open reports whether the dialog shows.
func (m *Modal) Open() bool { return m.open != nil && *m.open }

// SetOpen opens or closes the dialog and returns it, for chaining.
func (m *Modal) SetOpen(open bool) *Modal {
	if m.open != nil {
		*m.open = open
	}
	return m
}

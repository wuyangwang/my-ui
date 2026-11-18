package overlays

import (
	"strings"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// Command is one row of a command palette. OnSelect runs when the row
// is chosen; the palette then closes and clears the query.
type Command struct {
	Label    string
	Icon     *ui.SVG
	Shortcut string
	OnSelect func()
}

// NewCommandPalette builds a centered command dialog over the window.
// open and query belong to the app, like DialogBase's open flag.
func NewCommandPalette(c *ui.Context, open *bool, query *string, commands []Command) ui.Element {
	if query == nil {
		query = new(string)
	}
	t := c.Theme()
	return ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		backdrop.Background(t.Text.Alpha(0.45))
		panel.Width(t.Space(120)).Radius(Radius(t, RadiusXl)).
			Background(t.Background).Border(1, t.Border).
			Shadow(0, 16, 48, -12, t.Text.Alpha(0.32))
		panel.Children(func() {
			input := ui.TextInput(c, query).Placeholder(Message(c, "searchCommands")).
				FontSize(TextSize(t, TextMd)).Padding(t.Space(3), t.Space(4))
			input.Background(ui.Transparent).Border(0, ui.Transparent)
			ui.Box(c).FillWidth().Height(1).Background(t.Border)

			ui.Column(c).Gap(t.Space(1)).Padding(t.Space(2)).Children(func() {
				found := 0
				q := strings.ToLower(*query)
				for _, cmd := range commands {
					if q != "" && !strings.Contains(strings.ToLower(cmd.Label), q) {
						continue
					}
					found++
					row := ctrlpkg.NewButton(c, cmd.Label).Ghost().Block().Sm().
						Justify(ui.Start).Padding(t.Space(2), t.Space(3))
					row.Children(func() {
						if cmd.Icon != nil {
							ui.Icon(c, cmd.Icon).Size(IconSizeSm, IconSizeSm).
								TextColor(t.TextMuted)
						}
					})
					if cmd.Shortcut != "" {
						ui.Text(c, cmd.Shortcut).FontSize(TextSize(t, TextXs)).
							TextColor(t.TextMuted)
					}
					if row.Clicked() {
						if cmd.OnSelect != nil {
							cmd.OnSelect()
						}
						*open = false
						*query = ""
					}
				}
				if found == 0 {
					ui.Text(c, Message(c, "noCommands")).
						FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted).
						Padding(t.Space(4), t.Space(3))
				}
			})
		})
	})
}

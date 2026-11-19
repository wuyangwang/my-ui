package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// UploadOptions configures NewUpload.
type UploadOptions struct {
	// Accept filters dropped and picked paths; nil accepts all.
	Accept func(path string) bool
	// OnFiles receives the accepted paths, from either path.
	OnFiles func(paths []string)
	// OnPick, when set, adds a button that opens the system file picker.
	// The app supplies it: mygo's OpenFile lives in mygo's root package,
	// which a library must not import. It returns the chosen path and
	// whether a file was chosen; on true the path goes through Accept and
	// OnFiles just like a dropped one.
	OnPick func() (path string, ok bool)
}

// NewUpload builds a drop zone with an optional "open" button. MyGo hands
// the dropped file paths to the element; the library never touches the
// filesystem. Clicking the button calls OnPick, and the chosen path, when
// there is one, takes the same route as a drop.
func NewUpload(c *ui.Context, o UploadOptions) ui.Element {
	t := c.Theme()
	face := ui.Row(c).FillWidth().Height(t.Space(30)).
		AlignItems(ui.Center).Justify(ui.Center).Gap(t.Space(2)).
		Radius(Radius(t, RadiusMd)).Border(1, t.Border)
	if face.FileDragOver() {
		face = face.Background(t.Accent.Alpha(0.08)).Border(1, t.Accent)
	} else {
		face = face.Background(t.Surface)
	}
	if paths := face.DroppedFiles(); len(paths) > 0 {
		reportFiles(o, paths)
	}
	face.Children(func() {
		ui.Icon(c, icons.Upload()).Size(IconSizeLg, IconSizeLg).TextColor(t.TextMuted)
		ui.Column(c).Gap(t.Space(1)).Children(func() {
			ui.Text(c, Message(c, "dropFiles")).FontSize(TextSize(t, TextSm)).Bold()
			ui.Text(c, Message(c, "dropHint")).
				FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
		})
		if o.OnPick != nil {
			open := controls.NewButton(c, Message(c, "open")).Outline().Sm()
			if open.Clicked() {
				if path, ok := o.OnPick(); ok {
					reportFiles(o, []string{path})
				}
			}
		}
	})
	return face
}

// reportFiles filters paths through Accept and hands the kept ones to
// OnFiles. Both the drop and the pick path use it.
func reportFiles(o UploadOptions, paths []string) {
	kept := []string{}
	for _, path := range paths {
		if o.Accept == nil || o.Accept(path) {
			kept = append(kept, path)
		}
	}
	if len(kept) > 0 && o.OnFiles != nil {
		o.OnFiles(kept)
	}
}

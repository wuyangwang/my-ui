package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// UploadOptions configures NewUpload.
type UploadOptions struct {
	// Accept filters dropped paths; nil accepts all.
	Accept func(path string) bool
	// OnFiles receives the accepted paths.
	OnFiles func(paths []string)
}

// NewUpload builds a drop zone. MyGo hands the dropped file paths to
// the element; the library never touches the filesystem. Opening the
// system file picker needs mygo's root package, which this module does
// not import — so the zone is drag-and-drop only, and says so.
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
	face.Children(func() {
		ui.Icon(c, icons.Upload()).Size(IconSizeLg, IconSizeLg).TextColor(t.TextMuted)
		ui.Column(c).Gap(t.Space(1)).Children(func() {
			ui.Text(c, Message(c, "dropFiles")).FontSize(TextSize(t, TextSm)).Bold()
			ui.Text(c, Message(c, "dropHint")).
				FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
		})
	})
	return face
}

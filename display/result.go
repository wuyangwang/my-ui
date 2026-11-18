package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// ResultStatus is the outcome a Result page reports.
type ResultStatus uint8

const (
	// ResultSuccess is a completed operation.
	ResultSuccess ResultStatus = iota
	// ResultWarning is an operation that needs the user's attention.
	ResultWarning
	// ResultInfo is a neutral note.
	ResultInfo
	// ResultError is a failed operation.
	ResultError
)

// ResultOptions configures NewResult.
type ResultOptions struct {
	// Status picks the icon and its semantic color.
	Status ResultStatus
	// Title is the short outcome line.
	Title string
	// Description holds the details.
	Description string
	// Extra builds the controls under the description.
	Extra func()
}

// NewResult builds a centered outcome page: a status icon in a tinted
// disc, a title, a description and the app's own next-step controls.
func NewResult(c *ui.Context, o ResultOptions) ui.Element {
	t := c.Theme()
	icon, ink := resultLook(c, o.Status)
	return ui.Column(c).FillWidth().Padding(t.Space(10), t.Space(6)).
		AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
		disc := ui.Row(c).Size(t.Space(14), t.Space(14)).
			AlignItems(ui.Center).Justify(ui.Center).
			Radius(t.Space(7)).Background(ink.Alpha(0.12))
		disc.Children(func() {
			ui.Icon(c, icon).Size(IconSizeLg, IconSizeLg).TextColor(ink)
		})
		if o.Title != "" {
			ui.Text(c, o.Title).FontSize(TextSize(t, Text2Xl)).Bold()
		}
		if o.Description != "" {
			ui.Text(c, o.Description).FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
		}
		if o.Extra != nil {
			ui.Row(c).Gap(t.Space(2)).Children(o.Extra)
		}
	})
}

func resultLook(c *ui.Context, status ResultStatus) (*ui.SVG, ui.Color) {
	t := c.Theme()
	switch status {
	case ResultWarning:
		return icons.TriangleAlert(), SemanticColor(t, SemanticWarning)
	case ResultInfo:
		return icons.Info(), SemanticColor(t, SemanticInfo)
	case ResultError:
		return icons.CircleX(), t.Danger
	default:
		return icons.CircleCheck(), SemanticColor(t, SemanticSuccess)
	}
}

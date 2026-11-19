package layout

import (
	"strconv"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// StepStatus is one step's state; StepAuto derives it from Current.
type StepStatus int

const (
	// StepAuto derives the status from the current index.
	StepAuto StepStatus = iota
	// StepWait is a step not reached yet.
	StepWait
	// StepProcess is the current step.
	StepProcess
	// StepFinish is a completed step.
	StepFinish
	// StepError marks a failed step.
	StepError
)

// StepItem is one step's title, description and optional glyph.
type StepItem struct {
	Title, Description string
	Icon               *ui.SVG
	Status             StepStatus
}

// StepsProps configures NewSteps.
type StepsProps struct {
	Items    []StepItem
	Current  int // index in process; negative leaves all neutral
	Vertical bool
	Dot      bool
}

// Steps shows a process as numbered circles joined by connectors, or
// as dots. Current belongs to the app and writes the finished/active
// split; status overrides are per item.
type Steps struct {
	ui.Element
	c *ui.Context
	p StepsProps
}

// NewSteps builds the step row or column.
func NewSteps(c *ui.Context, p StepsProps) *Steps {
	s := &Steps{c: c, p: p}
	s.build()
	return s
}

// Vertical stacks the steps with the connector running down beside them.
func (s *Steps) Vertical(on bool) *Steps {
	s.p.Vertical = on
	s.build()
	return s
}

// Dot uses dots instead of numbered circles.
func (s *Steps) Dot(on bool) *Steps {
	s.p.Dot = on
	s.build()
	return s
}

func (s *Steps) build() {
	t := s.c.Theme()
	p := s.p
	var root ui.Element
	if p.Vertical {
		root = ui.Column(s.c).AlignSelf(ui.Stretch).Gap(t.Space(2))
	} else {
		root = ui.Row(s.c).AlignItems(ui.Center).Gap(t.Space(2)).AlignSelf(ui.Stretch)
	}
	root.Children(func() {
		for i, item := range p.Items {
			if i > 0 {
				done := p.Current >= i
				var line ui.Element
				if p.Vertical {
					line = ui.Box(s.c).Width(2).Height(t.Space(6))
				} else {
					line = ui.Box(s.c).Height(2).Grow(1)
				}
				if done {
					line = line.Background(t.Accent)
				} else {
					line = line.Background(t.Border)
				}
			}
			status := item.Status
			if status == StepAuto {
				switch {
				case p.Current < 0:
					status = StepWait
				case i < p.Current:
					status = StepFinish
				case i == p.Current:
					status = StepProcess
				default:
					status = StepWait
				}
			}
			ink, face := t.TextMuted, t.Surface
			switch status {
			case StepProcess:
				ink, face = t.AccentText, t.Accent
			case StepFinish:
				ink, face = t.AccentText, t.Accent.Alpha(0.16)
			case StepError:
				ink, face = t.AccentText, t.Danger
			}
			var marker ui.Element
			if p.Dot {
				marker = ui.Box(s.c).Size(t.Space(2), t.Space(2)).Radius(t.Space(1))
			} else {
				marker = ui.Row(s.c).Size(t.Space(7), t.Space(7)).
					Justify(ui.Center).AlignItems(ui.Center).Radius(Radius(t, RadiusMd))
			}
			marker = marker.Background(face).TextColor(ink)
			marker.Children(func() {
				if item.Icon != nil {
					ui.Icon(s.c, item.Icon).Size(12, 12).TextColor(ink)
					return
				}
				if !p.Dot {
					label := strconv.Itoa(i + 1)
					if status == StepFinish {
						ui.Icon(s.c, icons.Check()).Size(12, 12).TextColor(ink)
					} else {
						ui.Text(s.c, label).FontSize(12).TextColor(ink)
					}
				}
			})
			ui.Column(s.c).Gap(t.Space(0.5)).Children(func() {
				ui.Text(s.c, item.Title).FontSize(TextSize(t, TextSm)).Bold().TextColor(t.Text)
				if item.Description != "" {
					ui.Text(s.c, item.Description).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
				}
			})
			_ = marker
		}
	})
	s.Element = root
}

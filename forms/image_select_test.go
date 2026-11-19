package forms

import (
	"image"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func swatch() ui.ImageSource {
	return ui.NewBitmap(image.NewRGBA(image.Rect(0, 0, 40, 24)))
}

func TestImageSelectWritesTheChoice(t *testing.T) {
	selected := "auto"
	view := func(c *ui.Context) {
		NewImageSelect(c, &selected, ImageSelectProps[string]{
			Options: []ImageSelectItem[string]{
				{Value: "light", Label: "浅色", Image: swatch()},
				{Value: "dark", Label: "深色", Image: swatch()},
				{Value: "auto", Label: "跟随系统", Image: swatch()},
			},
		})
	}
	tt := ui.NewTester(view, 420, 240)
	tt.Advance(time.Second)
	if !tt.HasText("浅色") || !tt.HasText("深色") {
		t.Fatal("the image select tiles are not on screen")
	}
	if err := tt.Click("深色"); err != nil {
		t.Fatalf("the tile did not choose: %v", err)
	}
	tt.Advance(time.Second)
	if selected != "dark" {
		t.Errorf("the choice was %q, want dark", selected)
	}
}

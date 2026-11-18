package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestCheckboxLgGlyphIsCentered guards the large check box: its tick
// must be drawn once, in the middle of the square. A nested Row used
// to duplicate the glyph at the box's edge at this size.
func TestCheckboxLgGlyphIsCentered(t *testing.T) {
	checked := true
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		NewCheckbox(c, &checked, "").Lg()
	}
	tt := ui.NewTester(view, 80, 60)
	tt.Advance(time.Second)

	minX, minY, maxX, maxY := 999, 999, -1, -1
	var sumX, sumY float64
	ink := 0
	for y := 0; y < 60; y++ {
		for x := 0; x < 80; x++ {
			r, g, b := pixelAt(tt, x, y)
			if r < 100 && g < 150 && b > 150 { // the blue face
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			} else if r > 240 && g > 240 && b > 240 && maxX >= 0 &&
				x >= minX && x <= maxX && y >= minY && y <= maxY {
				sumX += float64(x)
				sumY += float64(y)
				ink++
			}
		}
	}
	if maxX < 0 {
		t.Fatal("the checked box was not painted")
	}
	if ink == 0 {
		t.Fatal("the tick was not painted")
	}
	cx, cy := float64(minX+maxX)/2, float64(minY+maxY)/2
	ix, iy := sumX/float64(ink), sumY/float64(ink)
	if diff := ix - cx; diff < -2 || diff > 2 {
		t.Errorf("the tick's x centroid is %.1f, box center %.1f", ix, cx)
	}
	if diff := iy - cy; diff < -2 || diff > 2 {
		t.Errorf("the tick's y centroid is %.1f, box center %.1f", iy, cy)
	}
}

func pixelAt(tt *ui.Tester, x, y int) (uint8, uint8, uint8) {
	p := tt.Image().At(x, y)
	r, g, b, _ := p.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

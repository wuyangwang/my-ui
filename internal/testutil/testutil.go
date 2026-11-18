// Package testutil holds the pixel helpers the component tests share.
package testutil

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
)

// Pixel reads what the tester drew at a DIP of the window.
func Pixel(tt *ui.Tester, x, y float32) (uint8, uint8, uint8) {
	p := color.NRGBAModel.Convert(tt.Image().At(int(x), int(y))).(color.NRGBA)
	return p.R, p.G, p.B
}

// SamePixel checks a pixel's color.
func SamePixel(t *testing.T, r, g, b uint8, want ui.Color, what string) {
	t.Helper()
	if r != want.R || g != want.G || b != want.B {
		t.Errorf("%s drew (%d, %d, %d), want #%02x%02x%02x", what, r, g, b, want.R, want.G, want.B)
	}
}

// HexRGB decodes a #rrggbb string.
func HexRGB(t *testing.T, s string) (uint8, uint8, uint8) {
	t.Helper()
	var r, g, b uint8
	if _, err := fmt.Sscanf(s, "#%02x%02x%02x", &r, &g, &b); err != nil {
		t.Fatal(err)
	}
	return r, g, b
}

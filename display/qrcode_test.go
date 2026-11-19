package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestQREncoder(t *testing.T) {
	grid := qrEncode("https://lobehub.com")
	if len(grid) == 0 {
		t.Fatal("encoder returned nothing")
	}
	if len(grid) != len(grid[0]) {
		t.Fatalf("not square: %dx%d", len(grid), len(grid[0]))
	}
	if len(grid) != 21 && len(grid) != 25 && len(grid) != 29 && len(grid) != 33 && len(grid) != 37 {
		t.Fatalf("unexpected version size %d", len(grid))
	}
	again := qrEncode("https://lobehub.com")
	if len(grid) != len(again) {
		t.Fatal("encoding is not deterministic")
	}
	for y := range grid {
		for x := range grid {
			if grid[y][x] != again[y][x] {
				t.Fatalf("encoding differs at %d,%d", x, y)
			}
		}
	}
	if qrEncode("") == nil {
		t.Fatal("empty value should still encode")
	}
}

func TestQRCodeComponent(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			NewQRCode(c, "https://lobehub.com", QRCodeOptions{Size: 120, Label: "site"})
			NewQRCode(c, "https://mygo.lobehub.com", QRCodeOptions{
				Size: 96, Status: QRExpired, Label: "expired",
				OnRefresh: func() {},
			})
		})
	}
	tt := ui.NewTester(view, 420, 240)
	tt.Advance(time.Second)
	for _, label := range []string{"site", "expired"} {
		if _, ok := tt.Access(label); !ok {
			t.Fatalf("qr %s missing", label)
		}
	}
	if !tt.HasText("Expired") {
		t.Fatal("expired overlay missing")
	}
}

package display

import (
	"strings"
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

// TestQRCodeShowsErrorForOverlongValue checks both outcomes: a short value
// builds a real code, while a value beyond the encoder's limit reports
// Valid() == false and renders a visible error instead of a blank card.
func TestQRCodeShowsErrorForOverlongValue(t *testing.T) {
	overlong := strings.Repeat("x", 200)
	shortValid, longValid := false, true
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			shortValid = NewQRCodeCard(c, "https://lobehub.com",
				QRCodeOptions{Size: 80, Label: "short"}).Valid()
			longValid = NewQRCodeCard(c, overlong,
				QRCodeOptions{Size: 80, Label: "long"}).Valid()
		})
	}
	tt := ui.NewTester(view, 360, 200)
	tt.SetLocale("zh-Hans")
	tt.Advance(time.Second)

	if !shortValid {
		t.Error("a short value should report Valid() == true")
	}
	if longValid {
		t.Error("an overlong value should report Valid() == false")
	}
	if _, ok := tt.Access("short"); !ok {
		t.Error("the short code did not render")
	}
	if _, ok := tt.Access("long"); !ok {
		t.Error("the overlong card did not render")
	}
	if !tt.HasText("内容过长，无法生成二维码") {
		t.Error("the overlong card has no visible error message")
	}
}

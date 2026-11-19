package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestAvatarFallbacks(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			NewAvatar(c, "Ada Lovelace", AvatarOptions{Fallback: AvatarText})
			NewAvatar(c, "🤯", AvatarOptions{})
			NewAvatar(c, "Lobe", AvatarOptions{}).Circle()
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if !tt.HasText("AD") {
		t.Fatal("text fallback missing")
	}
	if !tt.HasText("🤯") {
		t.Fatal("emoji avatar missing")
	}
}

func TestAvatarSizes(t *testing.T) {
	width := func(tt *ui.Tester, label string) float32 {
		info, ok := tt.Access(label)
		if !ok {
			t.Fatalf("%s missing", label)
		}
		return info.Bounds.W
	}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewAvatar(c, "A", AvatarOptions{Fallback: AvatarText}).Size(SizeSm)
			NewAvatar(c, "B", AvatarOptions{Fallback: AvatarText}).Size(SizeLg)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if !(width(tt, "A") < width(tt, "B")) {
		t.Fatalf("avatar sizes not increasing: %v %v", width(tt, "A"), width(tt, "B"))
	}
}

func TestAvatarDarkTheme(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewAvatar(c, "Grace Hopper", AvatarOptions{}).Pixels(40).Circle()
	}
	tt := ui.NewTester(view, 160, 120)
	tt.Advance(time.Second)
	if tt.Image() == nil {
		t.Fatal("no frame drawn")
	}
}

package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

func TestAlertShowsItsTitleAndDescription(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewAlert(c, "注意").Description("这条信息需要留意")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("注意"); !ok {
		t.Fatal("the title of the alert is not on screen")
	}
	if _, ok := tt.Find("这条信息需要留意"); !ok {
		t.Fatal("the description of the alert is not on screen")
	}
}

func TestAlertIsAnAlertToAssistiveTechnology(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewAlert(c, "出错了").Error()
		NewAlert(c, "好了").Success().Role(ui.RoleStatus)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.AccessByRole(ui.RoleAlert, ""); !ok {
		t.Fatal("no alert is in the tree of the assistive technology")
	}
	for _, name := range []string{"出错了", "好了"} {
		if _, ok := tt.Access(name); !ok {
			t.Errorf("the title %q of the alert is not on screen", name)
		}
	}
}

func TestAlertTypeAndVariantAreOrthogonal(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewAlert(c, "soft error").Error().Variant(AlertSoft)
			NewAlert(c, "outlined success").Success().Variant(AlertOutlined)
			NewAlert(c, "plain info").Info().Variant(AlertPlain)
			NewAlert(c, "secondary").Secondary().NoIcon(true)
		})
	}
	tt := ui.NewTester(view, 360, 320)
	tt.Advance(time.Second)
	for _, name := range []string{"soft error", "outlined success", "plain info", "secondary"} {
		if _, ok := tt.Find(name); !ok {
			t.Errorf("alert %q missing", name)
		}
	}
}

func TestAlertIconNoIconSelectableAndClosed(t *testing.T) {
	closed := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewAlert(c, "with icon").Warning().Icon(icons.TriangleAlert())
			NewAlert(c, "no icon").Warning().NoIcon(true)
			NewAlert(c, "selectable").Info().Selectable(true)
			NewAlert(c, "closable").Info().Closed(&closed)
		})
	}
	tt := ui.NewTester(view, 360, 360)
	tt.Advance(time.Second)
	if _, ok := tt.Find("with icon"); !ok {
		t.Fatal("the icon alert is missing")
	}
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close control of the alert is missing")
	}
	x, y := close.Bounds.X+close.Bounds.W/2, close.Bounds.Y+close.Bounds.H/2
	tt.Press(x, y)
	tt.Advance(time.Second)
	tt.Release(x, y)
	tt.Advance(time.Second)
	if !closed {
		t.Error("the close control did not write the app's flag")
	}
	if _, ok := tt.Find("closable"); ok {
		t.Error("a closed alert still shows")
	}
}

package window

import (
	"image/color"
	"testing"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	controls "github.com/wuyangwang/my-ui/controls"
)

// pixel reads what the tester drew at a DIP of the window.
func pixel(tt *ui.Tester, x, y int) (uint8, uint8, uint8) {
	c := color.NRGBAModel.Convert(tt.Image().At(x, y)).(color.NRGBA)
	return c.R, c.G, c.B
}

// shellTester runs one frame of the shell without a window: the
// controls' actions record their calls in the returned map, and the
// app's view marks that it ran.
func shellTester(t *testing.T, o Options) (*ui.Tester, map[string]bool, *bool) {
	t.Helper()
	fired := map[string]bool{}
	act := &actions{
		minimize:  func() { fired["minimize"] = true },
		maximize:  func() { fired["maximize"] = true },
		close:     func() { fired["close"] = true },
		setBg:     func(s string) error { fired["bg:"+s] = true; return nil },
		setSource: func(s mygo.ThemeSource) { fired["source:"+string(s)] = true },
	}
	ran := false
	if o.View == nil {
		o.View = func(c *ui.Context) { ran = true }
	} else {
		view := o.View
		o.View = func(c *ui.Context) { view(c); ran = true }
	}
	sh := &shell{o: o, act: act}
	tt := ui.NewTester(sh.render, 320, 240)
	tt.Advance(time.Second)
	return tt, fired, &ran
}

func TestShellLaysTheBarOverTheView(t *testing.T) {
	o := Options{Name: "MyGo App", Look: controls.TitleBarWindows}
	o.View = func(c *ui.Context) {
		ui.Text(c, "Body").Bold()
	}
	tt, _, ran := shellTester(t, o)

	if _, ok := tt.Find("MyGo App"); !ok {
		t.Fatal("the bar's name is missing")
	}
	if _, ok := tt.Find("Body"); !ok {
		t.Fatal("the app's view is missing")
	}
	for _, name := range []string{"Minimize", "Maximize", "Close"} {
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("the %s control is missing", name)
		}
	}
	// The name sits above the body: the bar owns the top strip.
	name, _ := tt.Find("MyGo App")
	body, _ := tt.Find("Body")
	if name.Y >= body.Y {
		t.Errorf("the name at y=%v does not sit over the body at y=%v", name.Y, body.Y)
	}
	if !*ran {
		t.Error("the app's view did not run")
	}
}

func TestShellWithoutAViewStillShowsTheBar(t *testing.T) {
	tt, _, _ := shellTester(t, Options{Name: "MyGo App", Look: controls.TitleBarMac})
	if _, ok := tt.Find("MyGo App"); !ok {
		t.Fatal("the bar's name is missing")
	}
	if _, ok := tt.Access("Close"); !ok {
		t.Fatal("the close light is missing")
	}
}

func TestShellThemeFeedsTheBar(t *testing.T) {
	o := Options{Name: "MyGo App", Look: controls.TitleBarWindows}
	o.Theme = func() myui.Options { return myui.Options{Appearance: myui.AppearanceDark} }
	o.View = func(c *ui.Context) {
		if !c.Theme().Dark {
			t.Error("the app's view did not get the dark theme")
		}
	}
	shellTester(t, o)
}

func TestShellWindowBackgroundFollowsTheAppearance(t *testing.T) {
	appearance := myui.AppearanceLight
	o := Options{Name: "MyGo App", Look: controls.TitleBarWindows}
	o.Theme = func() myui.Options { return myui.Options{Appearance: appearance} }
	tt, fired, _ := shellTester(t, o)
	if !fired["bg:#ffffff"] {
		t.Errorf("the light frame did not set the window background: %v", fired)
	}
	appearance = myui.AppearanceDark
	tt.Frame()
	if !fired["bg:#09090b"] {
		t.Errorf("the dark frame left a light window background: %v", fired)
	}
}

func TestDefaultsGiveTheWindowItsName(t *testing.T) {
	if got := withDefaults(Options{Name: "My App"}).Title; got != "My App" {
		t.Errorf("the default title is %q, want the app name", got)
	}
	if got := withDefaults(Options{Title: "Custom", Name: "My App"}).Title; got != "Custom" {
		t.Errorf("an explicit title %q was overwritten", got)
	}
}

func TestShellFrameAppearanceFollowsTheTheme(t *testing.T) {
	appearance := myui.AppearanceLight
	o := Options{Name: "MyGo App", Theme: func() myui.Options {
		return myui.Options{Appearance: appearance}
	}}
	tt, fired, _ := shellTester(t, o)
	if !fired["source:light"] {
		t.Errorf("the light frame did not set the native window's appearance: %v", fired)
	}
	appearance = myui.AppearanceDark
	tt.Frame()
	if !fired["source:dark"] {
		t.Errorf("the dark frame left a light native window border: %v", fired)
	}
}

func TestShellPaintsDarkFrameEdges(t *testing.T) {
	o := Options{
		Name:  "MyGo App",
		Look:  controls.TitleBarWindows,
		Theme: func() myui.Options { return myui.Options{Appearance: myui.AppearanceDark} },
	}
	tt, _, _ := shellTester(t, o)
	// The root would draw the system's light background over the
	// transparent bar; the shell has to paint its own on every edge.
	for _, at := range []struct{ x, y int }{{1, 1}, {160, 10}, {318, 1}, {1, 238}} {
		r, g, b := pixel(tt, at.x, at.y)
		if r > 60 || g > 60 || b > 60 {
			t.Errorf("the dark frame drew (%d, %d, %d) at (%d, %d), want the theme's dark background",
				r, g, b, at.x, at.y)
		}
	}
}

func TestShellControlsFireThroughTheActions(t *testing.T) {
	tt, fired, _ := shellTester(t, Options{Name: "MyGo App", Look: controls.TitleBarWindows})
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	for _, click := range []string{"Close", "Minimize", "Maximize"} {
		if err := tt.Click(click); err != nil {
			t.Fatal(err)
		}
		adv()
	}
	for _, name := range []string{"close", "minimize", "maximize"} {
		if !fired[name] {
			t.Errorf("the %s click did not reach the window action", name)
		}
	}
}

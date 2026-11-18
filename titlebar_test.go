package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// flags collects what the title bar's callbacks raised.
type flags struct {
	minimized, toggled, closed bool
}

// titleBarTester runs one frame with a title bar whose callbacks raise
// flags. The look is fixed, so the tests read the same on every
// platform the library builds on.
func titleBarTester(t *testing.T, look TitleBarLook, o TitleBarOptions) (*ui.Tester, *flags) {
	t.Helper()
	f := &flags{}
	if o.OnMinimize == nil {
		o.OnMinimize = func() { f.minimized = true }
	}
	if o.OnMaximize == nil {
		o.OnMaximize = func() { f.toggled = true }
	}
	if o.OnClose == nil {
		o.OnClose = func() { f.closed = true }
	}
	o.Look = look
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTitleBar(c, o)
	}
	tt := ui.NewTester(view, 320, 240)
	return tt, f
}

func TestTitleBarRendersPartsAndClicks(t *testing.T) {
	tt, f := titleBarTester(t, TitleBarWindows, TitleBarOptions{Name: "MyGo App"})

	if _, ok := tt.Find("MyGo App"); !ok {
		t.Fatal("the app name is missing")
	}
	for _, name := range []string{"Minimize", "Maximize", "Close"} {
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("the %s control is missing", name)
		}
	}
	// A control acts like a button to assistive technology.
	if info, ok := tt.Access("Close"); !ok || info.Role != ui.RoleButton {
		t.Errorf("the close control is not a button (found %v)", info.Role)
	}
	// The name sits left of the controls.
	name, _ := tt.Find("MyGo App")
	if close, ok := tt.Access("Close"); ok && name.X+name.W > close.Bounds.X {
		t.Errorf("the name at x=%v runs into the controls at x=%v",
			name.X+name.W, close.Bounds.X)
	}

	// The clicks reach the app's callbacks.
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	if err := tt.Click("Close"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.closed {
		t.Error("the close click did not reach the app")
	}
	if err := tt.Click("Minimize"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.minimized {
		t.Error("the minimize click did not reach the app")
	}
	if err := tt.Click("Maximize"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.toggled {
		t.Error("the maximize click did not reach the app")
	}
}

func TestTitleBarLeavesControlsOutWhenNil(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTitleBar(c, TitleBarOptions{Name: "MyGo App", Look: TitleBarWindows})
	}
	tt := ui.NewTester(view, 320, 240)
	for _, name := range []string{"Close", "Minimize", "Maximize"} {
		if _, ok := tt.Access(name); ok {
			t.Errorf("a title bar without a callback drew a %s control", name)
		}
	}
	if _, ok := tt.Find("MyGo App"); !ok {
		t.Error("the app name is missing")
	}
}

func TestTitleBarFlipsToRestoreWhenMaximized(t *testing.T) {
	var maximized bool
	tt, _ := titleBarTester(t, TitleBarWindows,
		TitleBarOptions{Name: "MyGo App", Maximized: &maximized})

	if _, ok := tt.Access("Restore"); ok {
		t.Fatal("a window at rest shows the restore glyph")
	}
	maximized = true
	tt.Frame()
	tt.Advance(time.Second)
	if _, ok := tt.Access("Restore"); !ok {
		t.Error("the maximized window does not show the restore glyph")
	}
	if _, ok := tt.Access("Maximize"); ok {
		t.Error("the maximized window still shows the maximize glyph")
	}
}

func TestTitleBarKeepsClearOfSystemControls(t *testing.T) {
	tt, _ := titleBarTester(t, TitleBarWindows, TitleBarOptions{Name: "MyGo App"})
	tt.SetTitleBar(ui.TitleBar{Height: 40, Left: 70, Right: 0})
	tt.Advance(time.Second)

	// The system's controls take the left strip, so ours stay out.
	for _, name := range []string{"Minimize", "Maximize", "Close"} {
		if _, ok := tt.Access(name); ok {
			t.Errorf("the bar drew its own %s over the system's", name)
		}
	}
	name, ok := tt.Find("MyGo App")
	if !ok {
		t.Fatal("the app name is missing")
	}
	if name.X < 70 {
		t.Errorf("the name sits at x=%v, inside the system controls' strip", name.X)
	}
}

func TestTitleBarClosesHoversInDanger(t *testing.T) {
	tt, _ := titleBarTester(t, TitleBarWindows, TitleBarOptions{Name: "MyGo App"})
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close control is missing")
	}
	// A pixel of the fill, clear of the glyph in the middle.
	x, y := close.Bounds.X+10, close.Bounds.Y+4
	tt.Advance(time.Second)
	if r, g, b := pixel(tt, x, y); r < 240 || g < 240 {
		t.Errorf("the close at rest drew (%d, %d, %d), want the bar's plain background", r, g, b)
	}
	tt.Move(x, y)
	tt.Advance(time.Second)
	r, g, b := pixel(tt, x, y)
	if r < 140 || r < g+40 {
		t.Errorf("the close hovered drew (%d, %d, %d), want a danger shade", r, g, b)
	}
}

func TestTitleBarDimsWhenTheWindowIsInactive(t *testing.T) {
	tt, _ := titleBarTester(t, TitleBarWindows, TitleBarOptions{Name: "MyGo App"})
	minimize, ok := tt.Access("Minimize")
	if !ok {
		t.Fatal("the minimize control is missing")
	}
	x, y := minimize.Bounds.X+minimize.Bounds.W/2, minimize.Bounds.Y+minimize.Bounds.H/2
	tt.Advance(time.Second)
	r1, g1, b1 := pixel(tt, x, y)
	tt.SetWindowActive(false)
	tt.Advance(time.Second)
	r2, g2, b2 := pixel(tt, x, y)
	if r1 == r2 && g1 == g2 && b1 == b2 {
		t.Error("the controls keep their active ink on an inactive window")
	}
}

func TestTitleBarMacDrawsTrafficLights(t *testing.T) {
	tt, f := titleBarTester(t, TitleBarMac, TitleBarOptions{Name: "MyGo App"})

	for _, name := range []string{"Close", "Minimize", "Maximize"} {
		if info, ok := tt.Access(name); !ok {
			t.Fatalf("the %s light is missing", name)
		} else if info.Role != ui.RoleButton {
			t.Errorf("the %s light is not a button (found %v)", name, info.Role)
		}
	}
	// The lights sit on the left, close first, as the system's do.
	close, _ := tt.Access("Close")
	minimize, _ := tt.Access("Minimize")
	maximize, _ := tt.Access("Maximize")
	if close.Bounds.X > minimize.Bounds.X || minimize.Bounds.X > maximize.Bounds.X {
		t.Errorf("the lights are out of order: close x=%v, minimize x=%v, maximize x=%v",
			close.Bounds.X, minimize.Bounds.X, maximize.Bounds.X)
	}
	if close.Bounds.X > 40 {
		t.Errorf("the close light sits at x=%v, not on the left edge", close.Bounds.X)
	}
	if name, ok := tt.Find("MyGo App"); !ok {
		t.Fatal("the app name is missing")
	} else if name.X < maximize.Bounds.X+maximize.Bounds.W {
		t.Errorf("the name at x=%v runs into the lights", name.X)
	}

	// The dots take the light's own color, and the clicks work.
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	dot := func(info ui.AccessInfo) (float32, float32) {
		return info.Bounds.X + info.Bounds.W/2 + 3, info.Bounds.Y + info.Bounds.H/2
	}
	x, y := dot(close)
	tt.Advance(time.Second)
	if r, g, b := pixel(tt, x, y); r < 200 || g > 160 {
		t.Errorf("the close light drew (%d, %d, %d), want its red", r, g, b)
	}
	if err := tt.Click("Close"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.closed {
		t.Error("the close light's click did not reach the app")
	}
	if err := tt.Click("Minimize"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.minimized {
		t.Error("the minimize light's click did not reach the app")
	}
	if err := tt.Click("Maximize"); err != nil {
		t.Fatal(err)
	}
	adv()
	if !f.toggled {
		t.Error("the maximize light's click did not reach the app")
	}
}

func TestTitleBarMacLightsReadGrayWhenInactive(t *testing.T) {
	tt, _ := titleBarTester(t, TitleBarMac, TitleBarOptions{Name: "MyGo App"})
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close light is missing")
	}
	x, y := close.Bounds.X+close.Bounds.W/2+3, close.Bounds.Y+close.Bounds.H/2
	tt.Advance(time.Second)
	r, g, b := pixel(tt, x, y)
	if r < 200 || g > 160 {
		t.Errorf("the active close light drew (%d, %d, %d), want its red", r, g, b)
	}
	tt.SetWindowActive(false)
	tt.Advance(time.Second)
	r, g, b = pixel(tt, x, y)
	if r != g || g != b {
		t.Errorf("the inactive close light drew (%d, %d, %d), want gray", r, g, b)
	}
}

func TestTitleBarMacKeepsItsLightsUnderRTL(t *testing.T) {
	tt, _ := titleBarTester(t, TitleBarMac, TitleBarOptions{Name: "MyGo App"})
	tt.SetDirection(ui.RightToLeft)
	tt.Advance(time.Second)
	close, ok := tt.Access("Close")
	if !ok {
		t.Fatal("the close light is missing")
	}
	if close.Bounds.X > 40 {
		t.Errorf("the close light sits at x=%v; the traffic lights stay on the left", close.Bounds.X)
	}
}

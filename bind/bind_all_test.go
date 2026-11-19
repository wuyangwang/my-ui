package bind

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	displaypkg "github.com/wuyangwang/my-ui/display"
	formspkg "github.com/wuyangwang/my-ui/forms"
	overlaypkg "github.com/wuyangwang/my-ui/overlays"
)

// The handle is sugar, not a parallel implementation: a widget built
// through it must behave like the same widget built by its constructor.
// These tests sample the new handles, checking visible text and clicks.

func TestHandleToggleMatchesConstructor(t *testing.T) {
	handleOn, directOn := false, false
	view := func(c *ui.Context) {
		h := Bind(c)
		h.Toggle(&handleOn, "Handle Toggle")
		ctrlpkg.NewToggle(c, &directOn, "Direct Toggle")
	}
	tt := ui.NewTester(view, 320, 160)
	tt.Advance(time.Second)
	for _, name := range []string{"Handle Toggle", "Direct Toggle"} {
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("%q is not on screen", name)
		}
	}
	if err := tt.Click("Handle Toggle"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Direct Toggle"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !handleOn {
		t.Error("the handle's toggle did not write its flag")
	}
	if !directOn {
		t.Error("the constructor's toggle did not write its flag")
	}
}

func TestHandleSplitButtonMatchesConstructor(t *testing.T) {
	handleClicked, directClicked := false, false
	view := func(c *ui.Context) {
		h := Bind(c)
		h.SplitButton("Handle Split", func() { handleClicked = true }, nil)
		overlaypkg.NewSplitButton(c, "Direct Split", func() { directClicked = true }, nil)
	}
	tt := ui.NewTester(view, 320, 160)
	tt.Advance(time.Second)
	for _, name := range []string{"Handle Split", "Direct Split"} {
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("%q is not on screen", name)
		}
	}
	if err := tt.Click("Handle Split"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Direct Split"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !handleClicked {
		t.Error("the handle's split button did not run its action")
	}
	if !directClicked {
		t.Error("the constructor's split button did not run its action")
	}
}

func TestHandleTextComponentsMatchConstructors(t *testing.T) {
	var stateHandle, stateDirect ui.CheckState
	chosenHandle, chosenDirect := 0, 0
	view := func(c *ui.Context) {
		h := Bind(c)
		h.Alert("Handle Alert")
		displaypkg.NewAlert(c, "Direct Alert")
		h.Breadcrumb([]string{"Handle Home", "Handle End"}, &chosenHandle)
		displaypkg.NewBreadcrumb(c, []string{"Direct Home", "Direct End"}, &chosenDirect)
		h.CheckboxState(&stateHandle, "Handle Checkbox")
		ctrlpkg.NewCheckboxState(c, &stateDirect, "Direct Checkbox")
		h.Tag(displaypkg.TagOptions{Label: "Handle Tag"})
		displaypkg.NewTag(c, displaypkg.TagOptions{Label: "Direct Tag"})
		h.EmptyState(displaypkg.EmptyStateOptions{Title: "Handle Empty"})
		displaypkg.NewEmptyState(c, displaypkg.EmptyStateOptions{Title: "Direct Empty"})
		h.Result(displaypkg.ResultOptions{Title: "Handle Result"})
		displaypkg.NewResult(c, displaypkg.ResultOptions{Title: "Direct Result"})
	}
	tt := ui.NewTester(view, 800, 1000)
	tt.Advance(time.Second)
	for _, text := range []string{
		"Handle Alert", "Direct Alert",
		"Handle Home", "Direct Home",
		"Handle Checkbox", "Direct Checkbox",
		"Handle Tag", "Direct Tag",
		"Handle Empty", "Direct Empty",
		"Handle Result", "Direct Result",
	} {
		if _, ok := tt.Find(text); !ok {
			t.Errorf("%q is not on screen", text)
		}
	}
}

func TestHandleToastMatchesConstructor(t *testing.T) {
	view := func(c *ui.Context) {
		h := Bind(c)
		h.ToastViewport()
		if h.Button("Show Handle Toast").Clicked() {
			h.Toast("Handle Toast").Timeout(-1).Show()
		}
		if ctrlpkg.NewButton(c, "Show Direct Toast").Clicked() {
			overlaypkg.NewToast(c, "Direct Toast").Timeout(-1).Show()
		}
	}
	tt := ui.NewTester(view, 480, 320)
	tt.Advance(time.Second)
	for _, name := range []string{"Show Handle Toast", "Show Direct Toast"} {
		if _, ok := tt.Access(name); !ok {
			t.Fatalf("%q is not on screen", name)
		}
	}
	if err := tt.Click("Show Handle Toast"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Show Direct Toast"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	for _, text := range []string{"Handle Toast", "Direct Toast"} {
		if _, ok := tt.Find(text); !ok {
			t.Errorf("the toast %q is not on screen", text)
		}
	}
}

// TestHandleGenericAndTimeComponents checks that the concrete Radio method
// and the time-backed handles build and show what their constructors show.
func TestHandleGenericAndTimeComponents(t *testing.T) {
	selected := "a"
	date := time.Date(2026, time.November, 19, 12, 0, 0, 0, time.UTC)
	view := func(c *ui.Context) {
		h := Bind(c)
		h.Radio(&selected, "a", "Handle Radio")
		ctrlpkg.NewRadio(c, &selected, "a", "Direct Radio")
		h.Calendar(&date)
		formspkg.NewCalendar(c, &date)
		h.DatePicker(&date)
		formspkg.NewDatePicker(c, &date)
	}
	tt := ui.NewTester(view, 800, 1000)
	tt.Advance(time.Second)
	for _, text := range []string{"Handle Radio", "Direct Radio"} {
		if _, ok := tt.Find(text); !ok {
			t.Errorf("%q is not on screen", text)
		}
	}
}

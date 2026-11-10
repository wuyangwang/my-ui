package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestFieldShowsItsLabelDescriptionAndError(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		}).Description("我们只用它发通知").Required(true).Error("邮箱格式不对")
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)

	if _, ok := tt.Find("邮箱"); !ok {
		t.Error("the label is not on screen")
	}
	if _, ok := tt.Find("我们只用它发通知"); !ok {
		t.Error("the description is not on screen")
	}
	info, ok := tt.Access("邮箱格式不对")
	if !ok {
		t.Fatal("the error message is not on screen")
	}
	// The error is announced as it appears.
	if info.Role != ui.RoleAlert {
		t.Errorf("the error message is %v, want an alert", info.Role)
	}
}

func TestFieldLabelFocusesTheControl(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !tt.Focused("邮箱") {
		t.Error("the label did not focus the control")
	}
}

func TestPasswordInputHidesAndShowsTheText(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewPasswordInput(c, &value).Outline().Label("密码")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("密码"); err != nil {
		t.Fatal(err)
	}
	tt.Type("secret")
	tt.Advance(time.Second)

	info, ok := tt.Access("密码")
	if !ok {
		t.Fatal("the password field is not on screen")
	}
	if info.Value != "" {
		t.Errorf("the hidden password shows %q", info.Value)
	}
	if _, ok := tt.Access("显示密码"); !ok {
		t.Fatal("the button that shows the password is not on screen")
	}
	if err := tt.Click("显示密码"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	info2, ok := tt.Access("密码")
	if !ok {
		t.Fatal("the password field is not on screen")
	}
	if info2.Value != "secret" {
		t.Errorf("the shown password holds %q, want %q", info2.Value, "secret")
	}
}

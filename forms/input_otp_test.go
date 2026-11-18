package forms

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestInputOTPBindsTheValue(t *testing.T) {
	value := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewInputOTP(c, &value, 4)
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if len(value) > 4 {
		t.Errorf("the OTP value %q is too long", value)
	}
	// The first box has the focus on a fresh value; typing fills it and
	// the app's string follows.
	tt.Type("7")
	tt.Advance(time.Second)
	if value != "7" {
		t.Errorf("the OTP value is %q, want 7", value)
	}
}

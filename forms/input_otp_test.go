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

func TestInputOTPFiltersCharacters(t *testing.T) {
	value := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOTP(c, &value, OTPOptions{Length: 4})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	// Letters are dropped from a digits-only code, even pasted together.
	tt.Type("1a2b")
	tt.Advance(time.Second)
	if value != "12" {
		t.Errorf("the OTP value is %q, want 12", value)
	}
}

func TestInputOTPAlphanumericKeepsLetters(t *testing.T) {
	value := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOTP(c, &value, OTPOptions{Length: 4, Alphanumeric: true})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	tt.Type("1a2b")
	tt.Advance(time.Second)
	if value != "1a2b" {
		t.Errorf("the OTP value is %q, want 1a2b", value)
	}
}

func TestInputOTPOnCompleteFiresOnce(t *testing.T) {
	value := ""
	var done []string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOTP(c, &value, OTPOptions{
			Length:     4,
			OnComplete: func(code string) { done = append(done, code) },
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	tt.Type("1234")
	tt.Advance(time.Second)
	if value != "1234" {
		t.Fatalf("the OTP value is %q, want 1234", value)
	}
	if len(done) != 1 || done[0] != "1234" {
		t.Fatalf("OnComplete reported %q, want 1234 once", done)
	}
	// Later frames with the code still complete must not report again.
	tt.Advance(time.Second)
	tt.Advance(time.Second)
	if len(done) != 1 {
		t.Errorf("OnComplete fired %d times, want once: %q", len(done), done)
	}
}

func TestInputOTPDefaultsToSix(t *testing.T) {
	value := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewOTP(c, &value, OTPOptions{})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	tt.Type("123456")
	tt.Advance(time.Second)
	if value != "123456" {
		t.Errorf("the OTP value is %q, want 123456", value)
	}
}

func TestInputOTPSizeTiers(t *testing.T) {
	widthOf := func(size Size) float32 {
		value := ""
		tt := ui.NewTester(func(c *ui.Context) {
			Apply(c, Options{Appearance: AppearanceLight})
			NewOTP(c, &value, OTPOptions{Length: 4, Size: size}).Label("otp")
		}, 320, 120)
		tt.Advance(time.Second)
		info, ok := tt.Access("otp")
		if !ok {
			t.Fatal("the OTP is not on screen")
		}
		return info.Bounds.W
	}
	if small, large := widthOf(SizeXs), widthOf(SizeLg); large <= small {
		t.Errorf("the large cells are not wider: small %v, large %v", small, large)
	}
}

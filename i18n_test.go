package myui

import (
	"testing"
)

func TestMessageFallsBackAlongTheChain(t *testing.T) {
	cases := []struct {
		tag, key, want string
	}{
		{"en", "clear", "Clear"},
		{"zh-Hans", "clear", "清除"},
		{"zh-Hant", "close", "關閉"},
		// A tag with a region falls back to its language.
		{"zh-Hans-CN", "clear", "清除"},
		{"zh-Hant-HK", "close", "關閉"},
		// A language the library has no strings for falls back to
		// English.
		{"ja", "close", "Close"},
		{"fr-CA", "clear", "Clear"},
	}
	for _, c := range cases {
		if got := message(c.tag, c.key); got != c.want {
			t.Errorf("the message %q in %q is %q, want %q", c.key, c.tag, got, c.want)
		}
	}
}

func TestMessageFallsBackToTheKey(t *testing.T) {
	if got := message("en", "nothing"); got != "nothing" {
		t.Errorf("the message of an unknown key is %q, want the key", got)
	}
}

func TestSetMessagesTakesThePlaceOfTheLibrarys(t *testing.T) {
	SetMessages("fr", map[string]string{"close": "Fermer"})
	if got := message("fr", "close"); got != "Fermer" {
		t.Errorf("the message of the app is %q, want %q", got, "Fermer")
	}
	// The library's own strings stay as they were.
	if got := message("en", "close"); got != "Close" {
		t.Errorf("the message of the library is %q, want %q", got, "Close")
	}
}

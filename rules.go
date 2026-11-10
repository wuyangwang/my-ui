package myui

import (
	"regexp"
	"strings"
)

// Rule is one rule of a field's value: it returns the message of the
// error, or nothing when the value holds.
type Rule func(value string) string

// Required is the rule of a value that may not be empty.
func Required(msg string) Rule {
	return func(v string) string {
		if strings.TrimSpace(v) == "" {
			return msg
		}
		return ""
	}
}

// MinLength is the rule of a value that may not be shorter than n.
func MinLength(n int, msg string) Rule {
	return func(v string) string {
		if len([]rune(v)) < n {
			return msg
		}
		return ""
	}
}

// MaxLength is the rule of a value that may not be longer than n.
func MaxLength(n int, msg string) Rule {
	return func(v string) string {
		if len([]rune(v)) > n {
			return msg
		}
		return ""
	}
}

// Pattern is the rule of a value that has to match the pattern.
func Pattern(re *regexp.Regexp, msg string) Rule {
	return func(v string) string {
		if !re.MatchString(v) {
			return msg
		}
		return ""
	}
}

// Email is the rule of a value that has to be an email address.
func Email(msg string) Rule {
	return Pattern(regexp.MustCompile(`^[^@\s]+@[^@\s.]+(\.[^@\s.]+)+$`), msg)
}

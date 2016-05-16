package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func setOrUnset(key, val string) {
	if v := os.Getenv(key); v != "" {
		os.Unsetenv(key)
	} else {
		os.Setenv(key, val)
	}
}

func TestScan(t *testing.T) {
	tests := []struct {
		text     string
		expected string
		setEnv   func()
	}{
		{
			text:     "foo",
			expected: "foo",
		},
		{
			text:     "foo bar baz",
			expected: "foo bar baz",
		},
		{
			text:     "foo ${bar} baz",
			expected: "foo ${bar} baz",
		},
		{
			text:     "foo ${BAR} baz",
			expected: "foo ${BAR} baz",
		},
		{
			text:     "${FOO} bar baz",
			expected: "${FOO} bar baz",
		},
		{
			text:     "${FOO_BAR} bar baz",
			expected: "${FOO_BAR} bar baz",
		},
		{
			text:     "${4OO} bar baz",
			expected: "${4OO} bar baz",
		},
		{
			text:     "${_OO} bar baz",
			expected: "${_OO} bar baz",
		},
		{
			text:     "${FOO BAR} bar baz",
			expected: "${FOO BAR} bar baz",
		},
		{
			text:     "${FOOBAR} bar baz",
			expected: "BAZ bar baz",
			setEnv:   func() { setOrUnset("FOOBAR", "BAZ") },
		},
		{
			text:     "/${FOO}/${BAR}/BAZ/",
			expected: "/foo/bar/BAZ/",
			setEnv:   func() { setOrUnset("FOO", "foo"); setOrUnset("BAR", "bar") },
		},
	}

	for _, test := range tests {
		r := bufio.NewReader(strings.NewReader(test.text))

		// set env
		if test.setEnv != nil {
			test.setEnv()
		}
		got := Scan(r)
		// unset env
		if test.setEnv != nil {
			test.setEnv()
		}

		if got != test.expected {
			t.Errorf("Got: %s, Expected: %s", got, test.expected)
		}
	}
}

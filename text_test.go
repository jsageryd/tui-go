package tui

import "testing"

func TestCols(t *testing.T) {
	for n, tt := range []struct {
		t      Text
		result int
	}{
		{"", 0},
		{"a", 1},
		{"abc", 3},
		{"あ", 2},
		{"あいう", 6},
		{"abcあいう123", 12},
	} {
		if got, want := tt.t.cols(), tt.result; got != want {
			t.Errorf("[%d] t.cols() of %s is %d, want %d", n, tt.t, got, want)
		}
	}
}

func TestTrimRightLen(t *testing.T) {
	for n, tt := range []struct {
		t      Text
		n      int
		result Text
	}{
		{"", 0, ""},
		{"", 1, ""},
		{"", -1, ""},
		{"abc", -1, "abc"},
		{"abc", 0, "abc"},
		{"abc", 1, "ab"},
		{"abc", 2, "a"},
		{"abc", 3, ""},
		{"abc", 4, ""},
		{"あいう", -1, "あいう"},
		{"あいう", 0, "あいう"},
		{"あいう", 1, "あい"},
		{"あいう", 2, "あ"},
		{"あいう", 3, ""},
		{"あいう", 4, ""},
	} {
		if got, want := trimRightLen(tt.t, tt.n), tt.result; got != want {
			t.Errorf("[%d] trimRightLen(%q, %d) = %q, want %q", n, tt.t, tt.n, got, want)
		}
	}
}

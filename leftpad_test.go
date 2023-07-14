package leftpad

import "testing"

func TestLeftpad(t *testing.T) {
	for _, testcase := range []struct {
		s string
		n int
		c byte
		e string
	}{
		{"foo", 5, ' ', "  foo"},
		{"foobar", 6, ' ', "foobar"},
		{"toolong", 3, ' ', "toolong"},
		{"foo", 5, '0', "00foo"},
		{"foobar", 6, '0', "foobar"},
		{"toolong", 3, '0', "toolong"},
		{"negative", -1, ' ', "negative"},
		{"negative", -1, '0', "negative"},
		{"", 0, ' ', ""},
	} {
		if r := Leftpad(testcase.s, testcase.n, testcase.c); r != testcase.e {
			t.Errorf("leftpad(%q, %d, %q) = %q, expected %q", testcase.s, testcase.n, testcase.c, r, testcase.e)
		}
	}
}

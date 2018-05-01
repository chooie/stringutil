package stringutil

import (
	"testing"

	"github.com/chooie/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		reversed := Reverse(c.in)
		assert.Equal(
			t,
			reversed,
			c.want,
			"Should be reversed",
			"> Reverse(\""+c.in+"\")",
		)
	}
}

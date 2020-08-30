package tests

import (
	"testing"

	"waheedsworkbench.com/toolbox/tools"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{16, 4},
		{9, 3},
		{4, 2},
		{3, 1.73205},
		{2, 1.41421},
		{1, 1},
	}
	for _, c := range cases {
		got := tools.Sqrt(c.in)
		if got != c.want {
			t.Errorf("Sqrt(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}

package multiply_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	ts := []struct {
		Num1 string
		Num2 string
		Want string
	}{
		{
			Num1: "123",
			Num2: "45",
			Want: "5535",
		},
	}

	for _, tc := range ts {
		got := multiply(tc.Num1, tc.Num2)
		assert.Equal(t, tc.Want, got)
	}
}

func TestStrPlus(t *testing.T) {
	ts := []struct {
		s1   string
		s2   string
		want string
	}{
		{s1: "1", s2: "2", want: "3"},
		{s1: "1", s2: "12", want: "13"},
		// {s1: "123", s2: "1123", want: "1246"},
	}

	for _, tc := range ts {
		got := strPlus(tc.s1, tc.s2)
		assert.Equal(t, tc.want, got)
	}
}

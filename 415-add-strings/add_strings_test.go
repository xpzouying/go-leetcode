package add_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {

	ts := []struct {
		n1   string
		n2   string
		want string
	}{
		{n1: "1", n2: "12", want: "13"},
		{n1: "9", n2: "9", want: "18"},
	}

	for _, tc := range ts {
		got := addStrings(tc.n1, tc.n2)
		assert.Equal(t, tc.want, got)
	}
}

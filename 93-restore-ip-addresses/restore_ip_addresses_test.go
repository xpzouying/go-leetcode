package restore_ip_addresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreIpAddresses(t *testing.T) {

	ts := []struct {
		Input string
		Want  []string
	}{
		{
			Input: "25525511135",
			Want:  []string{"255.255.11.135", "255.255.111.35"},
		},
	}

	for _, tc := range ts {
		got := restoreIpAddresses(tc.Input)
		assert.Equal(t, tc.Want, got)
	}

}

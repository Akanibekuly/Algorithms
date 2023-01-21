package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	s := "25525511135"
	assert.Equal(t, 255, toInt(s[0:3]))
}

func TestRestoreIpAddresses(t *testing.T) {
	cases := []struct {
		name        string
		s           string
		expectedRes []string
	}{
		{
			name:        "1",
			s:           "25525511135",
			expectedRes: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name:        "2",
			s:           "0000",
			expectedRes: []string{"0.0.0.0"},
		},
		{
			name:        "3",
			s:           "101023",
			expectedRes: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(
				tt,
				c.expectedRes,
				restoreIpAddresses(c.s),
				"case %d", i,
			)
		})
	}
}

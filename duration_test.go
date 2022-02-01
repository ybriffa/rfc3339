package rfc3339

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFormatDuration(t *testing.T) {
	for _, test := range []struct {
		expected string
		value    time.Duration
	}{
		// 0
		{
			expected: "PT23H12M42S",
			value:    23*time.Hour + 12*time.Minute + 42*time.Second,
		},

		// 1
		{
			expected: "PT2H",
			value:    2 * time.Hour,
		},

		// 2
		{
			expected: "P2DT1H",
			value:    49 * time.Hour,
		},

		//3
		{
			expected: "PT0S",
			value:    0,
		},

		//4
		{
			expected: "-PT2H",
			value:    -2 * time.Hour,
		},

		// 5
		{
			expected: "P2W1DT1H",
			value:    (2 * 7 * 24 * time.Hour) + (24 * time.Hour) + time.Hour,
		},
	} {
		result := FormatDuration(test.value)
		require.Equal(t, test.expected, result)
	}
}

func TestParseDuration(t *testing.T) {
	for _, test := range []struct {
		expected      time.Duration
		value         string
		errorExpected bool
	}{
		// 0
		{
			value:    "P3W1DT23H12M42S",
			expected: 3*7*24*time.Hour + 1*24*time.Hour + 23*time.Hour + 12*time.Minute + 42*time.Second,
		},

		// 1
		{
			value:    "P2D",
			expected: 2 * 24 * time.Hour,
		},

		// 2
		{
			value:    "PT42M",
			expected: 42 * time.Minute,
		},
	} {
		result, err := ParseDuration(test.value)
		if test.errorExpected {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expected, result)
	}
}

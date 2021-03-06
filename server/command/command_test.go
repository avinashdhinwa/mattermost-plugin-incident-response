package command

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTimeForMillis(t *testing.T) {
	t.Run("Now", func(t *testing.T) {
		expected := time.Now()

		millis := expected.Unix() * 1000
		actual := getTimeForMillis(millis)

		require.Equal(t, expected.Unix(), actual.Unix())
	})

	t.Run("Epoch 0", func(t *testing.T) {
		expected, err := time.Parse("2006-01-02 15:04:05 MST", "1970-01-01 00:00:00 UTC")
		require.NoError(t, err)

		actual := getTimeForMillis(0)

		require.Equal(t, expected.Unix(), actual.Unix())
	})

	t.Run("Go layout", func(t *testing.T) {
		// From https://golang.org/pkg/time/#pkg-constants:
		// "The reference time used in the layouts is the specific time:
		// Mon Jan 2 15:04:05 MST 2006
		// which is Unix time 1136239445"
		var unixTime int64 = 1136239445
		layout := "Mon Jan 02 15:04:05 -0700 2006"

		expected, err := time.Parse(layout, layout)
		require.NoError(t, err)

		actual := getTimeForMillis(unixTime * 1000)

		require.Equal(t, expected.Unix(), actual.Unix())
	})
}

func TestDurationString(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected string
	}{
		{
			name:     "Duration zero",
			start:    now,
			end:      now,
			expected: "0s",
		},
		{
			name:     "Only seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 10, 0, 25, 0, time.UTC),
			expected: "25s",
		},
		{
			name:     "Exact minutes",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 10, 15, 0, 0, time.UTC),
			expected: "15m0s",
		},
		{
			name:     "Minutes and seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 10, 30, 25, 0, time.UTC),
			expected: "30m25s",
		},
		{
			name:     "Exact hours",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 13, 0, 0, 0, time.UTC),
			expected: "3h0m0s",
		},
		{
			name:     "Hours and minutes",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 12, 45, 0, 0, time.UTC),
			expected: "2h45m0s",
		},
		{
			name:     "Hours, minutes and seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 1, 20, 59, 10, 0, time.UTC),
			expected: "10h59m10s",
		},
		{
			name:     "Exact days",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 2, 10, 0, 0, 0, time.UTC),
			expected: "1d0s",
		},
		{
			name:     "Days and seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 3, 10, 0, 25, 0, time.UTC),
			expected: "2d25s",
		},
		{
			name:     "Days and exact minutes",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 5, 10, 15, 0, 0, time.UTC),
			expected: "4d15m0s",
		},
		{
			name:     "Days, minutes and seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 10, 10, 30, 25, 0, time.UTC),
			expected: "9d30m25s",
		},
		{
			name:     "Days and hours",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 21, 13, 0, 0, 0, time.UTC),
			expected: "20d3h0m0s",
		},
		{
			name:     "Days, hours and minutes",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 26, 12, 45, 0, 0, time.UTC),
			expected: "25d2h45m0s",
		},
		{
			name:     "Days, hours, minutes and seconds",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 1, 31, 20, 59, 10, 0, time.UTC),
			expected: "30d10h59m10s",
		},
		{
			name:     "Days, hours, minutes and seconds over months",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2000, 2, 31, 20, 59, 10, 0, time.UTC),
			expected: "61d10h59m10s",
		},
		{
			name:     "Days, hours, minutes and seconds over years",
			start:    time.Date(2000, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2001, 2, 31, 20, 59, 10, 0, time.UTC),
			expected: "427d10h59m10s",
		},
		{
			name:     "An exact year",
			start:    time.Date(2001, 1, 1, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2002, 1, 1, 10, 0, 0, 0, time.UTC),
			expected: "365d0s",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := durationString(testCase.start, testCase.end)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

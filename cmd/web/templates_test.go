package main

import (
	"testing"
	"time"

	"chetraseng.com/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string

		tm time.Time

		want string
	}{

		{
			name: "UTC",
			tm:   time.Date(2025, 5, 6, 23, 45, 50, 0, time.UTC),
			want: "06 May 2025 at 23:45",
		},
		{name: "Empty", tm: time.Time{}, want: ""},
		{
			name: "CET",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, humanDate(test.tm), test.want)
		})
	}
}

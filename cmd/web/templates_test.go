package main

import (
	"testing"
	"time"

	"github.com/sam-maton/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 04, 20, 12, 35, 0, 0, time.UTC),
			want: "20 Apr 2024 at 12:35",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 04, 20, 12, 35, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "20 Apr 2024 at 11:35",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}

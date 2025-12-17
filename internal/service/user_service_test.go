package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		dob  time.Time
		want int
	}{
		{
			name: "birthday already passed this year",
			dob:  time.Date(now.Year()-25, now.Month()-1, 1, 0, 0, 0, 0, time.UTC),
			want: 25,
		},
		{
			name: "birthday yet to occur this year",
			dob:  time.Date(now.Year()-25, now.Month()+1, 1, 0, 0, 0, 0, time.UTC),
			want: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(tt.dob)
			if got != tt.want {
				t.Errorf("CalculateAge() = %d, want %d", got, tt.want)
			}
		})
	}
}

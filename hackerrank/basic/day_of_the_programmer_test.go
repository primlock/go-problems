package basic

import "testing"

func TestDayOfProgrammer(t *testing.T) {
	testCases := []struct {
		name     string
		year     int32
		expected string
	}{
		{name: "1", year: 1984, expected: "12.09.1984\n"},
		{name: "2", year: 2017, expected: "13.09.2017\n"},
		{name: "3", year: 2016, expected: "12.09.2016\n"},
		{name: "3", year: 1800, expected: "12.09.1800\n"},
		{name: "4", year: 1918, expected: "26.09.1918\n"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := DayOfProgrammer(tt.year)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

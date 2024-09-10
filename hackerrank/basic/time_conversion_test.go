package basic

import "testing"

func TestTimeConversion(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "1", s: "07:05:45PM", expected: "19:05:45"},
		{name: "2", s: "12:25:00PM", expected: "12:25:00"},
		{name: "3", s: "11:56:20PM", expected: "23:56:20"},
		{name: "4", s: "12:01:00AM", expected: "00:01:00"},
		{name: "5", s: "06:45:00AM", expected: "06:45:00"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeConversion(tt.s)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

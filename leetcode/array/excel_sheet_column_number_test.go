package array

import "testing"

func TestTitleToNumber(t *testing.T) {
	testCases := []struct {
		name        string
		columnTitle string
		expected    int
	}{
		{name: "1", columnTitle: "A", expected: 1},
		{name: "2", columnTitle: "AB", expected: 28},
		{name: "3", columnTitle: "ZY", expected: 701},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TitleToNumber(tt.columnTitle)
			want := tt.expected

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

package other

import "testing"

func TestGetMaxTotalArea(t *testing.T) {
	testCases := []struct {
		name        string
		sideLengths []int32
		expected    int32
	}{
		{name: "1", sideLengths: []int32{2, 6, 6, 2, 3, 5}, expected: 12},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMaxTotalArea(tt.sideLengths)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

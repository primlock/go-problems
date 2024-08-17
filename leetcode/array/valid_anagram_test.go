package array

import "testing"

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{name: "1", s: "anagram", t: "nagaram", expected: true},
		{name: "2", s: "rat", t: "car", expected: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidAnagram(tt.s, tt.t)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

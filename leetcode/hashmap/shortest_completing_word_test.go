package hashmap

import "testing"

func TestShortestCompletingWord(t *testing.T) {
	testCases := []struct {
		name         string
		licensePlate string
		words        []string
		expected     string
	}{
		{name: "1", licensePlate: "1s3 PSt", words: []string{"step", "steps", "stripe", "stepple"}, expected: "steps"},
		{name: "2", licensePlate: "1s3 456", words: []string{"looks", "pest", "stew", "show"}, expected: "pest"},
		{name: "3", licensePlate: "GrC8950", words: []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}, expected: "according"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestCompletingWord(tt.licensePlate, tt.words)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

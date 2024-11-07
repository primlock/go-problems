package queue

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	testCases := []struct {
		name     string
		senate   string
		expected string
	}{
		{name: "1", senate: "RD", expected: "Radiant"},
		{name: "2", senate: "RDD", expected: "Dire"},
		{name: "3", senate: "RRR", expected: "Radiant"},
		{name: "4", senate: "DDRRR", expected: "Dire"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PredictPartyVictory(tt.senate)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

package hashmap

import "testing"

func TestMostCommonWord(t *testing.T) {
	testCases := []struct {
		name      string
		paragraph string
		banned    []string
		expected  string
	}{
		{name: "1", paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.", banned: []string{"hit"}, expected: "ball"},
		{name: "2", paragraph: "a.", banned: []string{}, expected: "a"},
		{name: "3", paragraph: "a, a, a, a, b,b,b,c, c", banned: []string{"a"}, expected: "b"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MostCommonWord(tt.paragraph, tt.banned)
			want := tt.expected

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

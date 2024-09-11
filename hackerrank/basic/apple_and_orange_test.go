package basic

import "testing"

func TestAppleAndOrange(t *testing.T) {
	testCases := []struct {
		name    string
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}{
		{name: "1", s: 7, t: 11, a: 5, b: 15, apples: []int32{-2, 2, 1}, oranges: []int32{5, -6}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			CountApplesAndOranges(tt.s, tt.t, tt.a, tt.b, tt.apples, tt.oranges)
		})
	}
}

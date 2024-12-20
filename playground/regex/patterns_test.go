package playground

import "testing"

func TestNonAlphanumericFilter(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "Remove all the non-alphas", s: "zFtVWb&gF#N@MVxf55tKn&oyN", expected: "zFtVWbgFNMVxf55tKnoyN"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NonAlphanumericFilter(tt.s)

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

package playground

import "regexp"

// The carret means match any character other than what's defined.
var nonAlphaReg = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func NonAlphanumericFilter(s string) string {
	return nonAlphaReg.ReplaceAllString(s, "")
}

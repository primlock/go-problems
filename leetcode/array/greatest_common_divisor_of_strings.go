// https://leetcode.com/problems/greatest-common-divisor-of-strings

package array

func GcdOfStrings(str1 string, str2 string) string {
	// Check if the two strings have a common GCD
	if str1+str2 != str2+str1 {
		return ""
	}

	// Get the GCD of two ints with the Euclidean algroithm
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Return a substring that's length is equal to the GCD of str1 and str2
	return str1[:gcd(len(str1), len(str2))]
}

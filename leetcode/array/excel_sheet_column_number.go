// https://leetcode.com/problems/excel-sheet-column-number/

package array

func TitleToNumber(columnTitle string) int {
	// Each position is 26 times more significant than the position to its right
	columnNumber := 0

	for _, letter := range []byte(columnTitle) {
		// Multiply the current total by 26 to make space for the value of the next character.
		columnNumber = columnNumber*26 + int(letter-'@')
	}

	return columnNumber
}

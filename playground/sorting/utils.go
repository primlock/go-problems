package playground

func Swap(array []int, a, b int) []int {
	if array == nil {
		return nil
	}

	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp

	return array
}

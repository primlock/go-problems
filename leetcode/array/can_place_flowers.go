package array

func CanPlaceFlowers(flowerbed []int, n int) bool {
	remaining := n

	for i, plot := range flowerbed {
		// No more flowers left to plant
		if remaining == 0 {
			break
		}

		// We can only plant in an empty plot (aka 0)
		if plot == 0 {
			// Only one plot in the flowerbed
			if len(flowerbed) == 1 {
				flowerbed[i] = 1
				remaining--
				continue
			}

			// Guard the start and end of array for out of bounds
			if i == 0 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					remaining--
				}
				continue
			} else if i == len(flowerbed)-1 {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					remaining--
				}
				continue
			}

			// Check the neighbors to see if they are occupied
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				remaining--
			}
		}
	}

	return remaining == 0
}

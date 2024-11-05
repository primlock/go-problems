// https://leetcode.com/problems/asteroid-collision

package stack

func AsteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, asteroid := range asteroids {
		collision := false

		// Determine if left moving asteroid collides where stack top is a right moving asteroid.
		for asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			// Both asteroids are the same size
			if stack[len(stack)-1] == -asteroid {
				collision = true
				// Top asteroid explodes
				stack = stack[:len(stack)-1]
				break
			} else if stack[len(stack)-1] > -asteroid {
				// Top is larger so this asteroid explodes
				collision = true
				break
			} else {
				// Top is smaller so it explodes
				stack = stack[:len(stack)-1]
			}
		}

		if !collision {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

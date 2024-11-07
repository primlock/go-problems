// https://leetcode.com/problems/dota2-senate

package queue

func PredictPartyVictory(senate string) string {
	if len(senate) == 1 {
		if senate == "R" {
			return "Radiant"
		} else {
			return "Dire"
		}
	}

	// Two queues to store the voters from the each party.
	radiant := make([]int, 0)
	dire := make([]int, 0)
	n := len(senate)

	// Add the voting order of each senator to the respective queue.
	for order, senator := range senate {
		if senator == 'R' {
			radiant = append(radiant, order)
		} else {
			dire = append(dire, order)
		}
	}

	// Let the senators exercise their rights until a senator can announce the victory.
	for len(radiant) > 0 && len(dire) > 0 {
		// Bring forward the voter from each queue
		r := radiant[0]
		d := dire[0]

		radiant = radiant[1:]
		dire = dire[1:]

		// The voter with the lower order get to ban the other
		if r < d {
			radiant = append(radiant, r+n)
		} else {
			dire = append(dire, d+n)
		}
	}

	// Announce the victory
	if len(radiant) == 0 {
		return "Dire"
	}

	return "Radiant"
}

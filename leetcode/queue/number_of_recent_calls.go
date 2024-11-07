// https://leetcode.com/problems/number-of-recent-calls

package queue

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	// Push the new request into the queue
	this.queue = append(this.queue, t)

	// Return the number of requests that have happend in the last 3000ms
	count := 0
	for _, request := range this.queue {
		if request >= t-3000 {
			count++
		}
	}

	return count
}

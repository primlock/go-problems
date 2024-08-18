# Frequency Map Algorithm
The frequency map algorithm is useful in situations where you need to keep track of how many times you've encountered an element. Implemented as a hash map, the frequency map keeps a counter at each key for how many times the element has been seen. 

Using a frequency map will increase the required space to O(N) in the event that every element is stored in the map.

## Implementation of a Frequency Map
A typical frequency map is a hash map using a key with the same data type as the elements you want to track (string, integer, double, etc) and a value of type integer acting as the counter.

If the key **does not** exist in the map, add the key with an inital value of 1, otherwise, increment the value at key by 1.

### Sample Code
```
inputs := []int {1, 2, 3, 1}

freq := make(map[int]int)

for _, num := range inputs {
    freq[num]++
}

```
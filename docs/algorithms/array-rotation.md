# Performing Rotations on an Array
Array rotations are when elements are shifted to the left or to the right of their starting index. We say we are performing a **left** rotation when the elements of the array move toward index $0$ and wrap back around to index $n-1$. Conversely, a **right** rotation is when elements of an array move toward the last index and wrap around to index $0$.
```
Left: Elements 2 and 3 shifts to the left by 1 index and element 1 wraps around to index n-1.
[1, 2, 3] -> [2, 3, 1]

Right: Elements 1 and 2 shift to the right by 1 index and element 3 wraps around to index 0.
[1, 2, 3] -> [3, 1, 2]
```

## Don't do more work than you need to
Shifting every element to the left or to the right can become expensive (in terms of time) fast. To perform just one rotation without optimizations we need to iterate over the entire array, shifting each element to the left or to the right. We then need to repeat this process for the number of rotations required.

The trick is not to do any more work than what you need. This means finding out the **minimum** number of rotations needed to place the array in the expected order. If array is of size $n$ and the number of rotations is $k$, performing $n = k$ rotations puts the array back in the original order that you started with. To find this minumum we need the remainder of the number of rotaions divided by the size of the array:

$$
k = k \% n
$$

In programming, this is achieved using the modulo operator in the expression: `k = k % n`. Now we have the minimum number of *left* rotations we need to perform in order to rotate our elements in the expected order. To find the minimum number of *right* rotations, we need to perform an extra step which is to subtract the size of the array by number of rotations we need to perform, `x = n - k`. This has the effect of rotating left until we are at the position we would be in if we went right.

## Use a Temporary Array
There are multiple way to do the rotation, some more space efficient than others but I find the simplest is to create a temp array and append to it. Once we have our array, we are going to begin adding elements to it starting from $i = k$. When we reach the end of the original array, we being another loop to start from index $0$ and go until we reach index $k$. The result is a temp array with the order of the original array rotated $k$ times.

## Left Rotation Sample Code 
```
// nums = [1, 2, 3, 4, 5], k = 3

n := len(nums)
k := k % n

temp := make([]int, n)

for i := k; i < n; i++ {
    temp = append(temp, nums[i])
}

for i := 0; i < k; i++ {
    temp = append(temp, nums[i])
}

```

## Right Rotation Sample Code 
```
// nums = [1, 2, 3, 4, 5], k = 3

n := len(nums)
k := k % n
rotations := n - k

temp := make([]int, n)

for i := rotations; i < n; i++ {
    temp = append(temp, nums[i])
}

for i := 0; i < rotations; i++ {
    temp = append(temp, nums[i])
}

```
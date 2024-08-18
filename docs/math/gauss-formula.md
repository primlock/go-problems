# Gauss's Summation Formula
Johann Carl Friedrich Gauss was a German mathematician, astronomer, geodesist, and physicist who contributed to many fields in mathematics and science.

Say we need to add the numbers 1 through 10. We could take the approach of summing each number as 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 to get our result of 55. 

What if instead we lined up the numbers in two rows where the first is ordered 1 to 10 and the second is ordered 10 to 1.

```
1  2  3  4  5  6  7  8  9  10
10 9  8  7  6  5  4  3  2  1 
```

Each one of theses columns adds up to 11. The total of all the numbers above is the number of pairs (10) multiplied by the sum of each pair (11). We only want the sum of one row so we need to divide our answer by 2.

## Organize the terms
We can use $n$ to represent how many numbers are in the array. $n$ is the largest number and we now know that 1 is the other number that makes the pair $(n + 1)$. Since we are using 2 rows to make this pair, we need to divide this number by 2. This formula can be written as:

$$
\sum = \frac{n * (n + 1)}{2}
$$

### References
- Let's Talk Science. (2021). *Gauss Summation*. https://letstalkscience.ca/educational-resources/backgrounders/gauss-summation
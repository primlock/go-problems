# Euclidean Algorithm

The Euclidean algorithm is an efficient method for finding the Greatest Common Divisor (GCD) of two integers. The GCD of two numbers is the largest integer that divides both of them without leaving a remainder. 

The Euclidean algorithm works by repeatedly replacing the larger number with the remainder when the larger number is divided by the smaller number, until the remainder is zero. At that point, the last non-zero remainder is the GCD.

## Steps of the Euclidean Algorithm

Here is how you would use the Euclidean algorithm to find the GCD of two integers.

1. Start with two integers, $a$ and $b$ (where $a > b$).
2. Divide $a$ by $b$ and get the remainder $r$.
3. Replace $a$ with $b$ and $b$ with $r$.
4. Repeat steps 2 and 3 until $r$ becomes 0.
5. When $r$ is 0, $b$ (the last non-zero remainder) is the GCD of the original two numbers.

## Example

To find the GCD of 48 and 18.

1. $48 ÷ 18 = 2$ remainder **12**.
2. Replace $a$ with 18 and $b$ with 12.
3. $18 ÷ 12 = 118 ÷ 12 = 1$ remainder **6**.
4. Replace $a$ with 12 and $b$ with 6.
5. $12 ÷ 6 = 212 ÷ 6 = 2$ remainder **0**.

## Code

This is how you could implement the Euclidean algorithm in code.

```go
package main

import "fmt"

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    a := 48
    b := 18

    // Output: 6
    fmt.Println(gcd(a, b))
}
```

### References
- Math is Fun. (n.d.). *Euclidean Algorithm*. https://www.mathsisfun.com/numbers/euclidean-algo.html

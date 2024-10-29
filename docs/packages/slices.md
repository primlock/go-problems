# Slices

The [slices](https://pkg.go.dev/slices) package from the standard library provides various functions useful with slices of any type.

## Reverse

```go
func Reverse[S ~[]E, E any](s S)
```

Reverse reverses the elements of the slice in place. 

<details>

<summary>Example</summary>

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	slices.Reverse(names)
	fmt.Println(names)
}
```
```
Output:

[VERA Bob alice]
```
</details>

## Sort

```go
func Sort[S ~[]E, E cmp.Ordered](x S)
```

Sort sorts a slice of any ordered type in ascending order. When sorting floating-point numbers, NaNs are ordered before other values.
* If you would like descending order, call **Reverse** after calling **Sort**.

<details>

<summary>Example</summary>

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	smallInts := []int8{0, 42, -10, 8}
	slices.Sort(smallInts)
	fmt.Println(smallInts)
}
```
```
Output:

[-10 0 8 42]
```

</details>
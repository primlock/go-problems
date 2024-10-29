# Strings 

The [strings](https://pkg.go.dev/strings) package from the standard library provides functions to manipulate UTF-8 encoded strings.

## Fields

```go
func Fields(s string) []string
```

Fields splits the string `s` around each instance of one or more consecutive white space characters, as defined by `unicode.IsSpace`, returning a slice of substrings of `s` or an empty slice if `s` contains only white space.

<details>

<summary>Example</summary>

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
```
```
Output:

Fields are: ["foo" "bar" "baz"]
```

</details>

## Split

```go
func Split(s, sep string) []string
```

Split slices `s` into all substrings separated by `sep` and returns a slice of the substrings between those separators. 

* If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
* If sep is empty, Split splits after each UTF-8 sequence.
* If both s and sep are empty, Split returns an empty slice. 

<details>

<summary>Example</summary>

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
```
```
Output:

["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```

</details>
## Encapsulation in Golang

**Encapsulation:** Encapsulation is defined as the wrapping up of data into a single unit. In a different way, encapsulation is a protective shield that prevents the data from being accessed by the code outside this shield.

Golang provides two types of identifiers:
- Exported identifiers
- Unexported identifiers

### Exported identifiers
- Identifiers which are exported by the package.
- First letter is always capital.
- Can be used by a different package by importing that package.
- Just exports the name and not implementation.

**Example**

```go
// Go program to illustrate
// the concept of encapsulation
// using exported function
package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating a slice of strings
	slc := []string{"GeeksforGeeks", "geeks", "gfg"}

	// Convert the case of the
	// elements of the given slice
	// Using ToUpper() function
	for x := 0; x < len(slc); x++ {

		// Exported Method
		res := strings.ToUpper(slc[x])

		// Exported Method
		fmt.Println(res)
	}
}
```

### Unexported identifiers
- Identifiers that are not exported from any package.
- First letter is always small.
- Visibility is limited to that package.

**Example**
```go
// Go program to illustrate
// the unexported function
package main

import "fmt"

// The addition function returns
// the sum of the elements
// Unexported function
func addition(val ...int) int {
	s := 0

	for x := range val {
		s += val[x]
	}

	fmt.Println("Total Sum: ", s)
	return s
}

// Main function
func main() {

	addition(23, 546, 65, 42, 21, 24, 67)
}
```

### Benefits of encapsulation
- Hiding implementation details from the user.
- Increase the reusability of the code.
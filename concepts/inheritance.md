## Inheritance in Golang

Golang doesn't support classes, so inheritance in go takes place through struct embedding. We cannot directly extend structs but rather use a concept called **composition** where the struct is used to form other objects. In composition, base structs can be embedded into a child struct and the methods of the base struct can be directly called on the child struct.

**Example**

```go
// Golang program to illustrate the
// concept of inheritance
package main

import (
	"fmt"
)

// declaring a struct
type Comic struct{

	// declaring struct variable
	Universe string
}

// function to return the
// universe of the comic
func (comic Comic) ComicUniverse() string {

	// returns comic universe
	return comic.Universe
}

// declaring a struct
type Marvel struct{
	
	// anonymous field,
	// this is composition where
	// the struct is embedded
	Comic
}

// declaring a struct
type DC struct{
	
	// anonymous field
	Comic
}

// main function
func main() {
	
	// creating an instance
	c1 := Marvel{
	
		// child struct can directly
		// access base struct variables
			Comic{
			Universe: "MCU",
			},
		}
	
			
	// child struct can directly
	// access base struct methods
	
	// printing base method using child
		fmt.Println("Universe is:", c1.ComicUniverse())	
	
	c2 := DC{
		Comic{
			Universe : "DC",
		},
	}
	
	// printing base method using child
	fmt.Println("Universe is:", c2.ComicUniverse())
}

```

### Refs

- https://www.geeksforgeeks.org/inheritance-in-golang/
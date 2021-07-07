// package interfaces

// import (
// 	"fmt"
// )

// func RunInterfaceProgram() {
// 	chicken1 := &Chicken{
// 		eggCount: 0,
// 	}

// }

// type Bird interface {
// 	layEggs()
// }

// type Chicken struct {
// 	eggCount int
// }

// func (c *Chicken) layEggs() error {
// 	return nil
// }

// type Egg struct {}

// type (e Egg) hatch(b Bird) (b Bird, error) {
// 	if b.eggCount > 0 {
// 		return nil, fmt.Errof("Egg already hatched")
// 	}
// 	return nil, nil
// }
package align

import (
	"fmt"
)

func ExampleNavigator_Right() {
	nav := NewNavigator(0, 0, 0, 0, 2, 2)
	for x, y, more := nav.Right(); more; x, y, more = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 0,0
	// 1,0
	// 2,0
	// 0,1
	// 1,1
	// 2,1
	// 0,2
	// 1,2
	// 2,2
}

func ExampleNavigator_Up() {
	nav := NewNavigator(1, 1, 0, 0, 1, 1)
	for x, y, more := nav.Up(); more; x, y, more = nav.Up() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 1,0
	// 0,1
	// 0,0
}

func ExampleNavigator_Left() {
	nav := NewNavigator(1, 1, 0, 0, 1, 1)
	for x, y, more := nav.Left(); more; x, y, more = nav.Left() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 0,1
	// 1,0
	// 0,0
}

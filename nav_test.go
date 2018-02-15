package align

import (
	"fmt"
)

func ExampleNavigator() {
	x, y := 10, 5
	ix := make([][]int, y)
	for i := range ix {
		ix[i] = make([]int, x)
	}

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

package align_test

import (
	"fmt"

	"github.com/gregoryv/align"
)

func ExampleCompare() {
	fmt.Println(
		align.Compare(
			"GCATGCUAAAAAAA",
			"GATTACAAAAAAAABBB",
		),
	)
	fmt.Println(
		align.Compare(
			"ABC",
			"ABC",
		),
	)
	// output:
	// 0.78571427
	// 1
}


func Example_fuzzyMatchWithCompare() {
	a := "GCATGCUAAAAAAA"
	b := "GATTACAAAAAAAABBB"
	if align.Compare(a, b) > 0.7 {
		fmt.Println("similar enough")
	}
}

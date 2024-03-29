package align_test

// ExampleNeedlemanWunsch as described on
// https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm

import (
	"os"
	"testing"

	"github.com/gregoryv/align"
)

func ExampleResult_PrintAlignment() {
	a := []rune("GCATGCU")
	b := []rune("GATTACA")
	result := align.NeedlemanWunsch(a, b)
	result.PrintAlignment(os.Stdout)
	// output:
	// GCA-TGCU
	// G-ATTACA
	//
	// GCAT-GCU
	// G-ATTACA
	//
	// GCATG-CU
	// G-ATTACA
}

func ExampleResult_PrintScoreMatrix() {
	a := []rune("GCATGCU")
	b := []rune("GATTACA")
	result := align.NeedlemanWunsch(a, b)
	result.PrintScoreMatrix(os.Stdout)

	// output:
	//       G  C  A  T  G  C  U
	//    0 -1 -2 -3 -4 -5 -6 -7
	// G -1  1  0 -1 -2 -3 -4 -5
	// A -2  0  0  1  0 -1 -2 -3
	// T -3 -1 -1  0  2  1  0 -1
	// T -4 -2 -2 -1  1  1  0 -1
	// A -5 -3 -3 -1  0  0  0 -1
	// C -6 -4 -2 -2 -1 -1  1  0
	// A -7 -5 -3 -1 -2 -2  0  0
}

/* Origin arrows are represented using three bits, 0 = No origin, 1 = diagonal, 2 = left and 4 is top.
 */
func ExampleResult_PrintOrigins() {
	a := []rune("GCATGCU")
	b := []rune("GATTACA")
	result := align.NeedlemanWunsch(a, b)
	result.PrintOrigins(os.Stdout)

	// output:
	//       G  C  A  T  G  C  U
	//    0  0  0  0  0  0  0  0
	// G  0  1  2  2  2  3  2  2
	// A  0  4  1  1  2  2  2  2
	// T  0  4  5  4  1  2  2  2
	// T  0  4  5  4  5  1  3  3
	// A  0  4  5  1  4  5  1  3
	// C  0  4  1  4  4  5  1  2
	// A  0  4  4  1  6  5  4  1
}

func ExampleResult_PrintScoreMatrix_global() {
	a := []rune("GCATGCU")
	b := []rune("GAT")
	result := align.NeedlemanWunsch(a, b)
	result.PrintScoreMatrix(os.Stdout)

	// output:
	//       G  C  A  T  G  C  U
	//    0 -1 -2 -3 -4 -5 -6 -7
	// G -1  1  0 -1 -2 -3 -4 -5
	// A -2  0  0  1  0 -1 -2 -3
	// T -3 -1 -1  0  2  1  0 -1
}

func TestResult_MaxScore(t *testing.T) {
	cases := []struct {
		a, b                    []rune
		match, miss, indel, ext int
		exp                     int
	}{
		{[]rune("GCATGCU"), []rune("GATTACA"), 1, -1, -1, -1, 0},
		{[]rune("GCATGCU"), []rune("GATTACA"), 1, -1, -2, -1, -1},
		{[]rune("GCCU"), []rune("GCTTCU"), 1, -1, -2, -1, 1},
		{[]rune("GCTTCU"), []rune("GCCU"), 1, -1, -2, -1, 1},
		{[]rune("AAA"), []rune("AAA"), 1, -1, -2, -1, 3},
	}

	for i, c := range cases {
		result := align.NeedlemanWunschCustom(c.a, c.b, c.match, c.miss, c.indel, c.ext)
		res := result.MaxScore()
		if res != c.exp {
			t.Errorf("%v. expected to return %v, got %v\n", i, c.exp, res)
			result.PrintScoreMatrix(os.Stdout)
			result.PrintOrigins(os.Stdout)
			result.PrintAlignment(os.Stdout)
		}
	}
}

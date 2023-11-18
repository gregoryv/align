package align_test

import (
	"testing"

	"github.com/gregoryv/align"
)

func BenchmarkRunes(bench *testing.B) {
	a := []rune("GCATGCUGAGACCAC")
	b := []rune("GATTACAAGACGAAC")
	for i := 0; i < bench.N; i++ {
		align.NeedlemanWunsch(a, b)
	}
}

// Compare to nwalgo
func BenchmarkRunes_nwalgo_sequences(bench *testing.B) {
	a := []rune("GGAATTAATCCAGGTAATGGACCCCAAGAT")
	b := []rune("GCCAGGATTCCCAGATATGGCCAAGGTTCC")
	for i := 0; i < bench.N; i++ {
		align.NeedlemanWunsch(a, b)
	}
}

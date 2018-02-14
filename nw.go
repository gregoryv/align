// package align implements NeedlemanWunsch pairwise alignment
package align

type comparable interface {
	equal(ai, bj int) bool
}

// NeedlemanWunsch aligns sequences a and b with simple scores 1, -1 and -1
func NeedlemanWunsch(a, b sequence) *Result {
	return NeedlemanWunschCustom(a, b, 1, -1, -1)
}

// NeedlemanWunschCustom calculates the score matrix using custom scores match,
// missmatch and insert/delete.
func NeedlemanWunschCustom(a, b sequence, match, miss, indel int) *Result {
	F := newMatrix(len(a)+1, len(b)+1, miss)
	m := &Result{a: a, b: b, f: F}
	for y := 1; y < len(F); y++ {
		for x := 1; x < len(F[0]); x++ {
			diag := F[y-1][x-1].max + miss
			if m.equal(x-1, y-1) {
				diag = F[y-1][x-1].max + match
			}

			c := newCell(
				F[y-1][x].max+indel,
				F[y][x-1].max+indel,
				diag,
			)
			F[y][x] = c
		}
	}
	return m

}

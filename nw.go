// package align implements NeedlemanWunsch pairwise alignment
package align

import (
	position "github.com/gregoryv/matrix"
)

type comparable interface {
	equal(ai, bj int) bool
}

// NeedlemanWunsch aligns sequences a and b with simple scores 1, -1, -1 and -1.
// Start gap is no different from extending the gap.
func NeedlemanWunsch(a, b Sequence) *Result {
	return NeedlemanWunschCustom(a, b, 1, -1, -1, -1)
}

// NeedlemanWunschCustom calculates the score matrix using custom scores match,
// missmatch, insert/delete and extended gap. Sequences a and b must not be empty.
func NeedlemanWunschCustom(a, b Sequence, match, miss, indel, ext int) *Result {
	F := newMatrix(len(a)+1, len(b)+1, miss)
	m := &Result{a: a, b: b, f: F}
	x, y := 1, 1
	boundary := position.Rect{x, y, len(F[0]) - 1, len(F) - 1}
	nav := position.NewXYNavigator(x, y, boundary)
	for more := true; more; x, y, more = nav.Right() {
		diag := F[y-1][x-1].max + miss
		if m.equal(x-1, y-1) {
			diag = F[y-1][x-1].max + match
		}
		top := F[y-1][x]
		left := F[y][x-1]
		gap := indel
		switch {
		case hasBit(top.origin, topOrigin):
			gap = ext
		case hasBit(left.origin, leftOrigin):
			gap = ext
		}
		c := newCell(
			top.max+gap,
			left.max+gap,
			diag,
		)
		F[y][x] = c
	}
	return m
}

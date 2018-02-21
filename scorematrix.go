package align

// cell is part of the matrix containing score calculations
type cell struct {
	top, left, diag, max, origin int
}

const (
	diagOrigin = iota
	leftOrigin
	topOrigin
)

func newCell(top, left, diag int) cell {
	// Note, using top as max
	c := cell{top, left, diag, top, 0}
	if left > c.max {
		c.max = left
	}
	if diag > c.max {
		c.max = diag
	}
	// Set origin bits
	if top == c.max {
		c.origin |= (1 << topOrigin)
	}
	if left == c.max {
		c.origin |= (1 << leftOrigin)
	}
	if diag == c.max {
		c.origin |= (1 << diagOrigin)
	}
	return c
}

type scoreMatrix [][]cell

// newMatrix returns a matrix with initialized first row and column using miss score.
// lenA is columns, and lenB rows.
func newScoreMatrix(lenA, lenB, miss int) scoreMatrix {
	g := make(scoreMatrix, lenB)
	for y := 0; y < lenB; y++ {
		g[y] = make([]cell, lenA)
	}
	// Fill first row
	for x := 1; x < lenA; x++ {
		g[0][x] = cell{max: x * miss, origin: 0}
	}
	// Fill first column
	for y := 1; y < lenB; y++ {
		g[y][0] = cell{max: y * miss, origin: 0}
	}
	return g
}

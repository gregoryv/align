package align

import (
	"fmt"
	"io"
)

const (
	gap = '-'
)

type Sequence []rune

func (s Sequence) reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type alignment struct {
	a Sequence
	b Sequence
}

func (pair alignment) String() string {
	return fmt.Sprintf("%s\n%s", string(pair.a), string(pair.b))
}

func (pair alignment) copyAdd(A, B rune) alignment {
	copy := alignment{}
	for i, _ := range pair.a {
		copy.a = append(copy.a, pair.a[i])
		copy.b = append(copy.b, pair.b[i])
	}
	copy.a = append(copy.a, A)
	copy.b = append(copy.b, B)
	return copy
}

type Result struct {
	a Sequence
	b Sequence
	f scoreMatrix
}

func (m *Result) equal(i, j int) bool {
	return m.a[i] == m.b[j]
}

func (result *Result) findAlignments(x, y int, pair alignment) (all []alignment) {
	cell := result.f[y][x]
	if cell.origin == 0 {
		pair.a.reverse()
		pair.b.reverse()
		return append(all, pair)
	}
	A := result.a[x-1]
	B := result.b[y-1]
	if hasBit(cell.origin, diagOrigin) {
		copy := pair.copyAdd(A, B)
		all = append(all, result.findAlignments(x-1, y-1, copy)...)
	}
	if hasBit(cell.origin, leftOrigin) {
		copy := pair.copyAdd(A, gap)
		all = append(all, result.findAlignments(x-1, y, copy)...)
	}
	if hasBit(cell.origin, topOrigin) {
		copy := pair.copyAdd(gap, B)
		all = append(all, result.findAlignments(x, y-1, copy)...)
	}
	return all
}

// Score returns the score, ie. the max of the bottom right cell
func (result *Result) MaxScore() (score int) {
	x := len(result.f[0]) - 1
	y := len(result.f) - 1
	return result.f[y][x].max
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func (result *Result) PrintAlignment(w io.Writer) {
	for _, p := range result.Alignments() {
		fmt.Fprintf(w, "%s\n\n", p.String())
	}
}

// Alignments returns all possible alignments with max score
func (result *Result) Alignments() []alignment {
	x := len(result.f[0]) - 1
	y := len(result.f) - 1
	var pair alignment
	return result.findAlignments(x, y, pair)
}

type valueReader func(c cell) int

// PrintScoreMatrix writes out the sequences and their calculated score as shown on wikipedia.
func (result *Result) PrintScoreMatrix(w io.Writer) {
	result.printMatrix(func(c cell) int { return c.max }, w)
}

// PrintOrigins writes out the score origins for each score.
// 0 = None(top row and first column),  1 = Diagonal, 2 = Left and 4 = Top
// Eg. 5 = diagonal or top
func (result *Result) PrintOrigins(w io.Writer) {
	result.printMatrix(func(c cell) int { return c.origin }, w)
}

func (result *Result) printMatrix(read valueReader, w io.Writer) {
	// sequence a
	fmt.Fprintf(w, "%4s", "")
	for _, r := range result.a {
		fmt.Fprintf(w, "%3v", string(r))
	}
	fmt.Fprint(w, "\n ")
	// First row
	for x := 0; x < len(result.f[0]); x++ {
		fmt.Fprintf(w, "%3v", read(result.f[0][x]))
	}
	fmt.Fprint(w, "\n")
	for y, row := range result.f[1:] {
		fmt.Fprintf(w, "%s", string(result.b[y])) // sequence b
		for x := 0; x < len(row); x++ {
			fmt.Fprintf(w, "%3v", read(row[x]))
		}
		fmt.Fprint(w, "\n")
	}
}

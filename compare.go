package align

// Compare returns 1.0 if equal otherwise the value of matching
// characters.
func Compare(a, b string) float32 {
	result := NeedlemanWunschCustom(
		[]rune(a),
		[]rune(b),
		1,
		0,
		0,
		0,
	)
	return float32(result.MaxScore()) / float32(min(len(a), len(b)))
}

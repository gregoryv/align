package align

type Navigator struct {
	x, y, xi, yi, xj, yj int
}

func NewNavigator(x, y, xi, yi, xj, yj int) *Navigator {
	return &Navigator{x, y, xi, yi, xj, yj}
}

// Right returns next position to the right, if the end of a row is reached
// following position starting at xi is returned.
func (iter *Navigator) Right() (x, y int, more bool) {
	x = iter.x
	y = iter.y
	if y > iter.yj {
		return x, y, false
	}
	iter.x++
	if iter.x > iter.xj {
		// next row
		iter.y++
		iter.x = iter.xi
	}
	return x, y, true
}

func (iter *Navigator) Up() (x, y int, more bool) {
	x = iter.x
	y = iter.y
	if x < iter.xi {
		return x, y, false
	}
	iter.y--
	if iter.y < iter.yi {
		// previous column
		iter.x--
		iter.y = iter.yj
	}
	return x, y, true
}

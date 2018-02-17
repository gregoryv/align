package align

type Navigator struct {
	x, y, xi, yi, xj, yj int
}

func NewNavigator(x, y, xi, yi, xj, yj int) *Navigator {
	return &Navigator{x, y, xi, yi, xj, yj}
}

// Right returns next position to the right, if the end of a row is reached
// following position starting at xi is returned.
func (nav *Navigator) Right() (x, y int, more bool) {
	nav.x++
	if nav.x > nav.xj {
		// next row
		nav.y++
		nav.x = nav.xi
	}
	if nav.y > nav.yj {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

func (nav *Navigator) Up() (x, y int, more bool) {
	nav.y--
	if nav.y < nav.yi {
		// previous column
		nav.x--
		nav.y = nav.yj
	}
	if nav.x < nav.xi {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

func (nav *Navigator) Left() (x, y int, more bool) {
	nav.x--
	if nav.x < nav.xi {
		// next row
		nav.y--
		nav.x = nav.xj
	}
	if nav.y < nav.yi {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

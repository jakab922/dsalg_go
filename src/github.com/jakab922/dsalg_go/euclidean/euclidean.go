package dsalg

func Euclidean(a, b int) int {
	if a > b {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func ExtendedEuclidean(a, b int) (int, int, int) {
	swapped := false
	if b > a {
		a, b = b, a
		swapped = true
	}
	if b == 0 { // If the not-bigger number is 0 then we return the other number.
		if swapped {
			return a, 0, 1
		} else {
			return a, 1, 0
		}
	} else if a == b {
		return a, 1, 0
	}
	aCurr, bCurr := 0, 1
	aPrev, bPrev := 1, 0

	for b > 0 {
		t := a / b
		aPrev, aCurr = aCurr, aPrev-t*aCurr
		bPrev, bCurr = bCurr, bPrev-t*bCurr
		a, b = b, a%b
	}

	if swapped {
		return a, bPrev, aPrev
	} else {
		return a, aPrev, bPrev
	}
}

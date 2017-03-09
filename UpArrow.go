package lib

import "math"

/*UpArrow ...*/
func UpArrow(base uint64, exponant uint64, upArrowAmount uint64) uint64 {
	if upArrowAmount <= 0 {
		return uint64(base * exponant)
	} else if upArrowAmount == 1 {
		return uint64(math.Pow(float64(base), float64(exponant)))
	} else if exponant <= 1 {
		return uint64(base)
	}
	return UpArrow(base, UpArrow(base, exponant-1, upArrowAmount), upArrowAmount-1)
}

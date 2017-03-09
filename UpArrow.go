package lib

import "math"

/*UpArrow ...*/
func UpArrow(base int64, exponant int64, upArrowAmount int64) int64 {
	if upArrowAmount <= 0 {
		return int64(base * exponant)
	} else if upArrowAmount == 1 {
		return int64(math.Pow(float64(base), float64(exponant)))
	} else if exponant == 1 {
		return int64(base)
	}
	return UpArrow(base, UpArrow(base, exponant-1, upArrowAmount), upArrowAmount-1)
}

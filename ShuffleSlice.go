package lib

/*ShuffleIntSlice ...*/
func ShuffleIntSlice(slice []int) {
	for len(slice) > 0 {
		n := len(slice)
		randIndex := seededRand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

/*ShuffleStringSlice ...*/
func ShuffleStringSlice(slice []string) {
	for len(slice) > 0 {
		n := len(slice)
		randIndex := seededRand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

/*ShuffleByteSlice ...*/
func ShuffleByteSlice(slice []byte) {
	for len(slice) > 0 {
		n := len(slice)
		randIndex := seededRand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

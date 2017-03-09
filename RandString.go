package lib

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

/*RandStringWithCharset forms a string with the random seed and the charset*/
func RandStringWithCharset(lenght int, charset string) string {
	b := make([]byte, lenght)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/*RandString returns functionCall StringWithCharset*/
func RandString(lenght int) string {
	return RandStringWithCharset(lenght, charset)
}

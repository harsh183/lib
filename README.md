# lib

## func RandStringWithCharset(lenght int, charset string) string
creates a random string with lenght and a charset (e.g. "abcdefghijklmnopqrstuvwxyz0123456789")

## func RandString(lenght int) string
creates a random string with lenght and the charset ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

## func ShuffleIntSlice(slice []int)
shuffles all entrys of an []int

## func ShuffleStringSlice(slice []string)
shuffles all entrys of an []string

## func ShuffleByteSlice(slice []byte)
shuffles all entrys of an []byte

## func UpArrow(base uint64, exponant uint64, upArrowAmount uint64) uint64 {
executes the upArrow function (https://en.wikipedia.org/wiki/Knuth%27s_up-arrow_notation)

# lib

## func RandStringWithCharset(lenght int, charset string) string
creates a random string with length and a charset (e.g. "abcdefghijklmnopqrstuvwxyz0123456789")

## func RandString(lenght int) string
creates a random string with length and the charset ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

## func ShuffleIntSlice(slice []int)
shuffles all entries of an `[]int`

## func ShuffleStringSlice(slice []string)
shuffles all entries of an `[]string`

## func ShuffleByteSlice(slice []byte)
shuffles all entries of an `[]byte`

## func UpArrow(base uint64, exponant uint64, upArrowAmount uint64) uint64
computes the upArrow function (https://en.wikipedia.org/wiki/Knuth%27s_up-arrow_notation)

## func GetRequestContentFromRequest(req *http.Request) map[string]interface{}
returns the JSON from an `http.Request`

## func GetRequestContentFromResponse(resp *http.Response) map[string]interface{}
returns the JSON from an `http.Response`

## func IntRange(args ...int) <-chan int
takes 1 to 3 arguments
#### for one argument:
* `IntRange(10)` 
returns :
* `0 1 2 3 4 5 6 7 8 9`

#### for two arguments:
* `IntRange(10, 20)` 
returns :
* `10 11 12 13 14 15 16 17 18 19`

#### for three arguments:
* `IntRange(20, 0, -1)` 
returns :
* `20 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1`

## func SiveOfEratosthenes(numberOfPrimes int64) []int64
returns an slice of prime numbers smaller than the parameter `numberOfPrimes`

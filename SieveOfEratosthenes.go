package lib

/*SieveOfEratosthenes returns an Slice of primenumbers smaller than the given paramter*/
func SieveOfEratosthenes(numberOfPrimes int64) []int64{
	ch := make(chan int64)
	var returnValue []int64
	go generate(ch)
	var a int64 = 0
	for i := a; i < numberOfPrimes; i++ {
		prime := <-ch
		if prime <= numberOfPrimes {
			returnValue = append(returnValue, prime)
			ch1 := make(chan int64)
			go filter(ch, ch1, prime)
			ch = ch1
		} else {
			break
		}
	}
	return returnValue
}

func IsPrime(number int64) bool{
	if number <=1 {
		return false
	}
	var i int64 = 2
	for i*i <= number {
		if IsPrime(i){
			if number%i==0 {
				return false
			}
		}
		i++
	}
	return true
}

func filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generate(ch chan<- int64) {
	var a int64 = 2
	for i := a; ; i++ {
		ch <- i
	}
}

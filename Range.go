package lib

import "fmt"

func IntRange(args ...int) <-chan int {
	var start, stop, step int
	switch len(args) {
	case 1: // 1 argument: stop
		start = 0
		stop = args[0]
		step = 1
	case 2: // 2 arguments: start, stop
		start = args[0]
		stop = args[1]
		step = 1
	case 3: // 3 arguments: start, stop, step
		start = args[0]
		stop = args[1]
		step = args[2]
	default: // invalid argument count
		fmt.Println("IntRange takes 1 to 3 arguments.")
	}
	ch := make(chan int)
	if step >= 0 {
		// increment case
		go func() {
			for i := start; i < stop; i += step {
				ch <- i
			}
			close(ch)
		}()
	} else {
		// decrement case
		go func() {
			for i := start; i > stop; i += step {
				ch <- i
			}
			close(ch)
		}()
	}
	return ch
}

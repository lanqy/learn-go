package main
import (
	"fmt"
	"sync"
)

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		var sum int = 0
// 		for i := 0; i < 10; i++ {
// 			sum += i
// 		}
// 		ch <- sum
// 	}()

// 	fmt.Println(<-ch)
// }

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Println("Received %d ", i)
		fmt.Println()
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 1000000 integers on channel.
	for i := 1; i <= 1000000; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
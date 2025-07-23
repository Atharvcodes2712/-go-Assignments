package main

import (
	"fmt"
	"sync"
)

func squareWorker(num float64, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	square := num * num
	ch <- square
}
func aggregateSquares(ch chan float64, ok chan bool) {
	var sum float64
	for square := range ch {
		sum += square
	}
	fmt.Printf("Aggregate of squares is:%.2f", sum)
	ok <- true
}
func main() {
	//accepting list of numbers and storing in a slice
	var noOfDigits uint
	fmt.Println("Enter number of digits to perform squaring and aggregation")
	fmt.Scan(&noOfDigits)
	numbers := make([]float64, noOfDigits)
	fmt.Printf("Enter %d numbers\n", noOfDigits)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	//created channel ch of type float64 to square floating numbers as well
	ch := make(chan float64, noOfDigits)
	ok := make(chan bool)
	// initialized sync.waitgroup variable
	var wg sync.WaitGroup
	go aggregateSquares(ch, ok)
	//calling squareWorker on each number in slice
	for _, num := range numbers {
		//will wait until it becomes 0
		wg.Add(1)
		go squareWorker(num, ch, &wg)
	}
	//to keep main waiting
	wg.Wait()
	//closing channel
	close(ch)
	<-ok
}

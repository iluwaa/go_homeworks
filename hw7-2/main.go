package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numbers := make(chan int)
	result := make(chan string)
	from := 150
	to := 5000

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go GenerateNumbers(from, to, result, numbers, wg)
	go FindMinMax(from, to, numbers, result)

	wg.Wait()

}

func GenerateNumbers(from int, to int, in chan string, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < rand.Intn(15-3)+3; i++ {
		out <- rand.Intn(to-from) + from
	}
	close(out)

	fmt.Println(<-in)
}

func FindMinMax(from int, to int, in chan int, out chan string) {
	min := to
	max := from

	for number := range in {
		fmt.Printf("%d ", number)
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}

	out <- fmt.Sprintf("\nMin : %d, Max: %d", min, max)
}

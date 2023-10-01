package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Яка створює 3 горутини. Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину.
Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал.
Третя горутина виводить середнє значення на екран.
*/

type Numbers struct {
	numbers   []int
	channelId int
}

func main() {
	wg := &sync.WaitGroup{}
	channels := CreateChannels(3)

	for i, ch := range channels {
		go GenerateNumbers(ch, i, wg)
	}

	OutoutChannel(channels)

	wg.Wait()

}

func CreateChannels(count int) []chan Numbers {
	c := make([]chan Numbers, count)
	for k := 0; k < count; k++ {
		c[k] = make(chan Numbers)
	}
	return c

}

func GenerateNumbers(c chan Numbers, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	slicesCount := rand.Intn(15-3) + 3

	fmt.Printf("In channel %d will be generated %d slices\n", id, slicesCount)

	for i := 0; i < slicesCount; i++ {
		numbers := Numbers{make([]int, 0), id}

		for k := 0; k < rand.Intn(10-2)+2; k++ {
			numbers.numbers = append(numbers.numbers, rand.Intn(100))
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
		c <- numbers
	}
	close(c)
	return
}

func OutoutChannel(channels []chan Numbers) {
	output := make(chan Numbers)

	go func() {
		for _, ch := range channels {
			go func(ch chan Numbers) {
				for val := range ch {
					output <- val
				}
			}(ch)
		}
	}()

	go func() {
		for val := range output {
			sum := 0

			for _, i := range val.numbers {
				sum += i
			}

			fmt.Printf("Recieved %v from channel %d, avg = %d\n", val.numbers, val.channelId, sum/len(val.numbers))
		}
	}()

}

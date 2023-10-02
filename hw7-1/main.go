package main

import (
	"fmt"
	"math/rand"
)

/*
Яка створює 3 горутини. Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину.
Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал.
Третя горутина виводить середнє значення на екран.
*/

func main() {
	numbers := make(chan int)
	output := make(chan string)

	go GenerateNumbers(numbers)

	go FindAvarage(numbers, output)

	fmt.Println(<-output)

}

func FindAvarage(in chan int, out chan string) {
	count := 0
	var sum float32
	result := "Generated "

	for number := range in {
		sum += float32(number)
		count++
		result += fmt.Sprintf("%d ", number)
	}

	result += fmt.Sprintf(", sum is %v, count is %d, avarage is %v", sum, count, sum/float32(count))
	out <- result

}

func GenerateNumbers(c chan int) {
	for i := 0; i < rand.Intn(20-2)+2; i++ {
		c <- rand.Intn(100-1) + 1
	}
	close(c)
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	products := ProductsGenerator(ctx)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	CalculateCost(products, ctx, wg)

	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()

}

type Product struct {
	Name  string
	Cost  int
	Count int
}

func ProductsGenerator(ctx context.Context) <-chan Product {
	c := make(chan Product)

	go func() {
		for i := 0; ; {
			select {
			case c <- Product{
				Name:  "Product" + strconv.Itoa(rand.Intn(100)),
				Cost:  rand.Intn(100-1) + 1,
				Count: rand.Intn(10-1) + 1,
			}:
				i += 1
			case <-ctx.Done():
				fmt.Printf("Generated %d products\n", i)
				return

			}
		}
	}()

	return c
}

func CalculateCost(products <-chan Product, ctx context.Context, wg *sync.WaitGroup) {

	go func() {

		for i := 0; ; {
			select {
			case prudct := <-products:
				fmt.Printf("Recieved product %s with cost %d and count %d. Total cost is %d\n", prudct.Name, prudct.Cost, prudct.Count, prudct.Cost*prudct.Count)
				i += 1
				time.Sleep(1 * time.Second)
			case <-ctx.Done():
				fmt.Printf("End of work calculator, calculated %d products. \n", i)
				wg.Done()
				return
			}

		}
	}()
}

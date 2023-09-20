package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/lzap/deagon"
)

const subjects = "Animation,App development,Audio production,Computer programming,Computer repair,Film production,Graphic design,Media technology,Music production,Typing,Video game development,Web design,Web programming,Word processing"

func main() {
	scoresMap := generateScoresMap()
	printReport(scoresMap)
}

func printReport(scoresMap map[string][]float32) {
	fmt.Printf("%s scores report for %d quarter 2023: \n", deagon.RandomName(deagon.NewCapitalizedSpaceFormatter()), rand.Intn(4)+1)
	for key, value := range scoresMap {
		var sum float32
		fmt.Printf("- %s: ", key)
		for i, v := range value {
			sum = sum + v
			if i+1 != len(value) {
				fmt.Printf("%v, ", v)
			} else {
				fmt.Printf("%v. avg score: %v.", v, sum/float32(len(value)))
			}

		}
		fmt.Println()
	}
}

func generateScoresMap() map[string][]float32 {
	// Idk why we need float here
	var scoresMap = make(map[string][]float32)

	for _, subject := range strings.Split(subjects, ",") {
		for i := 0; i < rand.Intn(20)+1; i++ {
			scoresMap[subject] = append(scoresMap[subject], float32(rand.Intn(5)+1))
		}
	}

	return scoresMap
}

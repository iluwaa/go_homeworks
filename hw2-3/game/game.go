package game

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"structs/zoo"
)

const answersCount = 4

func PlayGame() {
	var choiseInt int
	var zooSize int
	var Zoo *zoo.Zoo

	fmt.Println("Hello, we have a different size zoos:")
	zooSizes := zooSizes()
	for i := 1; i <= len(zooSizes); i++ {
		fmt.Printf("%d) %s\n", i, zooSizes[i])
	}
	fmt.Printf("Which size you prefer? Type a number of your choise: ")
	fmt.Scan(&choiseInt)
	switch choiseInt {
	case 1:
		zooSize = 4
	case 2:
		zooSize = 8
	case 3:
		zooSize = 12
	}

	Zoo = zoo.MakeZoo(zooSize)
	questionsList := makeQuestions(Zoo, zooSize)

	for _, v := range questionsList.questions {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Printf("You choose a %s zoo, here we have a %d cages:\n", zooSizes[choiseInt], len(Zoo.Cages))
	for _, cage := range Zoo.Cages {
		fmt.Printf("In %d cage we have a %d animals. Now I will tell you about their taxonomy.\n", cage.Number, len(cage.Animals))
		for _, animal := range cage.Animals {
			fmt.Printf("%s:\n", animal.Name)
			for i := 0; i < reflect.TypeOf(animal.Taxonomy).NumField(); i++ {
				taxKey := reflect.Indirect(reflect.ValueOf(animal.Taxonomy)).Type().Field(i).Name
				taxValue := reflect.ValueOf(animal.Taxonomy).Field(i).String()
				fmt.Printf("  - %s: %s\n", taxKey, taxValue)
			}
			fmt.Printf("-----------------------------------\n")
		}
	}

RepeatGame:
	questionsList = makeQuestions(Zoo, zooSize)

	for i, question := range questionsList.questions {
		fmt.Printf("\n(%d/%d) %s\n", i+1, zooSize, question.description)
		for i, answer := range question.answers {
			fmt.Printf("%d) %s ", i+1, answer.value)
		}
		fmt.Printf("\nType a number of correct answer: ")
		fmt.Scan(&choiseInt)
		if question.answers[choiseInt-1].isCorrect {
			fmt.Println("You r right! Lets go to the next question")
		} else {
			fmt.Printf("Wrong answer, you r lost :( ")
			showCoreect(question.answers)
			askToRepeat()
			goto RepeatGame
		}
	}
	fmt.Println("Congrats, you are won!")
	askToRepeat()
}

func showCoreect(answers []answer) {
	for _, answer := range answers {
		if answer.isCorrect {
			fmt.Printf("Correct answer was: %s\n", answer.value)
		}
	}
}

func askToRepeat() {
	var choiseStr string
	fmt.Printf("Wanna play again? ")
	fmt.Scan(&choiseStr)
	if !strings.Contains("yes", choiseStr) {
		fmt.Println("Bye!")
		os.Exit(0)
	}
}

func makeQuestions(Zoo *zoo.Zoo, zooSize int) *questions {
	result := new(questions)
	for i := 0; i < zooSize; i++ {
		result.questions = append(result.questions, makeQuestion(Zoo))

	}
	return result
}

func makeAnswer(value string) answer {
	var result answer
	result.value = value

	return result
}

func makeQuestion(Zoo *zoo.Zoo) *question {
	result := new(question)

	randomAnimals := zoo.RandomAnimalsList(Zoo, answersCount)

	correctAnswerIndex := rand.Intn(answersCount)
	for i, animal := range randomAnimals {
		result.answers = append(result.answers, makeAnswer(animal.Name))
		if i == correctAnswerIndex {
			result.answers[correctAnswerIndex].markCorrect()
			taxonomy := Zoo.Cages[animal.CageIndex].Animals[animal.AnimalIndex].Taxonomy
			fieldsCount := reflect.TypeOf(taxonomy).NumField()
			randomTax := rand.Intn(fieldsCount)

			taxValue := reflect.ValueOf(taxonomy).Field(randomTax).String()
			taxKey := reflect.Indirect(reflect.ValueOf(taxonomy)).Type().Field(randomTax).Name
			result.description = fmt.Sprintf("Which animal %s is %s?", taxKey, taxValue)
		}
	}
	return result
}

func zooSizes() map[int]string {
	return map[int]string{
		1: "Small",
		2: "Medium",
		3: "Large",
	}
}

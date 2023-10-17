package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	quiz "quiz/pkg"
	"sync"
	"syscall"
	"time"
)

const category = "Linux"

var playersCount int
var questionsCount int

func main() {
	flag.IntVar(&playersCount, "players", 2, "Number of players.")
	flag.IntVar(&questionsCount, "questions", 5, "Number of questions.")
	flag.Parse()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	players := quiz.CreatePlayers(playersCount)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	questions := Generator(questionsCount, interupt, wg)

	answers := listenAnswers()

	go func() {
		for question := range questions {
			PlayRound(question, players, answers)
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-interupt

		fmt.Println()
		quiz.CheckWinner(players)
		os.Exit(0)

	}(wg)

	wg.Wait()

}

func Generator(questionsCount int, interupt chan os.Signal, wg *sync.WaitGroup) <-chan quiz.Question {

	c := make(chan quiz.Question)

	go func() {
		defer wg.Done()

		for i := 0; i < questionsCount; i++ {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			question := quiz.GetQuestion(category)
			question.Context = ctx

			select {
			case c <- question:
				question.PrintQuestion()
				time.Sleep(3 * time.Second)
				fmt.Printf("Correct answer is: %s\n", question.CorrectAnswer)
				cancel()
			}
		}
		interupt <- syscall.SIGINT
	}()

	return c
}

func PlayRound(question quiz.Question, players []*quiz.Player, answers chan<- quiz.Answer) {

	for _, player := range players {
		go func(player *quiz.Player, question quiz.Question) {
			latency := time.After(time.Duration(rand.Intn(4)) * time.Second)
			for {
				select {
				case <-latency:
					answers <- player.MakeAnswer(&question)
					return
				case <-question.Context.Done():
					fmt.Println("reached timeout, player", player.Id)
					return
					// default:
					// 	time.Sleep(1 * time.Second)
				}
			}
		}(player, question)
	}
}

func listenAnswers() chan<- quiz.Answer {
	answers := make(chan quiz.Answer)

	go func() {
		for {
			select {
			case answer := <-answers:
				fmt.Printf("Recieved answer from Player %d: %s. ", answer.Player.Id, answer.Answer)
				isCorrect := answer.Question.IsCorrect(answer.Answer)
				fmt.Println(answer.Question.IsCorrect(answer.Answer))
				if isCorrect {
					answer.Player.AddPoint()
				}
			default:
				time.Sleep(1 * time.Second)
			}

		}
	}()

	return answers
}

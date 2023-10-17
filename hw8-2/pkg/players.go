package quiz

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Id     int
	Points int
}

type Answer struct {
	Player   *Player
	Question *Question
	Answer   string
}

func CreatePlayers(count int) []*Player {
	players := make([]*Player, 0)
	for i := 0; i < count; i++ {
		players = append(players, &Player{
			Id: i + 1,
		})
	}
	return players
}

func (player *Player) MakeAnswer(question *Question) Answer {
	var answer Answer
	availableAnswersAnswers := question.AvailableAnswers()

	answer.Answer = availableAnswersAnswers[rand.Intn(len(availableAnswersAnswers))]
	answer.Player = player
	answer.Question = question

	return answer
}

func (player *Player) AddPoint() {
	player.Points += 1
}

func CheckWinner(players []*Player) {
	maxPoints := -1
	isDraw := false
	var winner *Player

	fmt.Println("Results:")
	for _, player := range players {
		if maxPoints < player.Points && player.Points != 0 {
			maxPoints = player.Points
			winner = player
		} else if maxPoints == player.Points {
			isDraw = true
		}
		fmt.Printf("Player %d has %d pts. \n", player.Id, player.Points)
	}

	if !isDraw {
		fmt.Printf("Player %d wins with %d pts!\n", winner.Id, winner.Points)
	} else {
		fmt.Println("Draw!")
	}
}

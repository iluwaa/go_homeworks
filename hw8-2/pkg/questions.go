package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Question struct {
	Context     context.Context
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Description any    `json:"description"`
	Answers     struct {
		AnswerA string `json:"answer_a"`
		AnswerB string `json:"answer_b"`
		AnswerC string `json:"answer_c"`
		AnswerD string `json:"answer_d"`
		AnswerE any    `json:"answer_e"`
		AnswerF any    `json:"answer_f"`
	} `json:"answers"`
	MultipleCorrectAnswers string `json:"multiple_correct_answers"`
	CorrectAnswers         struct {
		AnswerACorrect string `json:"answer_a_correct"`
		AnswerBCorrect string `json:"answer_b_correct"`
		AnswerCCorrect string `json:"answer_c_correct"`
		AnswerDCorrect string `json:"answer_d_correct"`
		AnswerECorrect string `json:"answer_e_correct"`
		AnswerFCorrect string `json:"answer_f_correct"`
	} `json:"correct_answers"`
	CorrectAnswer string `json:"correct_answer"`
	Explanation   any    `json:"explanation"`
	Tip           any    `json:"tip"`
	Tags          []struct {
		Name string `json:"name"`
	} `json:"tags"`
	Category   string `json:"category"`
	Difficulty string `json:"difficulty"`
}

func (question *Question) IsCorrect(answer string) bool {

	if question.CorrectAnswer == answer {
		return true
	} else {
		return false
	}
}

func (question *Question) PrintQuestion() {
	fmt.Println("--------------------------")
	fmt.Printf("Question: %s\n", question.Question)
	fmt.Println("Answers: ")
	fmt.Printf("1) %s\n", question.Answers.AnswerA)
	fmt.Printf("2) %s\n", question.Answers.AnswerB)
	fmt.Printf("3) %s\n", question.Answers.AnswerC)
	fmt.Printf("4) %s\n", question.Answers.AnswerD)
	if question.Answers.AnswerE != nil {
		fmt.Printf("5) %s\n", question.Answers.AnswerE)
	}
	if question.Answers.AnswerF != nil {
		fmt.Printf("6) %s\n", question.Answers.AnswerF)
	}
	fmt.Println("--------------------------")
}

func (question *Question) AvailableAnswers() map[int]string {
	answersMapping := map[int]string{
		0: "answer_a",
		1: "answer_b",
		2: "answer_c",
		3: "answer_d",
	}
	if question.Answers.AnswerE != nil {
		answersMapping[4] = "answer_e"
	}
	if question.Answers.AnswerF != nil {
		answersMapping[5] = "answer_f"
	}
	return answersMapping
}

func GetQuestion(category string) Question {
	var result []Question

	baseUrl, err := url.Parse("https://quizapi.io/api/v1/questions")
	if err != nil {
		fmt.Println(err)

	}
	urlParams := url.Values{}
	urlParams.Add("category", category)
	urlParams.Add("limit", "1")
	baseUrl.RawQuery = urlParams.Encode()

	request, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	// I know that api key it is sensitive info, but do not matter
	request.Header.Add("X-Api-Key", "MFR7csqeHgNC4ttVArZENAo7t1YaWzT5CiPn7LOa")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		fmt.Println(err)
	}

	return result[0]
}

// Debug
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

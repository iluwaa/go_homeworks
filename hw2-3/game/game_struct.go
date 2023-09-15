package game

type questions struct {
	questions []*question
}

type question struct {
	description string
	answers     []answer
}

type answer struct {
	value     string
	isCorrect bool
}

func (answer *answer) markCorrect() {
	answer.isCorrect = true
}

// // Maybe next time ;)
// type Tip struct {
// 	Name        string
// 	used        bool
// 	Id          int
// 	Description string
// }

// func (tip *Tip) UseTip() {
// 	tip.used = true
// }

// func makeTips() map[int]*Tip {
// 	return map[int]*Tip{
// 		1: &Tip{
// 			Description: "Clear 2 wrong answers.",
// 			Name:        "Fifty-fifty",
// 		},
// 		2: &Tip{
// 			Description: "Make a call to your friend to ask the question.",
// 			Name:        "Call a friend",
// 		},
// 		3: &Tip{
// 			Description: "Ask for audiance from spectators.",
// 			Name:        "Ask the audiance",
// 		},
// 	}
// }

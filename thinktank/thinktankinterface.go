package thinktank

type ThinkTank interface {
	Converse(userInput string, sid int64) (string, error)
}

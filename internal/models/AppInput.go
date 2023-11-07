package models

type AppInput struct {
	Question string
	Value    string
	Input    Input
}

func new(input string) AppInput {
	return AppInput{Question: input}
}
func NewTextField(input string) AppInput {
	question := new(input)
	model := newShortAnswerField()
	question.Input = model
	return question
}

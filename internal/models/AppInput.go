package models

type AppInput struct {
	question string
	value    string
	input    Input
	index    uint
}

func new(input string) AppInput {
	return AppInput{question: input}
}
func NewTextField(input string) AppInput {
	question := new(input)
	model := newShortAnswerField()
	question.input = model
	return question
}

package models

import "strings"

var Bool_choices = []string{"Yup", "Nuh-uh"}

type Radio struct {
	Cursor int
	Choice string
}

func (m Radio) View() string {
	s := strings.Builder{}
	s.WriteString("Generate VScode configuration file?\n\n")

	for i := 0; i < len(Bool_choices); i++ {
		if m.Cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(Bool_choices[i])
		s.WriteString("\n")
	}
	return s.String()
}

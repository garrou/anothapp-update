package helpers

import (
	"fmt"
	"strings"
)

func MapToString(m map[string]string) string {
	s := ""

	for _, element := range m {
		s += element + ";"
	}
	return strings.TrimSuffix(s, ";")
}

func Format[T any](title string, models []T) string {
	s := fmt.Sprintf("%s\n\n", title)

	for _, model := range models {
		s += fmt.Sprintln(model)
	}
	return s
}

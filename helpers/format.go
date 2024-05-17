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

func Format[T any](title string, items []T) string {
	s := fmt.Sprintf("%s\n\n", title)

	for _, item := range items {
		s += fmt.Sprintln(item)
	}
	return fmt.Sprintln(s)
}

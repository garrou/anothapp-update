package helpers

import (
	"anothapp_update/models"
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

func FormatMessage(shows []models.Show, updated []models.Season, deleted []models.Season) string {
	message := fmt.Sprintf("%d shows updated\n", len(shows))

	for _, show := range shows {
		message += format(show)
	}
	message += fmt.Sprintf("%d seasons updated\n", len(updated))

	for _, season := range updated {
		message += format(season)
	}
	message += fmt.Sprintf("%d seasons deleted\n", len(deleted))

	for _, season := range deleted {
		message += format(season)
	}
	return message
}

func format(toFormat models.Stringer) string {
	return strings.Repeat(" ", 4) + fmt.Sprintln(toFormat)
}

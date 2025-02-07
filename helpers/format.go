package helpers

import (
	"anothapp_update/models"
	"fmt"
	"strings"
)

func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func MapToString(m map[string]string) string {
	s := ""

	for _, element := range m {
		s += element + ";"
	}
	return strings.TrimSuffix(s, ";")
}

func FormatMessage(updatedShows []models.Show, deletedShows []models.Show, updatedSeasons []models.Season, deletedSeasons []models.Season) string {
	message := fmt.Sprintf("%d shows updated\n", len(updatedShows))
	for _, show := range updatedShows {
		message += format(show)
	}

	message += fmt.Sprintf("%d shows deleted\n", len(deletedShows))
	for _, show := range deletedShows {
		message += format(show)
	}

	message += fmt.Sprintf("%d seasons updated\n", len(updatedSeasons))
	for _, season := range updatedSeasons {
		message += format(season)
	}

	message += fmt.Sprintf("%d seasons deleted\n", len(deletedSeasons))
	for _, season := range deletedSeasons {
		message += format(season)
	}
	return message
}

func format(toFormat models.Stringer) string {
	return strings.Repeat(" ", 4) + fmt.Sprintln(toFormat)
}

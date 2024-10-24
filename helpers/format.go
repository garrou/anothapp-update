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

func FormatMessage(updatedShows []models.Show, deletedShows []models.Show, updatedSeasons []models.Season, deletedSeasons []models.Season) string {
	return strings.Join([]string{
		createMessage("shows updated", updatedShows),
		createMessage("shows deleted", deletedShows),
		createMessage("seasons updated", updatedSeasons),
		createMessage("seasons deleted", deletedSeasons),
	}, "\n")
}

func createMessage[T models.Stringer](title string, toWrite []T) string {
	message := fmt.Sprintf("%d %s\n", len(toWrite), title)
	for _, show := range toWrite {
		message += format(show)
	}
	return message
}

func format(toFormat models.Stringer) string {
	return strings.Repeat(" ", 4) + fmt.Sprintln(toFormat)
}

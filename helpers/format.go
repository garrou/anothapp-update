package helpers

import (
	"anothapp_update/models"
	"fmt"
)

func MapToString(m map[string]string) string {

	s := ""

	for _, element := range m {
		s += element + ";"
	}
	return strings.TrimSuffix(s, ";")
}


func FormatMsg(seasonsToUp []models.Season, seasonsToDel []models.Season) string {
	return format("UPDATED", seasonsToUp) + format("DELETED", seasonsToDel)
}

func format(title string, seasons []models.Season) string {

	if len(seasons) == 0 {
		return ""
	}
	var msg = fmt.Sprintf("%s : ", title)

	for _, n := range seasons {
		msg += fmt.Sprintf("%d ", n.ShowId)
	}
	return fmt.Sprintln(msg)
}

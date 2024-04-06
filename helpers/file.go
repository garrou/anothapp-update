package helpers

import (
	"os"
)

func Write(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, wErr := f.WriteString(content)

	if wErr != nil {
		panic(wErr)
	}
}

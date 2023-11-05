package helpers

import (
	"bytes"
	"fmt"
	"os"
)

func SendTelegramMessage(message string) {
	token := os.Getenv("TOKEN")
	chatId := os.Getenv("CHAT_ID")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	content := fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, chatId, message)
	body := bytes.NewBuffer([]byte(content))
	HttpPost(url, body)
}

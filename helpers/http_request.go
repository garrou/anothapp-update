package helpers

import (
	"bytes"
	"io"
	"net/http"
)

func SendTelegramMessage(message string) {
	token := os.Getenv("TOKEN")
	chatId := os.Getenv("CHAT_ID")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	content := fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, chatId, message)
	body := bytes.NewBuffer([]byte(content))
	httpPost(url, body)
}

func HttpGet(url, key string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-BetaSeries-Key", key)
	resp, getErr := client.Do(req)

	if getErr != nil {
		panic(getErr)
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	return body
}

func httpPost(url string, body *bytes.Buffer) {
	client := &http.Client{}
	req, reqErr := http.NewRequest(http.MethodPost, url, body)

	if reqErr != nil {
		panic(reqErr)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	res, resErr := client.Do(req)

	if resErr != nil {
		panic(resErr)
	}
	defer res.Body.Close()

	if _, err := io.ReadAll(res.Body); err != nil {
		panic(err)
	}
}
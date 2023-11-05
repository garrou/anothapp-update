package helpers

import (
	"bytes"
	"io"
	"net/http"
)

func HttpGet(url string) []byte {
	resp, getErr := http.Get(url)

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

func HttpPost(url string, body *bytes.Buffer) {
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

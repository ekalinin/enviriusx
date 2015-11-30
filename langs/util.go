package langs

import (
	"io/ioutil"
	"net/http"
)

const (
	Version = "0.1.0"
)

var (
	UserAgent = "enviriusx/" + Version + " (https://github.com/ekalinin/enviriusx)"
)

func HttpGet(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", UserAgent)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

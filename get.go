// Fetch provides a super simple API for get and
// post-requesting REST API's, plus converting to/from JSON
package fetch

import (
	"io/ioutil"
	"net/http"

	json "github.com/bitly/go-simplejson"
)


func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return body, nil
}

func GetJson(url string) (*json.Json, error) {
	body, err := Get(url)
	if err != nil {
		return nil, err
	}

	data, err := json.NewJson(body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

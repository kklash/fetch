// Fetch provides a super simple API for get and
// post-requesting REST API's, plus converting to/from JSON
package fetch

import (
	"io/ioutil"
	"bytes"
	"net/http"

	json "github.com/bitly/go-simplejson"
)

const JSON_CONTENT_TYPE = "application/json"

func Post(url string, postData []byte) ([]byte, error) {
	contentType := http.DetectContentType(postData)
	resp, err := http.Post(url, contentType, bytes.NewReader(postData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PostJson(url string, postData *json.Json) (*json.Json, error) {
	postBytes, err := postData.Bytes()
	if err != nil {
		return nil, err
	}

	respBody, err := Post(url, postBytes)
	if err != nil {
		return nil, err
	}

	respData, err := json.NewJson(respBody)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

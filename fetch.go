// Fetch provides a super simple API for get and
// post-requesting REST API's, plus converting to/from JSON
package fetch

import (
	"io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
)

/*
Get pulls the contents of a URL synchronously, if passed
only one argument, or concurrently if given a response channel as a
second argument. Example:

response, err := fetch.Get("http://google.com") // synchronous

resp_chan := make(chan fetch.Response)
go fetch.Get("http://google.com", resp_chan)
...
response := <-resp_chan // concurrent

fetched responses can then be decoded into JSON arrays, or maps, of interfaces:

js, err := response.toArray()
...
firstItem := js[0].(string) // Type of the resultant maps/arrays must be asserted
*/
func Get(url string, chans ...chan Response) Response {
	result := new(Response)
	var (
		_chan chan Response
		async bool
	)
	if len(chans) > 0 {
		_chan = chans[0]
		async = true
	} else {
		async = false
	}
	resp, http_err := http.Get(url)
	if http_err != nil {
		result.Error = http_err
		result.Body = nil
		if async {
			_chan <- *result
		}
		return *result
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Error = err
		result.Body = nil
		if async {
			_chan <- *result
		}
		return *result
	}
	resp.Body.Close()
	result.Error = nil
	result.Body = bytes
	if async {
		_chan <- *result
	}
	return *result
}

/*
Post provides a simple one-stop command for posting any format of JSON data.
The key is in using interfaces for the values of a JSON object, and as the
contents of a JSON array.

postData must either be in the form of:

postData := map[string]interface{} {
	"ID": 49194,
	"name": "Jimmy",
	"Parents": []string {"James", "Beth"}
}

..in the case of a mapping (JSON object) or as a JSON array:
postData := []interface{} {
	"This", "is", "an", "array", 322, 999.999,
}
*/
func Post(url string, postData interface{}, chans ...chan Response) Response {
	result := new(Response)
	var (
		_chan chan Response
		async bool
	)
	if len(chans) > 0 {
		_chan = chans[0]
		async = true
	} else {
		async = false
	}
	var (
		postMapping map[string]interface{}
		postArray []interface{}
		jsData []byte
		err error
	)
	postMapping, ok := postData.(map[string]interface{})
	if ok {
		jsData, err = json.Marshal(postMapping)
	} else {
		postArray = postData.([]interface{})
		jsData, err = json.Marshal(postArray)
	}
	if err != nil {
		result.Error = err
		result.Body = nil
		if async {
			_chan <- *result
		}
		return *result
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsData))
	if err != nil {
		result.Error = err
		result.Body = nil
		if async {
			_chan <- *result
		}
		return *result
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Error = err
		result.Body = nil
		if async {
			_chan <- *result
		}
		return *result
	}
	resp.Body.Close()
	result.Error = nil
	result.Body = bytes
	if async {
		_chan <- *result
	}
	return *result
}

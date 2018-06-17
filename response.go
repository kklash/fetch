package fetch

import (
	"encoding/json"
)

/*
Response structures are returned by fetch.Get. They
can be used to directly parse the body of a get request
into JSON structures, or into a string. All methods
check for errors from the response.
*/
type Response struct {
	Error error
	Body  []byte
}

func (r *Response) ToArray() ([]interface{}, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	var arr []interface{}
	err := json.Unmarshal(r.Body, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *Response) ToMap() (map[string]interface{}, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	var mapping map[string]interface{}
	err := json.Unmarshal(r.Body, &mapping)
	if err != nil {
		return nil, err
	}
	return mapping, nil
}

func (r *Response) ToString() (string, error) {
	if r.Error != nil {
		return "", r.Error
	}
	return string(r.Body), nil
}

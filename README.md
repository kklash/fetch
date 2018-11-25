# Fetch for Go!


Inspired by the lightweight `node-fetch` package in Node.js, `fetch` is meant as a light utility for fetching and parsing web pages, especially for JSON based REST API's.

Its two primary functions are `GetJson` and `PostJson` which accept and return Json data using https://github.com/bitly/go-simplejson 

Example usage:

```
package main

import (
	"github.com/kklash/fetch"
	"fmt"
)

func main() {
	data, err := fetch.GetJson("http://ipinfo.io/json")
	if err != nil {
		panic(err)
	}

	ip, _ := data.Get("ip").String()

	city, _ := data.Get("city").String()

	fmt.Printf("IP: %s\nLocation: %s\n", ip, city)
// => IP: 93.115.86.8
//    Location: Bucharest
}

```

# Fetch for Go!


Inspired by the lightweight `node-fetch` package in Node.js, `fetch` is meant as a light utility for fetching and parsing web pages, especially for JSON based REST API's. Example usage:


```
package main

import (
	"github.com/kklash/fetch"
	"fmt"
)

func main() {
	resp := fetch.Get("http://ipinfo.io/json")
	if resp.Error != nil {
		panic(resp.Error)
	}
	json, _ := resp.ToMap()
	ip := json["ip"].(string)
	city := json["city"].(string)
	fmt.Printf("IP: %s\nLocation: %s\n", ip, city)
// => IP: 93.115.86.8
//    Location: Bucharest
}
```

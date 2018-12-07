package fetch

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Sets fetch up to use a custom proxy. For example,
// to use the tor browser's socks5 proxy:
// fetch.UseProxy("socks5://127.0.0.1:9150")
func UseProxy(proxyURL string) error {
	// Parse Tor proxy URL string to a URL type
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return fmt.Errorf("UseProxy: error parsing proxy URL: %s - %e", proxyURL, err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	transport := &http.Transport{Proxy: http.ProxyURL(parsedURL)}
	HttpClient = &http.Client{Transport: transport, Timeout: time.Second * 10}

	return nil
}

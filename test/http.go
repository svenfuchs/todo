package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func StubRequest(url string, status int, headers map[string]string, body string) {
	http.DefaultClient.Transport = &stubTransport{url, status, headers, body}
}

type stubTransport struct {
	url     string
	status  int
	headers map[string]string
	body    string
}

func (t *stubTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	t.checkUrl(request.URL.String())

	response := http.Response{
		Request:    request,
		StatusCode: t.status,
		Header:     http.Header{},
		Body:       ioutil.NopCloser(strings.NewReader(t.body)),
	}
	for key, value := range t.headers {
		response.Header.Set(key, value)
	}
	return &response, nil
}

func (t stubTransport) checkUrl(url string) {
	if t.url != url {
		log.Fatal(fmt.Sprintf("Unexpected request to %s (expected %s).", url, t.url))
	}
}

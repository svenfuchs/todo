package service

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
)

var Client = *http.DefaultClient

func NewHttpClient() HttpClient {
  return HttpClient {}
}

type HttpClient struct {}

func (c HttpClient) run(method string, uri string, query *map[string]string, headers *map[string]string, body *bytes.Buffer) []byte {
  request := NewHttpRequest(method, uri, query, headers, body)
  return request.run(&Client)
}

func NewHttpRequest(method string, uri string, query *map[string]string, headers *map[string]string, body *bytes.Buffer) HttpRequest {
  if headers == nil {
    headers = &map[string]string{}
  }
  if body == nil {
    body = bytes.NewBuffer([]byte{})
  }
  if query != nil {
    uri = appendQueryString(uri, *query)
  }
  return HttpRequest{ method, uri, headers, body }
}

type HttpRequest struct {
  method string
  uri string
  headers *map[string]string
  body *bytes.Buffer
}

func (r HttpRequest) run(client *http.Client) []byte {
  request, err := http.NewRequest(r.method, r.uri, r.body)
  if err != nil { log.Fatal(err) }
  r.setRequestHeaders(request)

  response, err := client.Do(request)
  defer response.Body.Close()
  if err != nil { log.Fatal(err) }

  r.checkResponseStatus(response)
  return r.readResponseBody(response)
}

func (r HttpRequest) setRequestHeaders(request *http.Request) {
  for key, value := range *r.headers {
    request.Header.Add(key, value)
  }
}

func (r HttpRequest) checkResponseStatus(response *http.Response) {
  if response.StatusCode / 100 != 2 {
    body := r.readResponseBody(response)
    log.Fatal(fmt.Sprintf("%d %s %s (%q)\n", response.StatusCode, r.method, r.uri, body))
  }
}

func (c HttpRequest) readResponseBody(response *http.Response) []byte {
  body, err := ioutil.ReadAll(response.Body)
  if err != nil { log.Fatal(err) }
  return body
}

func appendQueryString(uri string, data map[string]string) string {
  url, err := url.ParseRequestURI(uri)
  if err != nil { log.Fatal(err) }

  url.RawQuery = toQueryString(data)
  return fmt.Sprintf("%v", url)
}

func toQueryString(data map[string]string) string {
  query := url.Values{}
  for key, value := range data {
    query.Set(key, value)
  }
  return query.Encode()
}

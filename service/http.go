package service

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
)

type HttpRequest struct {
  method string
  uri string
  headers *map[string]string
  body *bytes.Buffer
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

func (h HttpRequest) Run() []byte {
  request := h.newRequest()
  response, err := http.DefaultClient.Do(request)
  defer response.Body.Close()
  if err != nil { log.Fatal(err) }

  h.checkStatus(response)
  return h.readBody(response)
}

func (h HttpRequest) newRequest() *http.Request {
  request, err := http.NewRequest(h.method, h.uri, h.body)
  if err != nil { log.Fatal(err) }

  h.setHeaders(request, h.headers)
  return request
}

func (h HttpRequest) setHeaders(request *http.Request, headers *map[string]string) {
  for key, value := range *headers {
    request.Header.Add(key, value)
  }
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

func (h HttpRequest) checkStatus(response *http.Response) {
  if response.StatusCode / 100 != 2 {
    body := h.readBody(response)
    log.Fatal(fmt.Sprintf("%d %s %s (%q)\n", response.StatusCode, h.method, h.uri, body))
  }
}

func (h HttpRequest) readBody(response *http.Response) []byte {
  body, err := ioutil.ReadAll(response.Body)
  if err != nil { log.Fatal(err) }
  return body
}


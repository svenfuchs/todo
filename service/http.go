package service

import (
  "bytes"
  "errors"
  "fmt"
  "io/ioutil"
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
    uri, _ = appendQueryString(uri, *query) // TODO report error
  }
  return HttpRequest{ method, uri, headers, body }
}

func (h HttpRequest) run() ([]byte, error) {
  var body []byte
  request, err := h.newRequest()
  if err != nil { return body, err }

  response, err := http.DefaultClient.Do(request)
  defer response.Body.Close()
  if err != nil { return body, err }

  err = h.checkResponseStatus(response)
  if err != nil { return body, err }

  return h.readResponseBody(response)
}

func (h HttpRequest) newRequest() (*http.Request, error) {
  request, err := http.NewRequest(h.method, h.uri, h.body)
  if err != nil { return nil, err }

  h.setHeaders(request, h.headers)
  return request, nil
}

func (h HttpRequest) setHeaders(request *http.Request, headers *map[string]string) {
  for key, value := range *headers {
    request.Header.Add(key, value)
  }
}

func appendQueryString(uri string, data map[string]string) (string, error) {
  u, err := url.ParseRequestURI(uri)
  if err != nil { return uri, err }
  u.RawQuery = toQueryString(data)
  return fmt.Sprintf("%v", u), nil
}

func toQueryString(data map[string]string) string {
  query := url.Values{}
  for key, value := range data {
    query.Set(key, value)
  }
  return query.Encode()
}

func (h HttpRequest) checkResponseStatus(r *http.Response) error {
  if r.StatusCode / 100 == 2 {
    return nil
  }

  body, err := ioutil.ReadAll(r.Body)
  if err != nil { return err }

  msg := "%s request to %s failed: (%d) %s\n"
  return errors.New(fmt.Sprintf(msg, h.method, h.uri, r.StatusCode, string(body)))
}

func (h HttpRequest) readResponseBody(response *http.Response) ([]byte, error) {
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    msg := "Could not read response body for %s request to %s. Status: %d Error: %q\n"
    return body, errors.New(fmt.Sprintf(msg, h.method, h.uri, response.StatusCode, err))
  }
  return body, err
}


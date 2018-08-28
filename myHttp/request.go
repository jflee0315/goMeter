package myHttp

import (
    "goMeter/config"
    "net/http"
    "io/ioutil"
    "strings"
    "time"
    "sync"
    )

var once sync.Once
var client *http.Client
// For Lazy Loading
func getClient() *http.Client {
    once.Do(func() {
        client = &http.Client{Timeout: time.Second * 20}
    })
    return client
}

type MyResponse interface {
    StatusCode() int
    Body() string
}

type response struct {
    fullResponse *http.Response
    body string
}

func (r response) StatusCode() int{
    return r.fullResponse.StatusCode
}

func (r response) Body() string{
    return r.body
}

func MakeRequest(request config.HttpRequest) (MyResponse, error) {
    client := getClient()
    var resp *http.Response
    var err error
    switch request.Method {
        case http.MethodGet: 
            resp, err = client.Get(request.Url)
        case http.MethodPost:
            resp, err = client.Post(request.Url, request.ContentType, strings.NewReader(request.Body))
        default:
            panic("no such http method")
    }
    if(err != nil) {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    myResponse := response{body:string(body), fullResponse: resp}

    return myResponse, err
}
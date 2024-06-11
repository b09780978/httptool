package httptool

import "net/http"

type HttpClient struct {
	client *http.Client
}

var DefaultClient = &HttpClient{
	client: &http.Client{},
}

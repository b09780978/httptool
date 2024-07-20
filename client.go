package httptool

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/iunary/fakeuseragent"
)

/*
User-Agents
default use chrome
*/
var Chrome = fakeuseragent.GetUserAgent(fakeuseragent.BrowserChrome)
var Firefox = fakeuseragent.GetUserAgent(fakeuseragent.BrowserFirefox)
var Edge = fakeuseragent.GetUserAgent(fakeuseragent.BrowserEdge)
var Safari = fakeuseragent.GetUserAgent(fakeuseragent.BrowserSafari)

// HTTP method
var GET string = "GET"
var POST string = "POST"
var PUT string = "PUT"
var DELETE string = "DELETE"
var HEAD string = "HEAD"
var OPTIONS string = "OPTIONS"
var TRACE string = "TRACE"

func NewPostData() url.Values {
	return url.Values{}
}

type HttpClient struct {
	Client    *http.Client
	UserAgent string
	Header    http.Header
	Cookies   []*http.Cookie
}

func NewClient() *HttpClient {
	return &HttpClient{
		Client:    &http.Client{},
		UserAgent: Chrome,
		Header:    http.Header{},
		Cookies:   make([]*http.Cookie, 0),
	}
}

var DefaultClient = NewClient()

// set fake user-agent
func (c *HttpClient) AddFakeUserAgent() {
	c.SetHeader("User-Agent", c.UserAgent)
}

// Header
func (c *HttpClient) SetHeader(k, v string) {
	c.Header.Set(k, v)
}

func (c *HttpClient) AddHeader(k, v string) {
	c.Header.Add(k, v)
}

func (c *HttpClient) DelHeader(k string) {
	c.Header.Del(k)
}

func (c *HttpClient) CloneHeader() http.Header {
	return c.Header.Clone()
}

func (c *HttpClient) GetHeader(k string) string {
	return c.Header.Get(k)
}

func (c *HttpClient) GetHeadersArray(k string) []string {
	return c.Header.Values(k)
}

// Cookie

func (c *HttpClient) SetCookie(cookie *http.Cookie) {
	for i, co := range c.Cookies {
		if co.Name == cookie.Name {
			c.Cookies[i] = cookie
			return
		}
	}

	c.Cookies = append(c.Cookies, cookie)
}

func (c *HttpClient) GetCookie(k string) (*http.Cookie, bool) {
	for i, cookie := range c.Cookies {
		if cookie.Name == k {
			return c.Cookies[i], true
		}
	}
	return nil, false
}

func (c *HttpClient) DelCookie(k string) bool {
	for i, cookie := range c.Cookies {
		if cookie.Name == k {
			c.Cookies = append(c.Cookies[:i], c.Cookies[i+1:]...)
			return true
		}
	}
	return false
}

// when NewRequest add custom header
func (c *HttpClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	for k, v := range c.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	return req, nil
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *HttpClient) Get(url string) (*http.Response, error) {
	req, err := c.NewRequest(GET, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HttpClient) Post(url string, params url.Values) (*http.Response, error) {
	req, err := c.NewRequest(POST, url, strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HttpClient) PostForm(url string, params url.Values) (*http.Response, error) {
	req, err := c.NewRequest(POST, url, strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return DefaultClient.NewRequest(method, url, body)
}

func Get(url string) (*http.Response, error) {
	return DefaultClient.Get(url)
}

func Post(url string, params url.Values) (*http.Response, error) {
	return DefaultClient.Post(url, params)
}

func PostForm(url string, param url.Values) (*http.Response, error) {
	return DefaultClient.PostForm(url, param)
}

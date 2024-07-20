package httptool

import (
	"io"
	"net/http"
)

// Read response
func ReadResp(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, err
}

func ReadRespStr(resp *http.Response) (string, error) {
	data, err := ReadResp(resp)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// parse cookies
func ParseCookies(resp *http.Response) []*http.Cookie {
	header := http.Header{}

	if cookie, ok := resp.Header["Set-Cookie"]; ok {
		for _, v := range cookie {
			header.Add("Cookie", v)
		}
	} else {
		return []*http.Cookie{}
	}

	req := http.Request{Header: header}

	return req.Cookies()
}

func SetCookiesHeader(req *http.Request, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		req.Header.Add("Set-Cookie", cookie.String())
	}
}

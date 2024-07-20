package httptool

import (
	"net/http"
	"testing"
)

func TestParseCookieHeader(t *testing.T) {
	fakeHeader := http.Header{}
	fakeHeader.Add("Set-Cookie", "webp=1;   path=/; PHPSESSIONID=21399f0f9f")

	fakeResp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     fakeHeader,
	}

	cookies := ParseCookies(fakeResp)

	for _, cookie := range cookies {
		if cookie.Name == "webp" && cookie.Value == "1" {
			return
		}
	}

	t.Error("Set value fail")
}

package httptool

import (
	"io"
	"net/http"
)

func ReadResp(resp *http.Response) (*[]byte, error) {
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &data, err
}

func ReadRespStr(resp *http.Response) (string, error) {
	data, err := ReadResp(resp)

	if err != nil {
		return "", err
	}

	return string(*data), nil
}

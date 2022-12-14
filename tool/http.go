package tool

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string) (io.ReadCloser, error) {
	c := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("403")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}

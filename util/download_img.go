package util

import (
	"fmt"
	"net/http"
)

func DownloadImg(url string) (*http.Response, error) {
	fmt.Println("URL:", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

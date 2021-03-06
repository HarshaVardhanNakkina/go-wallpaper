package util

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

func DownloadImg(url string) (*http.Response, error) {
	fmt.Println("URL:", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ExtractFileExt(resp *http.Response) string {
	fileExt := "jpeg"
	contentType := resp.Header.Get("Content-Type")
	imgExtRegex := regexp.MustCompile(`(?i)(jpeg|jpg|png)`)
	imgExt := imgExtRegex.FindString(contentType)
	if imgExt != "" {
		fileExt = imgExt
	}

	return fileExt
}

func FileTypeCheck(resp *http.Response) (bool, error) {
	contentType := resp.Header.Get("Content-Type")
	contentTypeRegex := regexp.MustCompile(`(?i)(jpeg|jpg|png)`)
	if contentTypeRegex.MatchString(contentType) {
		return true, nil
	}
	return false, errors.New("downloaded file is not an image")
}

package unsplash

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"regexp"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set_wallpaper"
	"github.com/HarshaVardhanNakkina/go-wallpaper/util"
)

var imgNotFound = "https://images.unsplash.com/source-404?fit=crop&fm=jpg&h=800&q=60&w=1200"
var unsplashUrl string = "https://source.unsplash.com"
var defaultRes string = "1920x1080"

func DownloadFromUnsplash(resolution, tag string) error {
	fmt.Println("Downloading image from unsplash.com")
	if resolution != "" {
		defaultRes = resolution
	}

	if tag == "" {
		unsplashUrl = fmt.Sprintf("%v/%v/", unsplashUrl, defaultRes)
	} else {
		unsplashUrl = fmt.Sprintf("%v/%v/?%v", unsplashUrl, defaultRes, tag)
	}

	resp, err := util.DownloadImg(unsplashUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	_, err = check404Error(reqUrl)
	if err != nil {
		return err
	}

	fileExt := "jpeg"
	contentType := resp.Header.Get("Content-Type")
	imgExtRegex := regexp.MustCompile(`(?i)(jpeg|jpg|png)`)
	imgExt := imgExtRegex.FindString(contentType)
	if imgExt != "" {
		fileExt = imgExt
	}

	rawImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	return setwallpaper.SetWallpaper(filename, rawImg)

}

func check404Error(reqUrl *url.URL) (*url.URL, error) {
	if reqUrl.String() == imgNotFound {
		return nil, errors.New("no image found with the given tag/resolution")
	}
	return reqUrl, nil
}

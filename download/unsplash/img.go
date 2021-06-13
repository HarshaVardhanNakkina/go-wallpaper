package unsplash

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

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
	if !util.FileTypeCheck(resp) {
		return errors.New("dowloaded file is not an image")
	}

	reqUrl := resp.Request.URL
	_, err = check404Error(reqUrl)
	if err != nil {
		return err
	}

	fileExt := util.ExtractFileExt(resp)

	rawImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	return setwallpaper.SetWallpaper(filename, rawImg)

}

// TODO, change this function to return a bool, like FileTypeCheck
func check404Error(reqUrl *url.URL) (*url.URL, error) {
	if reqUrl.String() == imgNotFound {
		return nil, errors.New("no image found with the given tag/resolution")
	}
	return reqUrl, nil
}

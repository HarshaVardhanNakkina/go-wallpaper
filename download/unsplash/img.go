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

func DownloadFromUnsplash(resolution, tag string) error {
	fmt.Println("Downloading image from unsplash.com")

	if tag == "" {
		unsplashUrl = fmt.Sprintf("%v/%v/", unsplashUrl, resolution)
	} else {
		unsplashUrl = fmt.Sprintf("%v/%v/?%v", unsplashUrl, resolution, tag)
	}

	resp, err := util.DownloadImg(unsplashUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := util.FileTypeCheck(resp); err != nil {
		return err
	}

	reqUrl := resp.Request.URL
	if _, err = check404Error(reqUrl); err != nil {
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

func check404Error(reqUrl *url.URL) (bool, error) {
	if reqUrl.String() == imgNotFound {
		return false, errors.New("no image found with the given tag/resolution")
	}
	return true, nil
}

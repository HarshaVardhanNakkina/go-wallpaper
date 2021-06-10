package unsplash

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set-wallpaper"
	homedir "github.com/mitchellh/go-homedir"
)

var imgNotFound = "https://images.unsplash.com/source-404?fit=crop&fm=jpg&h=800&q=60&w=1200"
var defaultFileMode fs.FileMode = 0644
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

	rawImg, fileExt, err := downloadImg(unsplashUrl)
	if err != nil {
		return err
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	fileLocation := filepath.Join(homeDir, "Pictures", "go-wallpaper")
	filepath := filepath.Join(fileLocation, filename)
	createDirIfNotExists(fileLocation)

	fmt.Println("Saving image @", filepath)
	err = ioutil.WriteFile(filepath, rawImg, defaultFileMode)
	if err != nil {
		return err
	}
	setwallpaper.SetWallpaper(filepath)

	return nil
}

func downloadImg(url string) ([]byte, string, error) {
	fmt.Println("URL:", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, "", err
	}

	defer resp.Body.Close()

	reqURL := resp.Request.URL
	if reqURL.String() == imgNotFound {
		return nil, "", errors.New("no image found with the given tag/resolution")
	}

	contentType := resp.Header.Get("Content-Type")
	fileExt := strings.Split(contentType, "/")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return body, fileExt[len(fileExt)-1], nil
}

func createDirIfNotExists(fileLocation string) error {
	_, err := os.Stat(fileLocation)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(fileLocation, defaultFileMode)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

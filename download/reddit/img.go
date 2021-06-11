package reddit

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set-wallpaper"
	util "github.com/HarshaVardhanNakkina/go-wallpaper/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type Result struct {
	Data Children `json:"data"`
}

type Children struct {
	ChildrenArr []ChildData `json:"children"`
}

type ChildData struct {
	ChildObj ChildInfo `json:"data"`
}

type ChildInfo struct {
	Title                 string `json:"title"`
	SubredditNamePrefixed string `json:"subreddit_name_prefixed"`
	Url                   string `json:"url"`
}

var redditUrl string = "https://www.reddit.com/r/wallpaper/new.json"
var userAgent string = "win64:github.com/HarshaVardhanNakkina/go-wallpaper:/u/harsha602"
var defaultFileMode fs.FileMode = 0644

func DownloadFromReddit() error {
	url := redditUrl
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		cobra.CheckErr(err)
	}

	req.Header.Set("user-agent", userAgent)
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var data Result
	json.Unmarshal(body, &data)
	randInd := util.GetRandomNum(len(data.Data.ChildrenArr))
	pick := data.Data.ChildrenArr[randInd]

	rawImg, fileExt, err := downloadImg(pick.ChildObj.Url)
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

	fileExt := filepath.Ext(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return body, fileExt, nil
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

package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set_wallpaper"
	util "github.com/HarshaVardhanNakkina/go-wallpaper/util"
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

func DownloadFromReddit() error {
	fmt.Println("Downloading image from unsplash.com")
	url := redditUrl
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
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

	images := data.Data.ChildrenArr
	randInd := util.GetRandomNum(len(images))
	pick := images[randInd]
	imgUrl := pick.ChildObj.Url

	resp, err := util.DownloadImg(imgUrl)
	if err != nil {
		return err
	}

	fileExt := filepath.Ext(imgUrl)

	rawImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	return setwallpaper.SetWallpaper(filename, rawImg)

}

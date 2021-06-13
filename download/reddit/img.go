package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set_wallpaper"
	util "github.com/HarshaVardhanNakkina/go-wallpaper/util"
)

var redditUrl string = "https://www.reddit.com/r"
var sort string = "new"
var userAgent string = "/u/harsha602"

var subreddits []string = []string{"EarthPorn", "wallpaper", "wallpapers", "multiwall"}

func DownloadFromReddit() error {
	randIdx := util.GetRandomNum(len(subreddits))
	subreddit := subreddits[randIdx]
	url := fmt.Sprintf("%v/%v/%v/.json", redditUrl, subreddit, sort)
	fmt.Println("Downloading image from /r/", subreddit)

	client := &http.Client{}
	response, err := httpRequest(client, "GET", url)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	children := result["data"].(map[string]interface{})["children"]

	var imgData []map[string]interface{}
	for _, child := range children.([]interface{}) {
		// type casting "data" is possible when assigning itself
		data := child.(map[string]interface{})["data"]
		if val, ok := data.(map[string]interface{})["post_hint"]; ok && val == "image" {
			imgData = append(imgData, data.(map[string]interface{}))
		}
	}

	targetImgs := []string{}
	for _, img := range imgData {
		previewImgs := img["preview"].(map[string]interface{})["images"]
		firstImg := previewImgs.([]interface{})[0]
		source := firstImg.(map[string]interface{})["source"]
		imgUrl := source.(map[string]interface{})["url"]
		targetImgs = append(targetImgs, html.UnescapeString((imgUrl.(string))))
	}

	randIdx = util.GetRandomNum(len(targetImgs))
	targetImg := targetImgs[randIdx]

	resp, err := httpRequest(client, "GET", targetImg)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if !util.FileTypeCheck(resp) {
		return errors.New("dowloaded file is not an image")
	}

	fileExt := util.ExtractFileExt(resp)
	rawImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	return setwallpaper.SetWallpaper(filename, rawImg)

}

func httpRequest(client *http.Client, method, url string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// userAgent is global
	req.Header.Set("user-agent", userAgent)
	return client.Do(req)

}

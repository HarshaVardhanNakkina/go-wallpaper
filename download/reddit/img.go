package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"

	setwallpaper "github.com/HarshaVardhanNakkina/go-wallpaper/set_wallpaper"
	util "github.com/HarshaVardhanNakkina/go-wallpaper/util"
	"github.com/imroc/req"
)

// var redditUrl string = "https://www.reddit.com/r"
var redditSearchUrl string = "https://reddit.com/search.json?q=subreddit:"

var subreddits []string = []string{"EarthPorn", "wallpaper", "wallpapers", "multiwall"}

func DownloadFromReddit(sort string) error {
	randIdx := util.GetRandomNum(len(subreddits))
	subreddit := subreddits[randIdx]
	// url := fmt.Sprintf("%v/%v/%v/.json", redditUrl, subreddit, sort)
	url := fmt.Sprintf("%v(%v)+self:no&sort=%v&limit=15&type=t3&restrict_sr=true", redditSearchUrl, subreddit, sort)
	fmt.Printf("Downloading image from /r/%v\n", subreddit)

	// client := &http.Client{
	// 	Timeout: time.Second * 25,
	// }
	// response, err := httpRequest(client, "GET", url)
	r, err := req.Get(url)
	if err != nil {
		return err
	}
	defer r.Response().Body.Close()
	if r.Response().StatusCode != 200 {
		return errors.New(r.Response().Status)
	}

	body, err := ioutil.ReadAll(r.Response().Body)
	if err != nil {
		return err
	}

	targetImgs := getImageUrls(body)
	randIdx = util.GetRandomNum(len(targetImgs))
	targetImg := targetImgs[randIdx]

	// r, err = httpRequest(client, "GET", targetImg)
	r, err = req.Get(targetImg)
	if err != nil {
		return err
	}
	defer r.Response().Body.Close()

	if _, err := util.FileTypeCheck(r.Response()); err != nil {
		return err
	}

	fileExt := util.ExtractFileExt(r.Response())
	rawImg, err := ioutil.ReadAll(r.Response().Body)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("wallpaper.%v", fileExt)
	return setwallpaper.SetWallpaper(filename, rawImg)

}

func getImageUrls(body []byte) []string {
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

	return targetImgs
}

// func httpRequest(client *http.Client, method, url string) (*http.Response, error) {
// 	req, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// userAgent is global
// 	// req.Header.Set("User-Agent", userAgent)
// 	return client.Do(req)

// }

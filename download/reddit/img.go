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

var redditSearchUrl string = "https://reddit.com/search.json?q=subreddit:"
var subreddits []string = []string{"EarthPorn", "wallpaper", "wallpapers", "multiwall"}

func DownloadFromReddit(sort, top string) error {
	if sort == "" && top == "" {
		return errors.New("invalid flag value provided")
	}
	randIdx := util.GetRandomNum(len(subreddits))
	subreddit := subreddits[randIdx]
	url := constructURL(subreddit, sort, top)
	fmt.Printf("Downloading image from /r/%v\n", subreddit)

	headers := req.Header{
		"User-Agent": "go-wallpaper:v0.4.0:/u/harsha602",
	}

	r, err := req.Get(url, headers)
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

func constructURL(subreddit, sort, top string) string {
	commonParams := "limit=20&type=t3&restrict_sr=true"
	url := fmt.Sprintf("%v(%v)+self:no", redditSearchUrl, subreddit)

	if top != "" {
		return fmt.Sprintf("%v&t=%v&%v", url, top, commonParams)
	}

	if sort != "" {
		return fmt.Sprintf("%v&sort=%v&%v", url, sort, commonParams)
	}

	return fmt.Sprintf("%v&%v", url, commonParams)
}

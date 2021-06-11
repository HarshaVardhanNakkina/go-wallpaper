/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

// redditCmd represents the reddit command
var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reddit called")
		url := "https://www.reddit.com/r/wallpaper/new.json"
		client := &http.Client{
			Timeout: time.Second * 120,
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			cobra.CheckErr(err)
		}
		req.Header.Set("user-agent", "win64:github.com/HarshaVardhanNakkina/go-wallpaper:/u/harsha602")
		response, err := client.Do(req)
		if err != nil {
			cobra.CheckErr(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			cobra.CheckErr(err)
		}
		//Convert the body to type string
		// sb := string(body)
		// fmt.Printf("Results String: %v\n", sb)

		var data Result
		json.Unmarshal(body, &data)
		// fmt.Printf("Results: %v\n", data)
		for _, child := range data.Data.ChildrenArr {
			fmt.Println(child.ChildObj)
		}
	},
}

func init() {
	rootCmd.AddCommand(redditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	reddit "github.com/HarshaVardhanNakkina/go-wallpaper/download/reddit"
	"github.com/spf13/cobra"
)

var sort string
var top string

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Set wallpaper from Reddit",
	Long:  `Selects a random wallpaper from subreddits like /r/wallpaper, /r/wallpapers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return reddit.DownloadFromReddit(sort, top)
	},
}

func init() {
	rootCmd.AddCommand(redditCmd)

	// TODO: maybe convert these flags into sub-commands
	redditCmd.Flags().StringVarP(&sort, "sort", "s", "", `Choose from "new", "hot", and "top" sections,
"top" option defaults to today's top
use top flag instead, for multiple options
`)
	redditCmd.Flags().StringVarP(&top, "top", "t", "", `today - picks from today's top
 week - picks from week's top
month - picks from month's top
 year - picks from year's top
	all - picks from alltime top
this takes priority over "sort" flag
	`)
}

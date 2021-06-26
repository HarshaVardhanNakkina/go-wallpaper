package cmd

import (
	reddit "github.com/HarshaVardhanNakkina/go-wallpaper/download/reddit"
	"github.com/spf13/cobra"
)

var sort string

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Set wallpaper from Reddit",
	Long:  `Selects a random wallpaper from subreddits like /r/wallpaper, /r/wallpapers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return reddit.DownloadFromReddit(sort)
	},
}

func init() {
	rootCmd.AddCommand(redditCmd)

	// TODO: maybe convert these flags into sub-commands
	redditCmd.Flags().StringVarP(&sort, "sort", "s", "new", `Choose from "new", "hot", and "top" sections
	"top" option defaults today's top
	separate flag for "top" has to be implemented yet`)
}

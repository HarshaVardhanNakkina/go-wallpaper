package cmd

import (
	reddit "github.com/HarshaVardhanNakkina/go-wallpaper/download/reddit"
	"github.com/spf13/cobra"
)

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Set wallpaper from Reddit",
	Long:  `Selects a random wallpaper from new wallpapers uploaded to several different subreddits`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return reddit.DownloadFromReddit()
	},
}

func init() {
	rootCmd.AddCommand(redditCmd)
}

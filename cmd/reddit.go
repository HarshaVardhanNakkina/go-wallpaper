package cmd

import (
	reddit "github.com/HarshaVardhanNakkina/go-wallpaper/download/reddit"
	"github.com/spf13/cobra"
)

var sort string

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Set wallpaper from Reddit",
	Long:  `Selects a random wallpaper from new wallpapers uploaded to several different subreddits`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return reddit.DownloadFromReddit(sort)
	},
}

func init() {
	rootCmd.AddCommand(redditCmd)
	// rootCmd.PersistentFlags().StringVarP(&Resolution, "resolution", "r", "", "Resolution of wallpaper (defaults to 1920x1080)")
	redditCmd.Flags().StringVarP(&sort, "sort", "s", "new", `Choose from "new", "hot", and "top" sections`)
}

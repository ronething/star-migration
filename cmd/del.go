package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var unStarCmd = &cobra.Command{
	Use:   "unstar",
	Short: "UnStar a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")

		err := starring.UnStarRepo(owner, repo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Repo %s unstarred successfully.\n", repo)
	},
}

func init() {
	rootCmd.AddCommand(unStarCmd)

	unStarCmd.Flags().StringP("owner", "o", "", "Repository owner")
	unStarCmd.Flags().StringP("repo", "r", "", "Repository name")
}

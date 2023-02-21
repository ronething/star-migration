package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Star a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		err := starring.StarRepo(owner, repo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Starred %s/%s\n", owner, repo)
	},
}

func init() {
	rootCmd.AddCommand(starCmd)

	starCmd.Flags().StringP("owner", "o", "", "Repository owner")
	starCmd.Flags().StringP("repo", "r", "", "Repository name")
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "unstar repos from file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")

		repos, err := starring.ReadStarredReposFromFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			err := starring.UnStarRepo(repo.Owner, repo.Repo)
			if err != nil {
				fmt.Printf("Failed to Unstar repo %s/%s: %s\n", repo.Owner, repo.Repo, err)
			} else {
				fmt.Printf("UnStarred repo %s/%s.\n", repo.Owner, repo.Repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)

	eraseCmd.Flags().StringP("filename", "f", "erase.txt", "erase filename")
}

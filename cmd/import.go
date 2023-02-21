package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "add star repos from file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")

		repos, err := starring.ReadStarredReposFromFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			err := starring.StarRepo(repo.Owner, repo.Repo)
			if err != nil {
				fmt.Printf("Failed to star repo %s/%s: %s\n", repo.Owner, repo.Repo, err)
			} else {
				fmt.Printf("Starred repo %s/%s.\n", repo.Owner, repo.Repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringP("filename", "f", "import.txt", "import filename")
}

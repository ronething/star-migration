package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list current user star repos",
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetInt64("page")
		repos, err := starring.GetStarredRepos(page)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			fmt.Printf("%s/%s: %s\n", strings.Split(repo.Url, "/")[3], repo.Name, repo.Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().Int64P("page", "p", 1, "page")
}

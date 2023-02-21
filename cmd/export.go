package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"star-migration/starring"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export star repos to file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")

		repos, err := starring.GetAllStarredRepos()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = starring.ExportStarredReposToFile(filename, repos)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Starred repos exported to %s.\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringP("filename", "f", "export.txt", "export filename")
}

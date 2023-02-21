package starring

import (
	"fmt"
	"os"
	"strings"
)

func ExportStarredReposToFile(filename string, repos []Repo) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, repo := range repos {
		line := fmt.Sprintf("%s/%s %s\n", strings.Split(repo.Url, "/")[3], repo.Name, repo.Url)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

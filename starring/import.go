package starring

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StarRepoInfo struct {
	Owner string
	Repo  string
}

func ReadStarredReposFromFile(filename string) ([]StarRepoInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var repos []StarRepoInfo

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line in file: %s", scanner.Text())
		}

		info := strings.Split(parts[0], "/")
		if len(info) != 2 {
			return nil, fmt.Errorf("invalid line in info: %s", scanner.Text())
		}
		owner := info[0]
		repo := strings.TrimSpace(info[1])

		repos = append(repos, StarRepoInfo{
			Owner: owner,
			Repo:  repo,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return repos, nil
}

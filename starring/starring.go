package starring

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Repo struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

func GetAllStarredRepos() ([]Repo, error) {
	var allRepos []Repo
	var page int64 = 1
	for {
		log.Printf("get page: %v star repos", page)
		repos, err := GetStarredRepos(page)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if len(repos) < 100 {
			break
		}
		time.Sleep(300 * time.Millisecond)
		page++
	}

	return allRepos, nil
}

func GetStarredRepos(page int64) ([]Repo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/user/starred?page=%d&per_page=100", page), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch starred repos: %s", resp.Status)
	}

	var repos []Repo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func StarRepo(owner, repo string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.github.com/user/starred/%s/%s", owner, repo), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))
	req.Header.Set("Accept", "application/vnd.github.star+json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to star the repo: %s", resp.Status)
	}

	return nil
}

func UnStarRepo(owner, repo string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.github.com/user/starred/%s/%s", owner, repo), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))
	req.Header.Set("Accept", "application/vnd.github.star+json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to unstar repo %s/%s: %s", owner, repo, resp.Status)
	}

	return nil
}

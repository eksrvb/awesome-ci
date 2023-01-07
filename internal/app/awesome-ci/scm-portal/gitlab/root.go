package gitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// NewGitLabClient Creates a new GitLab Client
// Needs the Environment Variables: GITHUB_TOKEN
// Needs the optional Environment Variables: GITHUB_ENTERPRISE_SERVER_URL
func NewGitLabClient(serverUrl *string, repoUrl *string, token *string) (glrc *GitLabRichClient, err error) {
	glrc.Client, err = gitlab.NewClient(*token, gitlab.WithBaseURL(*serverUrl))
	if err != nil {
		return nil, fmt.Errorf("error at initializing gitlab client: %v", err)
	}
	return
}

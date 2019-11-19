package entity

import (
	"encoding/json"
	"fmt"
)

type (
	GitHub struct {
		Action       GitHubActionType `json:"action"`
		*Comment     `json:"comment",omitempty`
		*PullRequest `json:"pull_request",omitempty`
	}

	Comment struct {
		HTMLURL string `json:"html_url"`
		Body    string `json:"body"`
	}

	PullRequest struct {
		HTMLURL  string `json:"html_url"`
		Assignee struct {
			Login string `json:"login"`
		} `json:"assignee"`
		Assignees []struct {
			Login string `json:"login"`
		} `json:"assignees"`
	}
)

type GitHubActionType int

const (
	GitHubActionTypeUnknown = iota
	GitHubActionTypeCreated
	GitHubActionTypeAssigned
)

func (enum *GitHubActionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}

	var e GitHubActionType
	switch s {
	case "created":
		e = GitHubActionTypeCreated
	case "assigned":
		e = GitHubActionTypeAssigned
	default:
		// TODO: Catch unknown case
		e = GitHubActionTypeUnknown
	}
	*enum = e
	return nil
}

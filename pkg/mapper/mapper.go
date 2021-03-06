package mapper

import (
	"fmt"
	"log"

	"github.com/bannzai/notifier/pkg/logger"
	"github.com/bannzai/notifier/pkg/parser"
	"github.com/bannzai/notifier/pkg/sender"
)

type Mapper struct{}

func New() Mapper {
	return Mapper{}
}

func noMatchedError(users []User, content parser.Content, toContentType sender.ContentType) error {
	return fmt.Errorf("Not matched id from content of %v, to %d, with users %v", content, toContentType, users)
}

func (Mapper) MapIDs(content parser.Content, toContentType sender.ContentType) ([]string, error) {
	users, err := fetchUsers()
	if err != nil {
		return []string{}, fmt.Errorf("fetchIDMap error %w", err)
	}
	switch content.ContentType {
	case parser.GitHubMentionContent, parser.GitHubAssignedContent, parser.GitHubRequestReviewedContent:
		ids := []string{}
		logger.Logf("content.UserNames = %+v\n", content.UserNames)
		for _, username := range content.UserNames {
			slack, ok := extractUserFromGitHub(users, username, toContentType)

			if !ok {
				return ids, noMatchedError(users, content, toContentType)
			}

			ids = append(ids, slack.ID)
		}
		return ids, nil
	default:
		log.Printf("Unexpected pattern for content type %d", content.ContentType)
		return []string{}, noMatchedError(users, content, toContentType)
	}
}

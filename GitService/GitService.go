package GitService

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// git service
type GitService struct {
	azureDevOpsUri     string
	repositoryId       string
	authorizationToken string
	client             *http.Client
}

// new git service
func NewGitService(azureDevOpsUri string, repositoryId string, authorizationToken string) *GitService {
	return &GitService{
		azureDevOpsUri:     azureDevOpsUri,
		repositoryId:       repositoryId,
		authorizationToken: authorizationToken,
		client:             &http.Client{},
	}
}

// get commit changes
func (s *GitService) GetChanges(commitId string) (GitChanges, error) {
	var changes GitChanges

	requestUrl := fmt.Sprintf("%v/_apis/git/repositories/%v/commits/%v/changes?api-version=5.0", s.azureDevOpsUri, s.repositoryId, commitId)
	bytes, err := s.getRequest(requestUrl)
	if err != nil {
		errorText := fmt.Sprintf("get changes for commit: %v %v", commitId, err.Error())
		return changes, errors.New(errorText)
	}

	err = json.Unmarshal(bytes, &changes)
	if err != nil {
		errorText := fmt.Sprintf("parse changes for commit: %v %v", commitId, err.Error())
		return changes, errors.New(errorText)
	}

	return changes, nil
}

// get item
func (s *GitService) GetItem(itemUrl string) ([]byte, error) {
	bytes, err := s.getRequest(itemUrl)
	if err != nil {
		errorText := fmt.Sprintf("get item: %v %v", itemUrl, err.Error())
		return nil, errors.New(errorText)
	}

	return bytes, nil
}

// get request
func (s *GitService) getRequest(requestUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		errorText := fmt.Sprintf("new request: %v %v", requestUrl, err.Error())
		return nil, errors.New(errorText)
	}
	req.SetBasicAuth(s.authorizationToken, s.authorizationToken)

	resp, err := s.client.Do(req)
	if err != nil {
		errorText := fmt.Sprintf("get request: %v %v", requestUrl, err.Error())
		return nil, errors.New(errorText)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorText := fmt.Sprintf("read request: %v %v", requestUrl, err.Error())
		return nil, errors.New(errorText)
	}

	return bytes, nil
}

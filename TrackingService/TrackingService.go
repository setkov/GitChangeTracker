package TrackingService

import (
	"fmt"
	"os"
	"path/filepath"

	"main.go/Common"
	"main.go/GitService"
)

// tracking service
type TrackingService struct {
	parameters *Common.Parameters
	gitService *GitService.GitService
}

// new tracking service
func NewTrackingService(parameters *Common.Parameters) *TrackingService {
	return &TrackingService{
		parameters: parameters,
		gitService: GitService.NewGitService(parameters.AzureDevOpsUri, parameters.RepositoryId, parameters.АuthorizationToken),
	}
}

// track changes
func (s *TrackingService) Track() error {
	changes, err := s.gitService.GetChanges(s.parameters.CommitId)
	if err != nil {
		return err
	}

	err = s.saveChanges(changes)
	if err != nil {
		return err
	}

	return nil
}

// save changes
func (s *TrackingService) saveChanges(changes GitService.GitChanges) error {
	for _, change := range changes.Changes {
		if change.Item.GitObjectType == "blob" {
			fileName := filepath.Base(change.Item.Path)
			if filepath.Ext(fileName) == ".sql" {
				fmt.Println(fileName)

				data, err := s.gitService.GetItem(change.Item.Url)
				if err != nil {
					return err
				}

				err = os.WriteFile(fileName, data, 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

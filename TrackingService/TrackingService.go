package TrackingService

import (
	"fmt"

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

func (s *TrackingService) Track() error {
	changes, err := s.gitService.GetChanges(s.parameters.CommitId)
	if err != nil {
		return err
	} else {
		fmt.Println("changes", changes)
	}

	return nil
}

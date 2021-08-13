package TrackingService

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"main.go/Common"
	"main.go/GitService"
	"main.go/SqlParser"
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
	fmt.Printf("commitId: %v\n", s.parameters.CommitId)

	// get commit
	commit, err := s.gitService.GetCommit(s.parameters.CommitId)
	if err != nil {
		return err
	}
	commitInfo := "Powered by GitChangeTracker (https://github.com/setkov/GitChangeTracker)\n\n"
	commitInfo += fmt.Sprintf("commit:  %v\n", commit.CommitId)
	commitInfo += fmt.Sprintf("author:  %v <%v>\n", commit.Author.Name, commit.Author.Email)
	commitInfo += fmt.Sprintf("date:    %s\n", commit.Author.Date)
	commitInfo += fmt.Sprintf("comment: %v\n", commit.Comment)

	// get changes
	changes, err := s.gitService.GetChanges(s.parameters.CommitId)
	if err != nil {
		return err
	}

	// make SqlObjects
	var sqlObjects SqlParser.SqlObjects
	for _, change := range changes.Changes {
		if change.Item.GitObjectType == "blob" && strings.HasSuffix(change.Item.Path, ".sql") {
			sqlObject := SqlParser.SqlObject{
				Path: change.Item.Path,
			}
			fmt.Printf("file: %v\n", sqlObject.Path)

			data, err := s.gitService.GetItem(change.Item.Url)
			if err != nil {
				return err
			}
			sqlObject.Code = string(bytes.Trim(data, "\xef\xbb\xbf")) // remove BOM marker

			sqlObjects.Objects = append(sqlObjects.Objects, sqlObject)
		}
	}

	// parse to sql script
	parser := SqlParser.NewSqlParser(&sqlObjects, s.parameters.OutputPath)
	sqlScript := parser.Parse(commitInfo)
	//fmt.Println(sqlScript)

	// save sql script
	fileName := s.parameters.OutputPath + s.parameters.CommitId + ".sql"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(sqlScript)
	fmt.Printf("save sql script: %v\n", fileName)

	return nil
}

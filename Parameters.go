package main

import (
	"errors"
	"fmt"
	"os"
)

type Parameters struct {
	AzureDevOpsUri     string
	RepositoryId       string
	CommitId           string
	АuthorizationToken string
	OutputPath         string
}

// New parameters
func NewParameters() (*Parameters, error) {
	count := len(os.Args) - 1
	if count != 5 {
		errorText := fmt.Sprintf("wrong argements count=%v, run with parameters: AzureDevOpsUri, RepositoryId, CommitId, АuthorizationToken, OutputPath", count)
		return nil, errors.New(errorText)
	}

	return &Parameters{
		AzureDevOpsUri:     os.Args[1],
		RepositoryId:       os.Args[2],
		CommitId:           os.Args[3],
		АuthorizationToken: os.Args[4],
		OutputPath:         os.Args[5],
	}, nil
}

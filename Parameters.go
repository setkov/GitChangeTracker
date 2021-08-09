package main

import (
	"errors"
	"fmt"
	"os"
)

type Parameters struct {
	CommitId           string
	RepositoryPath     string
	АuthorizationToken string
	OutputPath         string
}

// New parameters
func NewParameters() (*Parameters, error) {
	count := len(os.Args) - 1
	if count != 4 {
		errorText := fmt.Sprintf("wrong argements count=%v, run with parameters: CommitId, RepositoryPath, АuthorizationToken, OutputPath", count)
		return nil, errors.New(errorText)
	}

	return &Parameters{
		CommitId:           os.Args[1],
		RepositoryPath:     os.Args[2],
		АuthorizationToken: os.Args[3],
		OutputPath:         os.Args[4],
	}, nil
}

package main

import (
	"fmt"
	"log"

	"main.go/GitService"
)

func main() {
	parameters, err := NewParameters()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("parameters", parameters)
	}

	gitService := GitService.NewGitService(parameters.AzureDevOpsUri, parameters.RepositoryId, parameters.АuthorizationToken)
	changes, err := gitService.GetChanges(parameters.CommitId)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("changes", changes)
	}
}

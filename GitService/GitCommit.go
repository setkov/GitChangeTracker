package GitService

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type GitCommit struct {
	CommitId string `json:"commitId"`
	Author   Author `json:"author"`
	Comment  string `json:"comment"`
}

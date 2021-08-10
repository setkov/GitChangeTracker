package GitService

type GitItem struct {
	GitObjectType string `json:"gitObjectType"`
	Path          string `json:"path"`
	Url           string `json:"url"`
}

type GitChange struct {
	Item       GitItem `json:"item"`
	ChangeType string  `json:"changeType"`
}

type GitChanges struct {
	Changes []GitChange `json:"changes"`
}

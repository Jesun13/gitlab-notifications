package models

type MergeRequestEvent struct {
	ObjectKind string `json:"object_kind"`

	ObjectAttributes struct {
		Action       string `json:"action"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		SourceBranch string `json:"source_branch"`
		Description  string `json:"description"`
		TargetBranch string `json:"target_branch"`
	} `json:"object_attributes"`

	User struct {
		Name string `json:"name"`
	} `json:"user"`

	Project struct {
		Name string `json:"name"`
	} `json:"project"`

	Reviewers []struct {
		Name string `json:"name"`
	} `json:"reviewers"`
}

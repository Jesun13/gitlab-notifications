package models

type PipelineEvent struct {
	ObjectKind string `json:"object_kind"`

	ObjectAttributes struct {
		Status string `json:"status"`
		URL    string `json:"url"`
	} `json:"object_attributes"`

	User struct {
		Name string `json:"name"`
	} `json:"user"`

	Project struct {
		Name string `json:"name"`
	} `json:"project"`
}

package aws

import "time"

type Object struct {
	Key          string     `json:"key"`
	LastModified *time.Time `json:"last_modified"`
	Size         int64      `json:"size"`
}

type Content struct {
	Data         string     `json:"data"`
	LastModified *time.Time `json:"last_modified"`
}

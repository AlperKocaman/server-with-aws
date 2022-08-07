package aws

import "time"

type Object struct {
	Key          string     `json:"key"`
	LastModified *time.Time `json:"last_modified"`
	Size         int64      `json:"size_in_bytes"`
}

type Content struct {
	Data          string     `json:"data"`
	LastModified  *time.Time `json:"last_modified"`
	ContentLength int64      `json:"content_length"`
	ContentType   string     `json:"content_type"`
}

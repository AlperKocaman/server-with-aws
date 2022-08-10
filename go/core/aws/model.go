package aws

type Object struct {
	Key  string `json:"key"`
	Size int64  `json:"size_in_bytes"`
}

type Content struct {
	Data          string `json:"data"`
	ContentLength int64  `json:"content_length"`
	ContentType   string `json:"content_type"`
}

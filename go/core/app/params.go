package app

type SaveObjectParam struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}
type GetObjectParam struct {
	Key string `json:"key" form:"key"`
}

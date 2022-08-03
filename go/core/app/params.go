package app

type SaveObjectParam struct {
	Data interface{} `json:"data"`
}
type GetObjectParam struct {
	Key string `json:"key" form:"key"`
}

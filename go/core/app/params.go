package app

type SaveObjectParam struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}
type GetObjectParam struct {
	Key string `json:"key" form:"key"`
}

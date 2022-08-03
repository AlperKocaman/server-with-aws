package app

// TODO: Change all searializers/responses

type ListObjectsSerializer struct {
	Key string
}

type ListObjectsResponse struct {
	Key string `json:"key"`
}

func (l ListObjectsSerializer) Response() interface{} {
	return ListObjectsResponse(l)
}

type SaveObjectSerializer struct {
	Key string
}

type SaveObjectResponse struct {
	Key string `json:"key"`
}

func (l SaveObjectSerializer) Response() interface{} {
	return SaveObjectResponse(l)
}

type GetObjectSerializer struct {
	Key string
}

type GetObjectResponse struct {
	Key string `json:"key"`
}

func (l GetObjectSerializer) Response() interface{} {
	return GetObjectResponse(l)
}

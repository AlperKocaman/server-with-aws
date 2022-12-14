package app

import (
	"github.com/AlperKocaman/server-with-aws/core/aws"
)

// TODO: Change all searializers/responses

type ListObjectsSerializer struct {
	Objects []aws.Object
}

type ListObjectsResponse struct {
	FoundObjectNumber int          `json:"found_object_number"`
	Objects           []aws.Object `json:"objects"`
}

func (l ListObjectsSerializer) Response() interface{} {
	return ListObjectsResponse{
		Objects:           l.Objects,
		FoundObjectNumber: len(l.Objects),
	}
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
	Content aws.Content
}

type GetObjectResponse struct {
	Content aws.Content `json:"content"`
}

func (l GetObjectSerializer) Response() interface{} {
	return GetObjectResponse(l)
}

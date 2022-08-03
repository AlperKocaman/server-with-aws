package app

import "github.com/AlperKocaman/server-with-aws/core/response"

type Controller interface {
	ListObjects() (response.Responder, error)
	SaveObject(param SaveObjectParam) (response.Responder, error)
	GetObject(param GetObjectParam) (response.Responder, error)
}

type controller struct {
}

func NewController() Controller {
	return controller{}
}

func NewDefaultController() Controller {
	return NewController()
}

func (c controller) ListObjects() (response.Responder, error) {
	// GET /picus/list
	// return objects on the S3 bucket

	return ListObjectsSerializer{Key: "Deneme"}, nil
}

func (c controller) SaveObject(param SaveObjectParam) (response.Responder, error) {
	// POST /picus/put
	// save given JSON data to the S3 bucket

	return SaveObjectSerializer{Key: "Deneme"}, nil
}

func (c controller) GetObject(param GetObjectParam) (response.Responder, error) {
	// GET /picus/get/:key
	// return content of given key from S3

	return GetObjectSerializer{Key: "Deneme"}, nil

}

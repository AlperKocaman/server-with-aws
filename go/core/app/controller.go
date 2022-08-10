package app

import (
	"bytes"
	"github.com/AlperKocaman/server-with-aws/core/aws"
	"github.com/AlperKocaman/server-with-aws/core/response"
	"github.com/google/uuid"
	"log"
)

type Controller interface {
	ListObjects() (response.Responder, error)
	SaveObject(params SaveObjectParam) (response.Responder, error)
	GetObject(params GetObjectParam) (response.Responder, error)
}

type controller struct {
	S3ServiceClient *aws.S3ServiceClient
}

func NewController() Controller {
	c, err := aws.GetS3ServiceClient()
	if err != nil {
		log.Panicln("Cannot instantiate controller due to aws client")
	}
	return controller{
		S3ServiceClient: c,
	}
}

func NewDefaultController() Controller {
	return NewController()
}

func (c controller) ListObjects() (response.Responder, error) {
	// GET /picus/list
	// return objects on the S3 bucket

	log.Println("location: ListObjectsController")

	objects, err := c.S3ServiceClient.ListObjects()
	if err != nil {
		return nil, response.ServiceUnavailableError
	}

	return ListObjectsSerializer{Objects: objects}, nil
}

func (c controller) SaveObject(params SaveObjectParam) (response.Responder, error) {
	// POST /picus/put
	// save given JSON data to the S3 bucket

	log.Println("location: SaveObjectController")

	// if user sends a key with data, save the data with that key, otherwise generate a key
	key := params.Key
	if key == "" {
		key = uuid.New().String()
	}

	err := c.S3ServiceClient.Upload(key, bytes.NewReader([]byte(params.Data.(string))))
	if err != nil {
		return nil, response.ServiceUnavailableError
	}

	return SaveObjectSerializer{Key: key}, nil
}

func (c controller) GetObject(params GetObjectParam) (response.Responder, error) {
	// GET /picus/get/:key
	// return content of given key from S3

	log.Println("location: GetObjectController")

	if params.Key == "" {
		log.Println("Request without key")
		return nil, response.BadRequestError
	}

	object, err := c.S3ServiceClient.GetObject(params.Key)
	if err != nil {
		return nil, err
	}

	return GetObjectSerializer{Content: object}, nil

}

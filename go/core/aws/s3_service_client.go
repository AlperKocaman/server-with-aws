package aws

import (
	"github.com/AlperKocaman/server-with-aws/cmd/config"
	"github.com/AlperKocaman/server-with-aws/core/response"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
)

type S3ServiceClient struct {
	awsS3    *s3.S3
	uploader *s3manager.Uploader
}

func GetS3ServiceClient() (*S3ServiceClient, error) {
	awsSession, err := getAWSSession()
	if err != nil {
		return nil, err
	}

	return &S3ServiceClient{
		awsS3:    s3.New(awsSession),
		uploader: s3manager.NewUploader(awsSession),
	}, nil
}

func getAWSSession() (*session.Session, error) {
	log.Println("Getting session of aws...")

	sess, err := session.NewSession()

	if err != nil {
		log.Printf("Error while getting AWS S3 session, err: %v", err)
		return nil, err
	}

	return sess, nil
}

func (c S3ServiceClient) GetObject(key string) (Content, error) {
	res, err := c.awsS3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.Get().AWS.S3.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		if err.(awserr.Error).Code() == s3.ErrCodeNoSuchKey {
			return Content{}, response.NotFoundError
		}
		log.Printf("Error while getting object with key: %s, err: %v", key, err)
		return Content{}, response.ServiceUnavailableError
	}

	cLength := *res.ContentLength
	buf := make([]byte, cLength)
	read, err := res.Body.Read(buf)
	if read <= 0 || (err != nil && err != io.EOF) {
		log.Println("Error while reading content into buffer")
		return Content{}, response.InternalServerError
	}

	log.Printf("Successfully get object with key: %s", key)
	return Content{
		Data:          string(buf),
		ContentLength: cLength,
		ContentType:   aws.StringValue(res.ContentType),
	}, nil
}

func (c S3ServiceClient) ListObjects() ([]Object, error) {
	objects, err := c.awsS3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(config.Get().AWS.S3.Bucket),
	})
	if err != nil {
		log.Printf("Error while listing objects, err: %v", err)
		return nil, err
	}

	res := make([]Object, 0)
	for _, content := range objects.Contents {
		res = append(res, Object{
			Key:  aws.StringValue(content.Key),
			Size: aws.Int64Value(content.Size),
		})
	}

	log.Println("Successfully list objects")
	return res, nil
}

func (c S3ServiceClient) Upload(key string, body io.ReadSeeker) error {
	res, err := c.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Get().AWS.S3.Bucket),
		Key:    aws.String(key),
		Body:   body,
	})

	if err != nil {
		log.Printf("Error while uploading body with key: %s, err: %v", key, err)
	} else {
		log.Printf("Successfully uploaded body with key: %s to %s", key, res.Location)
	}

	return err
}

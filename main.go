import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Credentials struct {
	Key    string
	Secret string
}

type Client struct {
	ServerUrl  string
	Credential Credentials
	Client     *minio.Client
}

var s3Client *Client

func NewClient() (*Client, error) {
	s3Client, err := minio.New("minio.backend.localhost:8000", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "miniostorage", ""),
		Secure: false,
		Region: "us-east-1",
	})

	if err != nil {
		fmt.Printf("s3 client init error: %v", err)

		return nil, err
	}

	return &Client{
		ServerUrl: objectStorage.Endpoint,
		Credential: Credentials{
			Key:    objectStorage.AccessKey,
			Secret: objectStorage.SecretKey,
		},
		Client: client,
	}, nil
}

func main() {
	s3Client, err := NewClient()

	path := "/home/s/Downloads/go.tar.gz"

	if _, err := s3Client.StatObject(context.Background(), "cache", path, minio.StatObjectOptions{}); err != nil {
		log.Println(path + " does not exist")
	}
	if err := s3Client.FGetObject(context.Background(), "cache", path, "go.tar.gz", minio.GetObjectOptions{}); err != nil {
		log.Println(path + " does not exist")
	}

	if _, err := s3Client.FPutObject(context.Background(), "cache", path, "go.tar.gz", minio.PutObjectOptions{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully uploaded to ", path)
}

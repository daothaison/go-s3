import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	s3Client, err := minio.New("minio.backend.localhost:8000", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "miniostorage", ""),
		Secure: false,
		Region: "us-east-1",
	})
	if err != nil {
		log.Fatalln(err)
	}

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

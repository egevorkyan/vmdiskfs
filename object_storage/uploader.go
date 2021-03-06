package object_storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go/types"
	"log"
	"os"
)

func Uploader(useSSL bool, filename string) (string, int64, bool, error) {
	bucket := os.Getenv("AWS_DEST_BUCKET")
	extension := os.Getenv("OUTPUT_FORMAT")
	s3Client := ObjectStorageProvider(useSSL)

	targetFile, _ := TargetFileName(filename)

	object, err := os.Open("output/" + targetFile + "." + extension)
	if err != nil {
		log.Println(err)
	}
	defer object.Close()

	objectStat, err := object.Stat()
	if err != nil {
		log.Println(err)
	}
	n, err := s3Client.PutObject(context.Background(), bucket, objectStat.Name(), object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return targetFile, n.Size, false, err
	} else {
		return targetFile, n.Size, true, types.Error{}
	}
}

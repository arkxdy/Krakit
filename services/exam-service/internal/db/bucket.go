package db

import (
	"fmt"
	"log"

	"github.com/krakit/exam-service/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials" // Correct v7 import
)

type BucketClient struct {
	minio *minio.Client
}

// Pass the specific MinioConfig struct
func connectMinio(cfg config.MinioConfig) (*BucketClient, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize minio: %w", err)
	}

	log.Printf("Successfully initialized Minio client at %s", cfg.Endpoint)
	return &BucketClient{minio: minioClient}, nil
}

func (c *BucketClient) Close() {
	// Minio client doesn't have a 'Close' method like Mongo/SQL,
	// but keeping this for interface consistency if needed.
	fmt.Println("Bucket client session finished")
}

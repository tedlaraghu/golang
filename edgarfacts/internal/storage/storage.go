package storage

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

// upload Bytes
func UploadBytes(data []byte, bucket, path string) error {

	//Create buffer
	buffer := bytes.NewBuffer(data)

	//Create client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	defer client.Close()

	//create writer
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	writer := client.Bucket(bucket).Object(path).NewWriter(ctx)

	//copy data from buffer to Google cloud storage
	_, err = io.Copy(writer, buffer)
	if err != nil {
		return err
	}

	//close writer
	err = writer.Close()
	if err != nil {
		return nil
	}

	// Return Result
	return nil
}

package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	//access_key := "D8E4K4KAKX71QO2X79IG"
	//secret_key := "L9RiS7xMOVjCCcHfP4bPxGPGHd1MpiUSEHuODVYR"
	//end_point := "http://192.168.1.40:32561" //endpoint设置，不要动
	//
	//sess, err := session.NewSession(&aws.Config{
	//	Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
	//	Endpoint:         aws.String(end_point),
	//	Region:           aws.String("us-east-1"),
	//	DisableSSL:       aws.Bool(true),
	//	S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	//})

	endpoint := "192.168.1.40:32561"
	accessKeyID := "D8E4K4KAKX71QO2X79IG"
	secretAccessKey := "L9RiS7xMOVjCCcHfP4bPxGPGHd1MpiUSEHuODVYR"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	// Make a new bucket called mymusic.
	bucketName := "video"

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	f, err := os.Open("test.jpg")
	fmt.Println(err)
	contentType := "image/jpeg"

	stte, err := f.Stat()
	info, err := minioClient.PutObject(ctx, bucketName, "111111", f, stte.Size(), minio.PutObjectOptions{
		ContentType: contentType,
	})
	fmt.Println(info)

	//a, err := minioClient.GetObject(ctx, bucketName, "111111", minio.GetObjectOptions{})
	//info, err := a.Stat()
	//fmt.Println(info.ContentType)
	//
	//data, err := ioutil.ReadAll(a)
	//
	//os.WriteFile("a.jpg", data, 0775)

	//// Upload the zip file with FPutObject
	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(info)
	//fmt.Println(err)

	////svc := s3.New(sess)
	////result, err := svc.ListBuckets(nil)
	////if err != nil {
	////	panic("Unable to list buckets")
	////}
	////
	////fmt.Println("Buckets:")
	////
	////for _, b := range result.Buckets {
	////	fmt.Printf("* %s created on %s\n",
	////		aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	////}
	//
	//uploader := s3manager.NewUploader(sess)
	//
	//_, err = uploader.Upload(&s3manager.UploadInput{
	//	Bucket: aws.String("video"),
	//	Key: aws.String("11111"),
	//	Body: bytes.NewBuffer([]byte("hello, word")),
	//})
	//if err != nil {
	//	// Print the error and exit.
	//	fmt.Println(err)
	//}

}

func TestUploadBigFile(t *testing.T) {
	endpoint := "192.168.1.40:32561"
	accessKeyID := "D8E4K4KAKX71QO2X79IG"
	secretAccessKey := "L9RiS7xMOVjCCcHfP4bPxGPGHd1MpiUSEHuODVYR"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

	link, err := minioClient.PresignedGetObject(context.Background(), "other", "62d372a4ec2e74d356131458", time.Hour, reqParams)
	fmt.Println(link)
	fmt.Println(err)

}

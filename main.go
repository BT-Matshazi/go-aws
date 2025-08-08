package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/joho/godotenv"
)

// main function
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// S3 bucket names
	bucketName := "bekithemba-go-bucket-learn"

	// Get the AWS region from environment variable loaded from .env
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatalf("AWS_REGION environment variable not set in .env file.")
	}
	fmt.Printf("Using AWS Region from .env: %s\n", region)

	// from the environment variables loaded by godotenv.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	// Create an S3 service client.
	s3Client := s3.NewFromConfig(cfg)

	// Call the function to create the S3 bucket.
	err = createS3Bucket(context.TODO(), s3Client, bucketName, region)
	if err != nil {
		log.Fatalf("failed to create S3 bucket %s: %v", bucketName, err)
	}

	fmt.Printf("Successfully created S3 bucket: %s in region %s\n", bucketName, region)
}

// createS3Bucket creates an S3 bucket with the specified name and region.
func createS3Bucket(ctx context.Context, client *s3.Client, bucketName, region string) error {
	// Prepare the input for the CreateBucket API call.
	// For regions other than us-east-1, a CreateBucketConfiguration is required
	// to specify the location constraint (the region).
	createBucketInput := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName), // The name of the bucket to create
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region), // Specify the region
		},
	}

	// Perform the CreateBucket API call.
	_, err := client.CreateBucket(ctx, createBucketInput)
	if err != nil {
		// Log the error and return it.
		return fmt.Errorf("error creating bucket %s: %w", bucketName, err)
	}

	return nil // Return nil if the bucket was created successfully.
}

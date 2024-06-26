package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	accessKeyID := "access-key"
	secretAccessKey := "secretAccessKey"
	region := "region"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := secretsmanager.NewFromConfig(cfg)

	secretValue, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("MySecret"),
	})
	if err != nil {
		log.Fatalf("Unable to retrieve secret: %v", err)
	}

	if secretValue.SecretString != nil {
		fmt.Println("Secret value:", *secretValue.SecretString)
	} else {
		fmt.Println("No secret string found.")
	}
}

# ecommerceUser
This repository contains a lambda function to create users in AWS RDS - MySQL database when a new user is created in AWS Cognito.

## Configuration
Follow all of these steps to configure the repository:

To create a Go module: 
- **go mod init github.com/luisrojas17/ecommerceUser

To get AWS SDK for this lambda:
- **go get github.com/aws/aws-lambda-go/lambda
- **go get github.com/aws/aws-lambda-go/events
- **go get github.com/aws/aws-sdk-go-v2/aws
- **go get github.com/aws/aws-sdk-go-v2/config
- **github.com/aws/aws-sdk-go-v2/service/secretsmanager
- **go get github.com/go-sql-driver/mysql

## Compilation
This project has to be compiled in linux format since AWS lambda runtime for Go only 
recognize lambda Linux executable. so you have to set next environment variables:

- set GOOS=linux
- set GOARCH=amd64

## Dependencies
This project has next AWS dependencies: 

- You have to create user pool in AWS Cognito.
- You have to create a database instance in AWS RDS. In this case, the instance will have to be MySQL.
- You have to create an schema with its tables for database. Load the script:
    ./db/ecommerceSchema.sql
- You have to create a secret in AWS Secrete Manager to store the database credentials.
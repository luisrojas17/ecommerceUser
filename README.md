# ecommerceUser
This repository contains a lambda function to create users in AWS RDS - MySQL database when a new user is created in AWS Cognito.

## Configuration project
Follow all of these steps to configure the repository:

To create a Go module: 
- **go mod init github.com/luisrojas17/ecommerceUser

To get AWS SDK for this lambda:
- **go get github.com/aws/aws-lambda-go/lambda
- **go get github.com/aws/aws-lambda-go/events
- **go get github.com/aws/aws-sdk-go-v2/aws
- **go get github.com/aws/aws-sdk-go-v2/config
- **go get github.com/aws/aws-sdk-go-v2/service/secretsmanager
- **go get github.com/go-sql-driver/mysql

## Dependencies
This project has next AWS dependencies: 

- You have to create user pool in AWS Cognito.
- You have to create a database instance in AWS RDS. In this case, the instance will have to be MySQL.
- You have to create an schema with its tables for database. Load the script:
    ./db/ecommerceSchema.sql
- You have to create a secret in AWS Secrete Manager to store the database credentials.
- Create and config Lambda function. This lambda function will insert a new users created in Cognito once the user confirm the email confirmation sent by Cognito.

## Compilation
This project has to be compiled in linux format since AWS lambda runtime for Go only 
recognize lambda Linux executable. so you have to set next environment variables:

- set GOOS=linux
- set GOARCH=amd64

See the scripts "compile.bat" or "compile.sh"

## AWS resources configuration

### Create user pool in Cognito


    Add new scopes to user pool application
        Go to user pool:            auth_user_group
        Recomendations section:     auth_user_app
        Go to Login Pages Tab and click on "Edit" button 
        Edit managed login pages configuration 
            OpenID Connect scopes: aws.cognito.signin.admin
            Click on "Save changes" button


## Create databse in RDS

### Create Lambda

Add trigger 

When you are going to load the code zip file for your lambda function you will have to add next  environment varirables:

- SecretName=ecommerce_secret


In next section it will be described the process to configure an API gateway to get and request objects from and to S3. Also, you can check next documentation:

https://repost.aws/knowledge-center/api-gateway-upload-image-s3

### Create bucket

    Go to S3 service
    Create bucket
    Select the region:  us-east-1 (is cheper)
    Bucket type:        General purpose
    Bucker name:        ecommerce-app-jlrg
    Object ownership:   ACLs disabled (Recommended)
    Block all public access in order to made the bucket public: Uncheck
    I acknowledge that the current settings might result in this bucket and the objects within becoming public:     Check.
    Bucket versioning:  Disable
    Default encription: Server-side encryption with Amazon S3 managed keys (SSE-S3)
        Bucket key:     Enabled
    Advanced settings
        Object lock:    Disabled.

    Add security policy to get any file allocated in the bucket
        Go to Bucket ecommerce-app-jlrg
            Go to Permissions tab
                Bucket policy Edit
                Add content in s3Policy.json policy
                Click on "Save changes" button

### Create role
    
    Go to IAM service
    Go to roles section 
    Create an IAM role "ecommerce-api-gateway-to-s3"
        Service type = AWS Service
        Service or use case = API Gateway
        Add any description.

        Note: By default, the new role only add permissions to push logs in CloudWatch. So you have to add permissions according to your necessity.

    Go to "ecommerce-api-gateway-to-s3"
    Go to Permissions section  
        Add permissions
            Add policies -> Create inline policy
                Select a service: S3 
                Actions allowed: PutOB -> PutObject
                Actions allowed: GetOB -> GetObject 
                Select checkbox: 
                    Anyone in this account if you wish read and write for any buckect created. If not, you have to specify the Bucket's ARN name. For example:
                        Resource bucket name:   ecommerce-app-jlrg
                        Resource object name:   * 
                        ARN name:               arn:aws:s3:::ecommerce-app-jlrg
                
                Policy name: ecommerce-s3-policy

                Select a service: Cognito identity
                Actions allowed: 
                    List (enumeración), 
                    Read (Lectura), 
                    Write (Escritura), 
                    Label (Tagging)  
                Select checkbox: Anyone in this account

                Policy name: ecommerce-cognito-policy

### Create an API Gateway

    Go to API Gateway service
    Create API
    Choose an API type: REST API
    API details:        New API 
    API name:           ecommerce-api-for-s3
    Description:        This API manage all operations to communicate with our ecommerce-app-jlrg 
                        S3 bucket.
    API endpoint type:  Regional
    Clicl on "Create API" button

    In Resources section add resource
    Click on "Create Resource"
        Resource details
            Resource path:                          /
            Resource name:                          bucket_name
            CORS (Cross Origin Resource Sharing):   Enabled

        Resource details. Select bucket_name created and create other resource
            Resource path:                          /bucket_name
            Resource name:                          file_name
            CORS (Cross Origin Resource Sharing):   Enabled

        Select file_name and create a method
            Method details
                Method type:        PUT
                Integration type:   AWS Service
                AWS Region:         us-east-1 (Virginia)
                AWS Service:        Simple Storage Service (S3)
                AWS subdomain:      
                HTTP Method:        PUT
                Action type:        Use path override
                Path override:      {bucket}/{key}
                Execution role:     arn:aws:iam::840004571262:role/ecommerce-api-gateway-to-s3
                                    Note: This the role that you create before.
                Credential cache:   Do not add caller credentials  to cache key
                Content handling:   Passthrough
                Integration timeout:29000    

    Go to ecommerce-api-for-s3 Authorizers
        Create Authorizer
            Authorizer details
                Authorizer name:    ecommer-authorizer
                Authorizer type:    Cognito
                Cognito user pool:  us-east-1 auth_user_group
                                    Note: This is the user pool name created before.
                Token source:       Authorization
                                    Note: This is the header in request "Authorization".
                Click on "Create authorizer" button

    Go to ecommerce-api-for-s3 resource section
    Select PUT Method in left panel
    Got to Method request settings and click on "Edit" button
        Authorizator:           ecommer-authorizer
        Authorization scopes:   openid
        Request validator:      None
        API key required:       Unchecked


    Go to ecommerce-api-for-s3 API Settings
    Binary media types
        Manage media types
        Click on "Add binary media type" button
            image/jpeg
            image/png
        Click on "Save changes" button

    Go to ecommerce-api-for-s3 Resource Section 
    Select PUT method 
    Click on "API actions"

    Go to ecommerce-api-for-s3 Resources
    Select PUT Method resource

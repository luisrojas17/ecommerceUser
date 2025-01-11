package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/luisrojas17/ecommerceuser/aws"
	"github.com/luisrojas17/ecommerceuser/db"
	"github.com/luisrojas17/ecommerceuser/models"
)

func main() {

	lambda.Start(Execute)

}

func Execute(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	// Init AWS configuration for lambda
	aws.Init()

	// If not exist secretName lambda function has to finalize
	if !checkEnvironmentParameters() {
		var msg string = "It was an error. The secret name was not specified."
		fmt.Println(msg)
		err := errors.New(msg)

		return event, err
	}

	// Get user data from Cognito
	var user models.User
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			user.Email = att
			fmt.Println("Email gotten: ", user.Email)
		case "sub":
			user.Uuid = att
			fmt.Println("UserId gotten: ", user.Uuid)
		}
	}

	/*err := db.GetSecret()
	if err != nil {
		fmt.Println("It was error to connect to database.", err.Error())
		return event, err
	}*/

	// Insert user to database
	err := db.Create(user)

	return event, err

}

/*
This function check if the secret name was prvided into environment.
*/
func checkEnvironmentParameters() bool {
	var exist bool
	_, exist = os.LookupEnv("SecretName")

	return exist
}

package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/luisrojas17/ecommerceuser/helper"
	"github.com/luisrojas17/ecommerceuser/models"
)

func Create(user models.User) error {

	fmt.Println("Starting to insert user in database...")

	err := Connect()

	if err != nil {
		return err
	}

	defer Close()

	statement := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + user.Email + "', '" + user.Uuid + "', '" + helper.GetDate() + "')"
	fmt.Println(statement)

	_, err = Connection.Exec(statement)
	if err != nil {
		fmt.Println("It was an error to insert user.", err.Error())
		return err
	}

	fmt.Println("User insert was made it database.")

	return nil
}

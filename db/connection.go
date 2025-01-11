package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/luisrojas17/ecommerceuser/aws"
	"github.com/luisrojas17/ecommerceuser/models"
)

//var secretModel models.Secret
//var err error

var Connection *sql.DB

/*func GetSecret() error {

	secretModel, err = aws.GetSecret(os.Getenv("SecretName"))

	return err
}*/

func Connect() error {
	secretModel, err := aws.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		fmt.Println("It was error to connect to database. It was not possible to get and parse the secret name.",
			err.Error())
		return err
	}

	Connection, err := sql.Open("mysql", buildConnectionString(secretModel))

	if err != nil {
		fmt.Println("It was an error when the application tried to connect to database.", err.Error())
		return err
	}

	err = Connection.Ping()
	if err != nil {
		fmt.Println("It was not possible to ping to database.", err.Error())
		return err
	}

	fmt.Println("Connection to database was successfull.")

	return nil
}

func Close() {
	fmt.Println("Closing connection...")

	Connection.Close()

	fmt.Println("Connection was closed.")
}

func buildConnectionString(secret models.Secret) string {
	var user, pass, dbEndpoint, dbName string

	user = secret.Username
	pass = secret.Password
	dbEndpoint = secret.Host
	dbName = "ecommerce_db"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPassword=true", user, pass, dbEndpoint, dbName)

	fmt.Println("Connection string is: ", connectionString)

	return connectionString
}

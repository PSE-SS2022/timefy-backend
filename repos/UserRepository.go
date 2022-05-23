package repos

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions *options.ClientOptions
var dbclient *mongo.Client

func InitDB() error {
	// Define Address of Database
	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	// Try to connect to Database, save error if one is thrown
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// If there was an error connecting to the DB (DB not running, wrong URI, ...) return the error
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	dbclient = client
	return nil
}

func getDb(dbname string) *mongo.Database {
	return dbclient.Database(dbname)
}

func DbExists(id string) bool {
	return getDb(id) != nil
}

// returns date in following format: 23-05-2022
func GetCurrentDate() string {
	currentTime := time.Now()
	month := fmt.Sprintf("%02d", int(currentTime.Month()))
	return strconv.Itoa(currentTime.Day()) + "-" + month + "-" + strconv.Itoa(currentTime.Year())
}

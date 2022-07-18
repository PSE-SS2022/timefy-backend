package database

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseMgrInstance databaseMgr = newDatabaseMgr()

const databaseName string = "timefy"
const host string = "localhost"
const port int = 27017

type databaseMgr struct {
	clientOptions *options.ClientOptions
	dbclient      *mongo.Client
}

func newDatabaseMgr() databaseMgr {
	databaseMgr := databaseMgr{}
	databaseMgr.initDB()

	return databaseMgr
}

func (database *databaseMgr) initDB() error {
	// Define Address of Database
	var portString string = strconv.Itoa(port)
	var uri string = "mongodb://" + host + portString
	database.clientOptions = options.Client().ApplyURI(uri)
	// Try to connect to Database, save error if one is thrown
	client, err := mongo.Connect(context.TODO(), database.clientOptions)
	// If there was an error connecting to the DB (DB not running, wrong URI, ...) return the error
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	database.dbclient = client
	return nil
}

func (database *databaseMgr) getDb() *mongo.Database {
	return database.dbclient.Database(databaseName)
}

func (database *databaseMgr) getCollection(collection string) *mongo.Collection {
	return database.getDb().Collection(collection)
}

// returns date in following format: 23-05-2022
/*func GetCurrentDate() string {
	currentTime := time.Now()
	month := fmt.Sprintf("%02d", int(currentTime.Month()))
	return strconv.Itoa(currentTime.Day()) + "-" + month + "-" + strconv.Itoa(currentTime.Year())
}*/

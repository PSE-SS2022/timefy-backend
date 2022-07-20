package database

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseMgrInstance databaseMgrBase

const dbName string = "timefy"
const host string = "localhost"
const port int = 27017

type DatabaseTypes int

const (
	DatabaseTypeNormal DatabaseTypes = iota
	DatabaseTypeTestingInMemory
	DatabaseTypeTestingNormal
)

type databaseMgrDefault struct {
	clientOptions *options.ClientOptions
	dbclient      *mongo.Client
	databaseName  string
}

func newDatabaseMgrDefault(databaseName string) databaseMgrDefault {
	databaseMgr := databaseMgrDefault{databaseName: databaseName}
	databaseMgr.InitDb()
	databaseMgr.CreateCollections()
	return databaseMgr
}

func newDatabaseMgrTesting() databaseMgrInMemory {
	databaseMgr := databaseMgrInMemory{}
	databaseMgr.InitDb()
	databaseMgr.CreateCollections()
	return databaseMgr
}

type databaseMgrInMemory struct {
	mongoURI     string
	databaseName string
	database     *mongo.Database
}

type databaseMgrBase interface {
	GetDatabase() *mongo.Database
	GetCollection(collectionName string) *mongo.Collection
	CreateCollections()
}

func (databaseMgr databaseMgrDefault) GetDatabase() *mongo.Database {
	return databaseMgr.dbclient.Database(databaseMgr.databaseName)
}

func (databaseMgr databaseMgrDefault) GetCollection(collectionName string) *mongo.Collection {
	return databaseMgr.GetDatabase().Collection(collectionName)
}

func (databaseMgr *databaseMgrDefault) InitDb() error {
	// Define Address of Database
	var portString string = strconv.Itoa(port)
	var uri string = "mongodb://" + host + portString
	databaseMgr.clientOptions = options.Client().ApplyURI(uri)
	// Try to connect to Database, save error if one is thrown
	client, err := mongo.Connect(context.TODO(), databaseMgr.clientOptions)
	// If there was an error connecting to the DB (DB not running, wrong URI, ...) return the error
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	databaseMgr.dbclient = client
	return nil
}

func (databaseMgr databaseMgrDefault) CreateCollections() {
}

func (databaseMgr databaseMgrInMemory) CreateCollections() {
}

func (databaseMgr databaseMgrInMemory) GetDatabase() *mongo.Database {
	return databaseMgr.database
}

func (databaseMgr databaseMgrInMemory) GetCollection(collectionName string) *mongo.Collection {
	return databaseMgr.GetDatabase().Collection(collectionName)
}

func (databaseMgr *databaseMgrInMemory) InitDb() error {
	mongoServer, err := strikememongo.StartWithOptions(&strikememongo.Options{MongoVersion: "4.2.0", ShouldUseReplica: true})
	if err != nil {
		return err
	}
	databaseMgr.mongoURI = mongoServer.URIWithRandomDB()
	splitedDatabaseName := strings.Split(databaseMgr.mongoURI, "/")
	databaseMgr.databaseName = splitedDatabaseName[len(splitedDatabaseName)-1]

	dbClient, err := databaseMgr.initInMemoryDB()
	if err != nil {
		//log.Fatal("error connecting to database", err)
		return nil
	}

	databaseMgr.database = dbClient.Database(databaseMgr.databaseName)

	return nil
}

func (databaseMgr *databaseMgrInMemory) initInMemoryDB() (client *mongo.Client, err error) {
	uri := fmt.Sprintf("%s%s", databaseMgr.mongoURI, "?retryWrites=false")
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	return
}

func SetupDatabase(databseType DatabaseTypes) {
	switch databseType {
	case DatabaseTypeNormal:
		databaseMgrInstance = newDatabaseMgrDefault(dbName)
	case DatabaseTypeTestingInMemory:
		databaseMgrInstance = newDatabaseMgrTesting()
	case DatabaseTypeTestingNormal:
		databaseMgrInstance = newDatabaseMgrDefault(dbName + "_testing")
	}
}

// returns date in following format: 23-05-2022
/*func GetCurrentDate() string {
	currentTime := time.Now()
	month := fmt.Sprintf("%02d", int(currentTime.Month()))
	return strconv.Itoa(currentTime.Day()) + "-" + month + "-" + strconv.Itoa(currentTime.Year())
}*/

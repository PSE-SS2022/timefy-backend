package repos

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"timefy-backend/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetReports() []models.ExtendedReport {
	return demoReports
}

// TODO: need to add something in front of id as id may not start with an number --> jquery error
var demoReports = []models.ExtendedReport{
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Abdullah#123", "Abdullah", "Yildirim", "21.05.2022", "1", "Mittagessen", "Hier treffen wir uns zum Mittagessen in der Mensa"},
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Talip#124", "Talip", "Göksu", "19.05.2022", "2", "Fußball", "Hi, wer hat Lust auf Fußball"},
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Barrack#125", "Barrack", "Obama", "10.05.2022", "3", "Murriicaa", "Murriicaaaaaa"},
}

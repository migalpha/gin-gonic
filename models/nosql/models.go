package nosql

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var ctx context.Context

func init() {
	var (
		err        error
		connString string
	)

	connString = os.Getenv("CONN_STRING")
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("gin-gonic")
}

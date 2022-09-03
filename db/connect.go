package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VJ-Vijay77/mongodb/config"
	"github.com/VJ-Vijay77/mongodb/errpkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var DbNAME = os.Getenv("DB_NAME")
// var CollectionNAME = os.Getenv("COLLECTION_NAME")

func Connect(uri string) (*mongo.Client, context.Context, error) {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	// fmt.Println("Mongodb Connection successful..")

	return client, ctx, err
}

func Ping(client *mongo.Client, ctx context.Context) {
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Connection to MongoDB successful...")
	}

}

func GetMongodbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	Uri := os.Getenv("MONGO_URI")
	client, ctx, err := Connect(Uri)
	_ = ctx
	errpkg.LogPrint(err)

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func NewDbData() *config.DbVariable {
	
	DbName := os.Getenv("DB_NAME")
	CollectionNAME := os.Getenv("COLLECTION_NAME")
	U := &config.DbVariable{
		DbName:         DbName,
		CollectionName: CollectionNAME,
	}
	return U
}

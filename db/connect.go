package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/VJ-Vijay77/mongodb/config"
	"github.com/VJ-Vijay77/mongodb/errpkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbNAME = os.Getenv("DB_NAME")
var CollectionNAME = os.Getenv("COLLECTION_NAME")


func Connect(uri string) (*mongo.Client, error) {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	fmt.Println("Mongodb Connection successful..")

	return client, err
}

// func Ping(client *mongo.Client, ctx context.Context) {
// 	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
// 		// errpkg.LogPrint(err)
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Connection is Successfull...")
// 	}

// }


func GetMongodbCollection(DbName string,CollectionName string) (*mongo.Collection,error) {

	Uri := os.Getenv("MONGO_URI")
	client,err := Connect(Uri)
	errpkg.LogPrint(err)

	collection := client.Database(DbName).Collection(CollectionName)
	
	return collection,nil
}


func NewDbData() *config.DbVariable {
	

	U := &config.DbVariable{
		DbName: DbNAME,
		CollectionName: CollectionNAME,
	}
	return U
}
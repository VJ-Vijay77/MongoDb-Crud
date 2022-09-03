package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)


//? This method closes mongoDB connection and cancel context.
func closeDisconnect(client *mongo.Client,c context.Context,cancel context.CancelFunc) {
	
	 //? CancelFunc to cancel to context
	defer cancel()

	
    // client provides a method to close
    // a mongoDB connection.
	defer func() {
		if err := client.Disconnect(c); err != nil {
			panic(err)
		}
	}()
}
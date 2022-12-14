package controllers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/VJ-Vijay77/mongodb/db"
	"github.com/VJ-Vijay77/mongodb/errpkg"
	"github.com/VJ-Vijay77/mongodb/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var D *config.DbVariable

func GetPerson(c *fiber.Ctx) error {
	k := db.NewDbData()
	if k.CollectionName == "" || k.DbName == "" {
		log.Fatalln("couldnt get env data")
	}
	 collection,err := db.GetMongodbCollection(k.DbName,k.CollectionName)
	 if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objId,_ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id":objId}
	}

	var results []bson.M
	cur,err := collection.Find(context.Background(),filter)
	errpkg.StatusFiveHundred(c,err)
	
	defer cur.Close(context.Background())

	cur.All(context.Background(),&results)

	if results == nil {
		c.SendStatus(404)
		return nil
	}

	json,_ := json.Marshal(results)
	c.Send(json)


	log.Println("Got One Record Successfully...")

	return nil
}

func CreatePerson(c *fiber.Ctx) error {

	k := db.NewDbData()
	if k.CollectionName == "" || k.DbName == "" {
		log.Fatalln("couldnt get env data")
	}
	collection, err := db.GetMongodbCollection(k.DbName, k.CollectionName)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	var person models.Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	errpkg.StatusFiveHundred(c, err)

	response, _ := json.Marshal(res)
	c.Send(response)

	log.Println("One Record Inserted Successfully...")
	return nil
}

func UpdatePerson(c *fiber.Ctx) error {
	k := db.NewDbData()
	if k.CollectionName == "" || k.DbName == "" {
		log.Fatalln("couldnt get env data")
	}
	collection, err := db.GetMongodbCollection(k.DbName, k.CollectionName)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	var person models.Person
	json.Unmarshal([]byte(c.Body()),&person)

	update := bson.M{
		"$set":person,
	}
	objID,_ := primitive.ObjectIDFromHex(c.Params("id"))
	res,err := collection.UpdateOne(context.Background(),bson.M{"_id":objID},update)
	errpkg.StatusFiveHundred(c,err)

	response,_ := json.Marshal(res)
	c.Send(response)

	log.Println("Updated One Record Successfully...")
	
	return nil
}

func DeletePerson(c *fiber.Ctx) error {
	k := db.NewDbData()
	if k.CollectionName == "" || k.DbName == "" {
		log.Fatalln("couldnt get env data")
	}
	collection, err := db.GetMongodbCollection(k.DbName, k.CollectionName)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	objID,_ := primitive.ObjectIDFromHex(c.Params("id"))
	res,err := collection.DeleteOne(context.Background(),bson.M{"_id":objID})
	errpkg.StatusFiveHundred(c,err)

	json,_ := json.Marshal(res)
	c.Send(json)
	
	log.Println("Deleted One Record Successfully...")

	return nil
}

package controllers

import (
	"context"
	"encoding/json"

	"github.com/VJ-Vijay77/mongodb/db"
	"github.com/VJ-Vijay77/mongodb/errpkg"
	"github.com/VJ-Vijay77/mongodb/models"
	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	return nil
}

func CreatePerson(c *fiber.Ctx) error{

	k := db.NewDbData()

	collection, err := db.GetMongodbCollection(k.DbName, k.CollectionName)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))	 
	}

	var person models.Person
	json.Unmarshal([]byte(c.Body()),&person)

	res,err := collection.InsertOne(context.Background(),person)
	errpkg.StatusFiveHundred(c,err)

	response,_ := json.Marshal(res)
	c.Send(response)

 return nil
}

func UpdatePerson(c *fiber.Ctx) error {
	return nil
}

func DeletePerson(c *fiber.Ctx) error {
	return nil
}

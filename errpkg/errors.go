package errpkg

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Panic(err error) {
	if err != nil {
		Panic(err)
	}

}

func Fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func LogPrint(err error) {
	if err != nil {
		log.Println(err)
	}

}

func StatusFiveHundred(c *fiber.Ctx, err error) {
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))	 
		return
	}
}

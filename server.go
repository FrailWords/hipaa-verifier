package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hipaa-verifier/hipaa"
	"log"
	"os"
	"strconv"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	app.Post("/protect", func(c *fiber.Ctx) error {
		source := c.Query("source")
		destination := c.Query("destination")
		rule := serviceRule(source, destination)
		if rule == nil {
			return c.JSON(nil)
		}
		protectedData, err := hipaa.ProtectData(c, rule)
		if err != nil {
			log.Printf("Error protecting incoming data %s:", err.Error())
			return c.JSON(nil)
		}
		return c.JSON(protectedData)
	})
	app.Post("/verify", func(c *fiber.Ctx) error {
		source := c.Query("source")
		destination := c.Query("destination")
		rule := serviceRule(source, destination)
		if rule == nil {
			return c.JSON(nil)
		}
		verified, err := hipaa.VerifyData(c, rule)
		if err != nil {
			log.Printf("Error verifying incoming data %s:", err.Error())
			return c.JSON(nil)
		}
		return c.SendString(strconv.FormatBool(verified))
	})
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func serviceRule(source string, destination string) *hipaa.InterServiceRule {
	var rule *hipaa.InterServiceRule
	for _, serviceRule := range hipaa.RegisteredServices {
		if serviceRule.Source == source && serviceRule.Destination == destination {
			rule = &serviceRule
			break
		}
	}
	return rule
}

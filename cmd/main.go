package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nubrid/go-api-demo/internal/handlers"
)

func main() {
	// const express = require("express")
	// const app = express()
	// app.use(express.json())
	app := fiber.New()

	// app.use((req, res, next) => {
	// 	console.log("Sample middleware")
	//  next()
	// })
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println(("Sample middleware"))

		return c.Next()
	})

	// app.get('/', (req, res) => {
	// 	res.send("OK")
	// })
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/products", handlers.GetAllProducts)
	app.Post("/api/products", handlers.CreateProduct)

	// console.log("Hello world")
	fmt.Println("Hello world")

	// try { app.listen(3000) } catch (err) { console.log(err); process.exit(1) }
	log.Fatal(app.Listen(":3000"))
}

package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
)

func homeHandler(c fiber.Ctx) error {
	return c.SendString("homeHandler")
}

func projects(c fiber.Ctx) error {
	return c.SendString("projects")
}

func project(c fiber.Ctx) error {
	if c.Params("name") == "" {
		return c.SendString("project")
	}

	return c.SendString("project " + c.Params("name"))
}

func articles(c fiber.Ctx) error {
	return c.SendString("articles")
}

func article(c fiber.Ctx) error {
	if c.Params("name") == "" {
		return c.SendString("article")
	}

	return c.SendString("article " + c.Params("name"))
}

func contacts(c fiber.Ctx) error {
	return c.SendString("contacts")
}

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//PATH_VAULT := os.Getenv("E:\\Sync\\obsidian-dextreme")

	app.Get("/", homeHandler)
	app.Get("/projects", projects)
	app.Get("/project/:name", project)
	app.Get("/articles", articles)
	app.Get("/article/:name", article)
	app.Get("/contacts", contacts)

	log.Fatal(app.Listen(":8080"))
}

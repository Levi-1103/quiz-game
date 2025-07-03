package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

func main() {
	app := fiber.New()

	setupDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/quizzes", getQuizzes)

	app.Handler()

	log.Fatal(app.Listen(":3000"))
}

func setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://root:example@localhost:27017/?authSource=admin"))
	if err != nil {
		panic(err)
	}

	quizCollection = client.Database("quiz").Collection("quizzes")

}

func getQuizzes(c *fiber.Ctx) error {
	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	quizzes := []map[string]any{}
	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		return err
	}

	return c.JSON(quizzes)
}

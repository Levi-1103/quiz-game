package internal

import (
	"backend/internal/collection"
	"backend/internal/controller"
	"backend/internal/service"
	"context"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

type App struct {
	server      *fiber.App
	database    *mongo.Database
	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupDB()
	a.setupServices()
	a.setupServer()

	log.Fatal(a.server.Listen(":3000"))
}

func (a *App) setupServer() {
	app := fiber.New()

	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)

	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.Ws(a.netService)
	app.Get("/ws/", websocket.New(wsController.Ws))

	app.Handler()

	log.Fatal(app.Listen(":3000"))
	a.server = app
}

func (a *App) setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://root:example@localhost:27017/?authSource=admin"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("quiz")

}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)
}

package service

import (
	"fmt"
	"strings"

	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService
	host        *websocket.Conn
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case "host":
		{
			fmt.Println("host quiz:", argument)
			c.host = con
			break
		}
	case "join":
		{
			fmt.Println("join code:", argument)
			c.host.WriteMessage(websocket.TextMessage, []byte("A player has joined"))
			break
		}
	}

}

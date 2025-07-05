package service

import (
	"backend/internal/entity"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NetService struct {
	quizService *QuizService
	games       []*Game
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
		games:       []*Game{},
	}
}

type Packet struct {
	Type string          `json:"code"`
	Data json.RawMessage `json:"data"`
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HostGamePacket struct {
	QuizId string `json:"quizId"`
}

type QuestionShowPacket struct {
	Question entity.QuizQuestion `json:"question"`
}

type ChangeGameStatePacket struct {
	State GameState `json:"state"`
}

const (
	PacketConnect      = "connect"
	PacketHost         = "host"
	PacketQuestionShow = "question"
	ChangeGameState    = "state"
)

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {

	var base Packet

	err := json.Unmarshal(msg, &base)
	if err != nil {
		fmt.Println("Invalid base packet", err)
		return
	}

	switch base.Type {
	case "connect":
		var connectPacket ConnectPacket
		var err = json.Unmarshal(base.Data, &connectPacket)
		if err != nil {
			fmt.Println("Problem with packet")
			return
		}
		game := c.getGameByCode(connectPacket.Code)
		if game == nil {
			return
		}
		game.OnPlayerJoin(connectPacket.Name, con)

	case "host":

		var hostGamePacket HostGamePacket
		var err = json.Unmarshal(base.Data, &hostGamePacket)
		if err != nil {
			fmt.Println("Problem with packet")
			return
		}

		quizId, err := primitive.ObjectIDFromHex(hostGamePacket.QuizId)
		if err != nil {
			fmt.Println(err)
			return
		}

		quiz, err := c.quizService.quizCollection.GetQuizById(quizId)
		if err != nil {
			fmt.Println(err)
			return
		}

		if quiz == nil {
			return
		}

		newGame := newGame(*quiz, con, c)

		fmt.Println("User wants to host quiz", newGame.Code)

		c.games = append(c.games, &newGame)
		c.SendPacket(con, ChangeGameState, ChangeGameStatePacket{
			State: LobbyState,
		})

		// go func() {
		// 	time.Sleep(time.Second * 2)
		// 	c.SendPacket(con, PacketQuestionShow, QuestionShowPacket{
		// 		Question: entity.QuizQuestion{
		// 			Name: "What is 2 + 2?",
		// 			Choices: []entity.QuizChoice{
		// 				{
		// 					Name: "4",
		// 				},
		// 				{
		// 					Name: "9",
		// 				},
		// 				{
		// 					Name: "8",
		// 				},
		// 				{
		// 					Name: "7",
		// 				},
		// 			},
		// 		},
		// 	})
		// }()

	default:
		fmt.Println("Unknown packet type:", base.Type)

	}
}

func (c *NetService) getGameByCode(code string) *Game {
	for _, game := range c.games {
		if game.Code == code {
			return game
		}
	}
	return nil
}

func (c *NetService) SendPacket(connection *websocket.Conn, code string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	packet := Packet{
		Type: code,
		Data: data,
	}

	msg, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	return connection.WriteMessage(websocket.TextMessage, msg)
}

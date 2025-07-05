package service

import (
	"backend/internal/entity"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService
	host        *websocket.Conn
	tick        int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
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

const (
	PacketConnect      = "connect"
	PacketHost         = "host"
	PacketQuestionShow = "question"
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
		fmt.Println(connectPacket.Name, "wants to join the game", connectPacket.Code)

	case "host":

		var hostGamePacket HostGamePacket
		var err = json.Unmarshal(base.Data, &hostGamePacket)
		if err != nil {
			fmt.Println("Problem with packet")
			return
		}
		fmt.Println("User wants to host quiz", hostGamePacket.QuizId)
		go func() {
			time.Sleep(time.Second * 2)
			c.SendPacket(con, PacketQuestionShow, QuestionShowPacket{
				Question: entity.QuizQuestion{
					Name: "What is 2 + 2?",
					Choices: []entity.QuizChoice{
						{
							Name: "4",
						},
						{
							Name: "9",
						},
					},
				},
			})
		}()

	default:
		fmt.Println("Unknown packet type:", base.Type)

	}
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

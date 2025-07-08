package service

import (
	"backend/internal/entity"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Game struct {
	Id      uuid.UUID
	Quiz    entity.Quiz
	Code    string
	State   GameState
	Players []Player
	Time    int

	Host       *websocket.Conn
	netService *NetService
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

type Player struct {
	Name       string
	Connection *websocket.Conn
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func newGame(quiz entity.Quiz, host *websocket.Conn, netservice *NetService) Game {
	return Game{
		Id:      uuid.New(),
		Quiz:    quiz,
		Code:    generateCode(),
		State:   LobbyState,
		Players: []Player{},
		Time:    60,

		Host:       host,
		netService: netservice,
	}
}

func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	fmt.Println(name, "joined the game")
	g.Players = append(g.Players, Player{
		Name:       name,
		Connection: connection,
	})

	g.netService.SendPacket(connection, PacketChangeGameState, ChangeGameStatePacket{
		State: g.State,
	})
}

func (g *Game) Start() {
	g.ChangeState(PlayState, PacketChangeGameState)
	g.netService.SendPacket(g.Host, PacketQuestionShow, QuestionShowPacket{
		Question: entity.QuizQuestion{
			Name: "What is 2 + 2?",
			Choices: []entity.QuizChoice{
				{
					Name: "4",
				},
				{
					Name: "9",
				},
				{
					Name: "8",
				},
				{
					Name: "7",
				},
			},
		},
	})

	// fmt.Println("GAME START")

	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {
	g.Time--
	g.netService.SendPacket(g.Host, PacketTick, TickPacket{
		Tick: g.Time,
	})
}

func (g *Game) ChangeState(state GameState, packetCode string) {
	g.State = state
	g.BroadcastPacket(ChangeGameStatePacket{
		State: state,
	}, packetCode, true)
}
func (g *Game) BroadcastPacket(packet any, packetType string, includeHost bool) error {
	for _, player := range g.Players {
		err := g.netService.SendPacket(player.Connection, packetType, packet)
		if err != nil {
			return err
		}
	}

	if includeHost {
		err := g.netService.SendPacket(g.Host, packetType, packet)
		if err != nil {
			return err
		}
	}

	return nil
}

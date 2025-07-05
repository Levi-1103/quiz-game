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

	g.netService.SendPacket(connection, ChangeGameState, ChangeGameStatePacket{
		State: g.State,
	})
}

func (g *Game) Start() {
	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {
	panic("unimplemented")
}

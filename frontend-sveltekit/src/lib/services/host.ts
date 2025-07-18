import { GameState, PacketCode, type Packet } from "$lib/model/net";
import type { Player, QuizQuestion } from "$lib/model/quiz";
import { writable, type Writable } from "svelte/store";
import { NetService } from "./net";

export const gameState: Writable<GameState> =  writable(GameState.LobbyState)
export const players: Writable<Player[]> =  writable([])
export const tick: Writable<number> =  writable(0)
export const currentQuestion: Writable<QuizQuestion | null> =  writable(null)



export class HostGame {
    private net: NetService;

    constructor() {
        this.net = new NetService();
        this.net.connect();
        this.net.onPacket(p => {
            console.log(p)
            this.onPacket(p)
        });
    }
    onPacket(packet: Packet) {
        switch (packet.code) {
            case PacketCode.ChangeGameState: {
                let data = packet.data;
                gameState.set(data.state)
                break;
            }
            case PacketCode.PlayerJoin: {
                let data = packet.data;
                players.update(p => [...p, data.player]);
                break;
            }
            case PacketCode.Tick: {
                let data = packet.data;
                tick.set(data.tick)
                break;
            }
             case PacketCode.QuestionShow: {
                 let data = packet.data;
                currentQuestion.set(data.question)
                break;
            }
            
        }
    };

    hostQuiz(quizId: string) {
        this.net.sendPacket({
            code: PacketCode.Host,
            data: {
                quizId: quizId,
            },
        });
    };

    startGame() {
        this.net.sendPacket({
      code: PacketCode.StartGame,
    });
    }
}
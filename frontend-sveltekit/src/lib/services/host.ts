import { GameState, PacketCode, type Packet } from "$lib/model/net";
import type { Player } from "$lib/model/quiz";
import { writable, type Writable } from "svelte/store";
import { NetService } from "./net";

export const gameState: Writable<GameState> =  writable(GameState.LobbyState)
export const players: Writable<Player[]> =  writable([])



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
                break
            }
             case PacketCode.QuestionShow: {

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
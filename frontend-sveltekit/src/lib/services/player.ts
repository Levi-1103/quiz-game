import { GameState, PacketCode, type Packet } from "$lib/model/net";
import { writable, type Writable } from "svelte/store";
import { NetService } from "./net";

export const gameState: Writable<GameState> =  writable(GameState.LobbyState)


export class PlayerGame {
    private net: NetService;

    constructor() {
        this.net = new NetService();
        this.net.connect();
        this.net.onPacket(p => this.onPacket(p));
    }
    onPacket(packet: Packet) {
        switch (packet.code) {
            case PacketCode.ChangeGameState: {
                let data = packet.data;
                gameState.set(data.state)
                break;
            }
        }
    };

    join(code: string, name: string) {
        this.net.sendPacket({
            code: PacketCode.Connect,
            data: {
                name: name,
                code: code,
            },
        })
    }

}
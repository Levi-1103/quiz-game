import { PacketCode, type Packet } from "$lib/model/net";
import { gameState, players } from "./host";
import { NetService } from "./net";

export class PlayerGame {
    private net: NetService;

    constructor() {
        this.net = new NetService();
        this.net.connect();
        this.net.onPacket(p => this.onPacket(p));
    }
    onPacket(packet: Packet) {
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
/* eslint-disable no-console */
import type { Packet } from "../model/net";

export class NetService {
  private websocket!: WebSocket;

  private onPacketCallback?: (packet: any) => void;

  connect() {
    this.websocket = new WebSocket("ws://localhost:3000/ws");
    this.websocket.onopen = () => {
      console.log("connection opened");
    };

    this.websocket.onmessage = async (event) => {
      const packet: Packet = JSON.parse(event.data);

      if (this.onPacketCallback) {
        this.onPacketCallback(packet);
      }
    };
  }

  onPacket(callback: (packet: any) => void) {
    this.onPacketCallback = callback;
  }

  sendPacket(packet: Packet) {
    const packetToSend = packet;

    console.log(packetToSend);

    const jsonPacket = JSON.stringify(packetToSend);

    console.log(jsonPacket);

    this.websocket.send(jsonPacket);
  }
}

export function hello() {
  return "hello";
}

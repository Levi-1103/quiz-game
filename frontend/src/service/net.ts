export class NetService {
  private websocket!: WebSocket;
  private textDecoder: TextDecoder = new TextDecoder();
  private textEncoder: TextEncoder = new TextEncoder();

  private onPacketCallback?: (packet: any) => void;

  connect() {
    this.websocket = new WebSocket("ws://localhost:3000/ws");
    this.websocket.onopen = () => {
      console.log("connection opened");
    };

    this.websocket.onmessage = async (event) => {
      const arrayBuffer = await event.data.arrayBuffer();
      const bytes = new Uint8Array(arrayBuffer);
      const packetId = bytes[0];
      const packet = JSON.parse(this.textDecoder.decode(bytes.subarray(1)));
      packet.id = packetId;

      if (this.onPacketCallback) {
        this.onPacketCallback(packet);
      }
    };
  }

  onPacket(callback: (packet: any) => void) {
    this.onPacketCallback = callback;
  }

  sendPacket(packet: any) {
    const packetId = packet.id;
    const packetData = JSON.stringify(packet, (key, value) =>
      key == "id" ? undefined : value);

    const packetIdArray = new Uint8Array([packetId]);
    const packetDataArray = this.textEncoder.encode(packetData);

    const mergedArray = new Uint8Array(packetIdArray.length + packetDataArray.length);

    mergedArray.set(packetIdArray);
    mergedArray.set(packetDataArray, packetIdArray.length);

    this.websocket.send(mergedArray);
  }
}

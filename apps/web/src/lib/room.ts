import { Room as RoomModel } from "@/proto/room_pb";
import { RPC } from "./rpc";

export class Room {
  private roomInfo?: RoomModel;
  private rpc: RPC;

  constructor(private url: string, public token: string) {
    this.rpc = new RPC(this.url, this.token);
  }

  async startConnection(url: string, token: string): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const query = new URLSearchParams();
      query.append("access_token", token);

      const ws = new WebSocket(`${url}/rtc?${query.toString()}`);

      ws.onerror = (event) => {
        console.error("ws error", event);
        reject(false);
      };

      ws.onclose = (event) => {
        console.error("ws connection closed", event);
        reject(false);
      };

      ws.onmessage = (message) => {
        console.log("new meesage received", message);
        resolve(true);
      };
    });
  }
}

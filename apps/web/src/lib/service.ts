import {
  CreateRoomRequest,
  CreateRoomResponse,
  VerifyRoomRequest,
  VerifyRoomResponse,
} from "@/proto/twirp/v1/room_pb";

import { RPC } from "./rpc";

export class RTCService {
  private rpc: RPC;

  constructor(private url: string) {
    this.rpc = new RPC(this.url, "");
  }

  async createRoom(): Promise<CreateRoomResponse> {
    const data = await this.rpc.request(
      "RoomService",
      "CreateRoom",
      new CreateRoomRequest({ name: "Hernan" }).toJson()
    );

    if (!data) throw new Error("invalid response from rpc request");

    const res = CreateRoomResponse.fromJson(data);

    return res;
  }

  async verifyRoom(roomId: string) {
    const data = await this.rpc.request(
      "RoomService",
      "VerifyRoom",
      new VerifyRoomRequest({ roomId }).toJson()
    );

    if (!data) throw new Error("invalid response from rpc request");

    const res = VerifyRoomResponse.fromJson(data);

    return res;
  }

  async connectToRoom(roomId: string, token: string): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const query = new URLSearchParams();
      query.append("access_token", token);
      query.append("room_id", roomId);
      const wsUrl = this.url.replace("http", "ws");
      const ws = new WebSocket(`${wsUrl}/rtc?${query.toString()}`);

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

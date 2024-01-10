import {
  CreateRoomRequest,
  CreateRoomResponse,
  VerifyRoomRequest,
  VerifyRoomResponse,
} from "@/proto/twirp/v1/room_pb";

import { RPC } from "./rpc";
import { Room } from "./room";

export class RTCService {
  private rpc: RPC;

  constructor(private url: string) {
    this.rpc = new RPC(url, "");
  }

  async createRoom(): Promise<Room> {
    const data = await this.rpc.request(
      "RoomService",
      "CreateRoom",
      new CreateRoomRequest({ name: "Hernan" }).toJson()
    );

    if (!data) throw new Error("invalid response from rpc request");

    const res = CreateRoomResponse.fromJson(data);

    return new Room(res.roomId, this.url, res.accessToken);
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
}

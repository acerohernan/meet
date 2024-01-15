import { JoinResponse } from "@/proto/rtc_pb";
import { Room as RoomModel } from "@/proto/room_pb";

import { RPC } from "./rpc";
import { SignalClient } from "./signal";

export class Room {
  private roomInfo: RoomModel;
  private rpc: RPC;
  private signalClient: SignalClient;

  constructor(url: string, token: string, ws: WebSocket, join: JoinResponse) {
    if (!join.room) throw Error("no room found in join response");

    this.rpc = new RPC(url, token);
    this.signalClient = new SignalClient(ws);
    this.roomInfo = join.room;
  }
}

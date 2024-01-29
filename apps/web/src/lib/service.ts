import {
  CreateRoomRequest,
  CreateRoomResponse,
  VerifyRoomRequest,
  VerifyRoomResponse,
} from "@/proto/twirp/v1/room_pb";
import { SignalResponse } from "@/proto/rtc_pb";
import { GuestJoinResponse } from "@/proto/guest_pb";

import { RPC } from "./rpc";
import { Room } from "./room";
import { logger } from "./logger";
import { setupRoomListeners } from "./listeners";

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

  askJoin(roomId: string, name: string) {
    return new Promise<GuestJoinResponse | null>((resolve, reject) => {
      const query = new URLSearchParams();
      query.set("roomId", roomId);
      query.set("name", name);

      const eventSource = new EventSource(
        `${this.url}/join?${query.toString()}`
      );

      const abortHandler = () => {
        eventSource.close();
        reject(null);
      };

      eventSource.onmessage = (event) => {
        // close the connection at first message
        eventSource.close();

        if (!event.data) return abortHandler();

        try {
          const data = JSON.parse(event.data);
          const msg = new GuestJoinResponse().fromJson(data);
          resolve(msg);
        } catch (error) {
          logger.error("error at parsing join message", { data: event.data });
          abortHandler();
        }
      };

      eventSource.onerror = (event) => {
        logger.error("error at starting event source for join request", {
          event,
        });
        abortHandler();
      };
    });
  }

  connectToRoom(roomId: string, token: string): Promise<Room> {
    return new Promise((resolve, reject) => {
      const wsTimeout = setTimeout(reject, 3000);

      const query = new URLSearchParams();
      query.append("access_token", token);
      query.append("room_id", roomId);

      const wsUrl = this.url.replace("http", "ws");
      const ws = new WebSocket(`${wsUrl}/rtc?${query.toString()}`);
      ws.binaryType = "arraybuffer";

      const abortFn = () => {
        clearTimeout(wsTimeout);
        reject();
      };

      ws.onerror = abortFn;
      ws.onclose = abortFn;

      ws.onmessage = (event) => {
        let resp: SignalResponse | null = null;

        if (typeof event.data === "string") {
          const json = JSON.parse(event.data);
          resp = SignalResponse.fromJson(json);
        } else if (event.data instanceof ArrayBuffer) {
          resp = SignalResponse.fromBinary(new Uint8Array(event.data));
        } else {
          logger.error(
            `could not decode websocket message: ${typeof event.data}`
          );
          abortFn();
          return;
        }

        if (resp.response.case === "joinResponse") {
          clearTimeout(wsTimeout);

          const room = new Room(this.url, token, ws, resp.response.value);
          setupRoomListeners(room);
          resolve(room);
        } else {
          abortFn();
        }
      };
    });
  }
}

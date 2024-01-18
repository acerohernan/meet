import EventEmitter from "eventemitter3";

import {
  SignalResponse,
  SignalRequest,
  type ParticipantConnected,
  type ParticipantDisconnected,
  type ParticipantUpdated,
  type RefreshToken,
} from "@/proto/rtc_pb";

import { logger } from "./logger";
import { SignalEvents } from "./events";

export class SignalClient extends EventEmitter<SignalEventCallbacks> {
  constructor(private ws: WebSocket) {
    super();
    // setup listeners
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
        return;
      }

      this.handleSignalResponse(resp);
    };

    ws.onerror = (event) => {
      logger.error("ws error in signal client", event);
    };

    ws.onclose = (event) => {
      logger.error("ws connection closed in signal client", event);
    };
  }

  private handleSignalResponse(res: SignalResponse) {
    const msg = res.response;
    if (!msg) {
      logger.error("unsupported message received");
      return;
    }

    switch (msg.case) {
      case "refreshToken":
        this.emit(SignalEvents.RefreshToken, msg.value);
        break;

      case "participantConnected":
        this.emit(SignalEvents.ParticipantConnected, msg.value);
        break;

      case "participantUpdated":
        this.emit(SignalEvents.ParticipantUpdated, msg.value);
        break;

      case "participantDisconnected":
        this.emit(SignalEvents.ParticipantDisconnected, msg.value);
        break;

      default:
        logger.error("invalid ws message received", { res });
    }
  }

  private sendRequest(request: SignalRequest["request"]) {
    const req = new SignalRequest({ request });
    try {
      this.ws.send(req.toBinary());
    } catch (error) {
      logger.error("error sending signal request", { error });
    }
  }
}

interface SignalEventCallbacks {
  refreshToken: (res: RefreshToken) => void;
  participantConnected: (res: ParticipantConnected) => void;
  participantUpdated: (res: ParticipantUpdated) => void;
  participantDisconnected: (res: ParticipantDisconnected) => void;
}

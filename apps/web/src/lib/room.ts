import EventEmitter from "eventemitter3";

import {
  JoinResponse,
  ParticipantConnected,
  ParticipantDisconnected,
  ParticipantUpdated,
  RefreshToken,
} from "@/proto/rtc_pb";
import { Participant, Room as RoomModel } from "@/proto/room_pb";

import { RPC } from "./rpc";
import { logger } from "./logger";
import { SignalClient } from "./signal";
import { RoomEvents, SignalEvents } from "./events";

export class Room extends EventEmitter<RoomEventCallbacks> {
  private rpc: RPC;
  private signalClient: SignalClient;

  private roomInfo: RoomModel;
  participants: Map<string, Participant>;

  constructor(url: string, token: string, ws: WebSocket, join: JoinResponse) {
    super();
    this.rpc = new RPC(url, token);
    this.signalClient = new SignalClient(ws);
    this.setupSignalListeners();

    // apply join response information
    this.roomInfo = join.room!;
    this.participants = new Map();
    for (const p of join.participants) {
      this.participants.set(p.id, p);
    }
  }

  get information(): RoomModel {
    return this.roomInfo.toJson({
      emitDefaultValues: true,
    }) as unknown as RoomModel;
  }

  async closeConnection() {
    return this.signalClient.close();
  }

  // extend emitter to log all emitted events for development
  /** @internal */
  emit<E extends keyof RoomEventCallbacks>(
    event: E,
    ...args: EventEmitter.ArgumentMap<RoomEventCallbacks>[Extract<
      E,
      keyof RoomEventCallbacks
    >]
  ): boolean {
    logger.info(`room event ${event}`, { event, args });
    return super.emit(event, ...args);
  }

  /** @internal */
  private setupSignalListeners() {
    this.signalClient
      .on(SignalEvents.RefreshToken, this.handleRefreshToken)
      .on(SignalEvents.ParticipantConnected, this.handleParticipantConnected)
      .on(SignalEvents.ParticipantUpdated, this.handleParticipantUpdated)
      .on(
        SignalEvents.ParticipantDisconnected,
        this.handleParticipantDisconnected
      );
  }

  /** @internal */
  private handleRefreshToken = (res: RefreshToken) => {
    this.rpc.token = res.token;
    this.emit(RoomEvents.RefreshToken, res.token);
  };

  /** @internal */
  private handleParticipantConnected = (res: ParticipantConnected) => {
    if (!res.participant) return;
    this.participants.set(res.participant.id, res.participant);
    this.emit(RoomEvents.ParticipantConnected, res.participant);
  };

  /** @internal */
  private handleParticipantUpdated = (res: ParticipantUpdated) => {
    if (!res.participant) return;
    this.participants.set(res.participant.id, res.participant);
    this.emit(
      RoomEvents.ParticipantUpdated,
      res.participant.id,
      res.participant
    );
  };

  /** @internal */
  private handleParticipantDisconnected = (res: ParticipantDisconnected) => {
    this.participants.delete(res.participantId);
    this.emit(RoomEvents.ParticipantDisconnected, res.participantId);
  };
}

interface RoomEventCallbacks {
  refreshToken: (token: string) => void;
  participantConnected: (participant: Participant) => void;
  participantUpdated: (participantId: string, participant: Participant) => void;
  participantDisconnected: (participantID: string) => void;
}

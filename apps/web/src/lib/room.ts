import EventEmitter from "eventemitter3";
import { PlainMessage } from "@bufbuild/protobuf";

import {
  GuestAccepted,
  GuestDenied,
  GuestReqCancelled,
  JoinResponse,
  NewGuestRequest,
  ParticipantConnected,
  ParticipantDisconnected,
  ParticipantUpdated,
  RefreshToken,
} from "@/proto/rtc_pb";
import { Guest } from "@/proto/guest_pb";
import { Participant, Room as RoomModel } from "@/proto/room_pb";

import { RPC } from "./rpc";
import { logger } from "./logger";
import { SignalClient } from "./signal";
import { RoomEvents, SignalEvents } from "./events";

export class Room extends EventEmitter<RoomEventCallbacks> {
  private rpc: RPC;
  private signalClient: SignalClient;

  private roomInfo: RoomModel;

  private participantsMap: Map<string, PlainMessage<Participant>>;

  private guestsMap: Map<string, PlainMessage<Guest>>;

  constructor(url: string, token: string, ws: WebSocket, join: JoinResponse) {
    super();
    this.rpc = new RPC(url, token);
    this.signalClient = new SignalClient(ws);
    this.setupSignalListeners();

    // apply join response information
    this.roomInfo = join.room!;
    this.participantsMap = new Map();
    for (const p of join.participants) {
      this.participantsMap.set(p.id, p.toJson() as PlainMessage<Participant>);
    }
    this.guestsMap = new Map();
    for (const g of join.guests) {
      this.guestsMap.set(g.id, g.toJson() as PlainMessage<Guest>);
    }
  }

  get information(): RoomModel {
    return this.roomInfo.toJson({
      emitDefaultValues: true,
    }) as unknown as RoomModel;
  }

  get participants(): PlainMessage<Participant>[] {
    return Array.from(this.participantsMap.values());
  }

  get guests(): PlainMessage<Guest>[] {
    return Array.from(this.guestsMap.values());
  }

  async closeConnection() {
    return this.signalClient.close();
  }

  async acceptGuest(guestId: string): Promise<boolean> {
    const guest = this.guestsMap.get(guestId);
    if (!guest) return false;

    const success = await this.signalClient.sendGuestAnswer(guestId, {
      case: "guestAccepted",
      value: new GuestAccepted(),
    });
    if (!success) return false;

    // delete guest from local state
    this.guestsMap.delete(guestId);
    return true;
  }

  async denyGuest(guestId: string): Promise<boolean> {
    const guest = this.guestsMap.get(guestId);
    if (!guest) return false;

    const success = await this.signalClient.sendGuestAnswer(guestId, {
      case: "guestDenied",
      value: new GuestDenied(),
    });
    if (!success) return false;

    // delete guest from local state
    this.guestsMap.delete(guestId);
    return true;
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
      )
      .on(SignalEvents.GuestRequestReceived, this.handleGuestRequestReceived)
      .on(SignalEvents.GuestRequestCancelled, this.handleGuestRequestCancelled);
  }

  /** @internal */
  private handleRefreshToken = (res: RefreshToken) => {
    this.rpc.token = res.token;
    this.emit(RoomEvents.RefreshToken, res.token);
  };

  /** @internal */
  private handleParticipantConnected = (res: ParticipantConnected) => {
    if (!res.participant) return;
    this.participantsMap.set(
      res.participant.id,
      res.participant.toJson() as PlainMessage<Participant>
    );
    this.emit(RoomEvents.ParticipantConnected, res.participant);
  };

  /** @internal */
  private handleParticipantUpdated = (res: ParticipantUpdated) => {
    if (!res.participant) return;
    this.participantsMap.set(
      res.participant.id,
      res.participant.toJson() as PlainMessage<Participant>
    );
    this.emit(
      RoomEvents.ParticipantUpdated,
      res.participant.id,
      res.participant.toJson() as PlainMessage<Participant>
    );
  };

  /** @internal */
  private handleParticipantDisconnected = (res: ParticipantDisconnected) => {
    this.participantsMap.delete(res.participantId);
    this.emit(RoomEvents.ParticipantDisconnected, res.participantId);
  };

  /** @internal */
  private handleGuestRequestReceived = (res: NewGuestRequest) => {
    if (!res.guest) {
      logger.error("guest request received without guest", { res });
      return;
    }
    this.guestsMap.set(res.guest.id, res.guest.toJson() as PlainMessage<Guest>);
    this.emit(
      RoomEvents.GuestRequestReceived,
      res.guest.toJson() as PlainMessage<Guest>
    );
  };

  /** @internal */
  private handleGuestRequestCancelled = (res: GuestReqCancelled) => {
    this.guestsMap.delete(res.guestId);
    this.emit(RoomEvents.GuestRequestCancelled, res.guestId);
  };
}

interface RoomEventCallbacks {
  refreshToken: (token: string) => void;
  participantConnected: (participant: PlainMessage<Participant>) => void;
  participantUpdated: (
    participantId: string,
    participant: PlainMessage<Participant>
  ) => void;
  participantDisconnected: (participantID: string) => void;
  guestRequestReceived: (guest: PlainMessage<Guest>) => void;
  guestRequestCancelled: (guestId: string) => void;
}

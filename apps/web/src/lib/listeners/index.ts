import { store } from "@/store/store";
import { roomActions } from "@/store/room";

import { Room } from "../room";
import { RoomEvents } from "../events";
import { handleGuestCancelled, handleGuestReceived } from "./guest";
import {
  handleParticipantConnected,
  handleParticipantDisconnected,
  handleParticipantUpdated,
} from "./participant";

export const setupRoomListeners = (room: Room) => {
  store.dispatch(roomActions.guestsReceived(room.guests));
  store.dispatch(roomActions.participantsReceived(room.participants));

  // setup listeners
  room.on(RoomEvents.ParticipantConnected, handleParticipantConnected(room));
  room.on(RoomEvents.ParticipantUpdated, handleParticipantUpdated(room));
  room.on(
    RoomEvents.ParticipantDisconnected,
    handleParticipantDisconnected(room)
  );
  room.on(RoomEvents.GuestRequestReceived, handleGuestReceived(room));
  room.on(RoomEvents.GuestRequestCancelled, handleGuestCancelled(room));
};

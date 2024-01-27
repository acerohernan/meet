import { store } from "@/store/store";
import { roomActions } from "@/store/room";

import { Room } from "./room";
import { RoomEvents } from "./events";

export const setupStateListener = (room: Room) => {
  store.dispatch(roomActions.guestsReceived(room.guests));
  store.dispatch(roomActions.participantsReceived(room.participants));

  const participantListener = () => {
    store.dispatch(roomActions.participantsReceived(room.participants));
  };

  room.on(RoomEvents.ParticipantConnected, participantListener);
  room.on(RoomEvents.ParticipantUpdated, participantListener);
  room.on(RoomEvents.ParticipantDisconnected, participantListener);

  const guestListener = () => {
    store.dispatch(roomActions.guestsReceived(room.guests));
  };

  room.on(RoomEvents.GuestRequestReceived, guestListener);
  room.on(RoomEvents.GuestRequestCancelled, guestListener);
};

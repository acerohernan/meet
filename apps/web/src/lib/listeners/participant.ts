import { store } from "@/store/store";
import { roomActions } from "@/store/room";

import { Room } from "../room";

export const handleParticipantConnected = (room: Room) => () => {
  store.dispatch(roomActions.participantsReceived(room.participants));
};
export const handleParticipantUpdated = (room: Room) => () => {
  store.dispatch(roomActions.participantsReceived(room.participants));
};

export const handleParticipantDisconnected = (room: Room) => () => {
  store.dispatch(roomActions.participantsReceived(room.participants));
};

import { store } from "@/store/store";
import { roomActions } from "@/store/room";

import { Room } from "../room";

export const handleGuestReceived = (room: Room) => () => {
  store.dispatch(roomActions.guestsReceived(room.guests));
};

export const handleGuestCancelled = (room: Room) => () => {
  store.dispatch(roomActions.guestsReceived(room.guests));
};

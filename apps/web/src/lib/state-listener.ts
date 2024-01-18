import { store } from "@/store/store";
import { roomActions } from "@/store/room";

import { Room } from "./room";

export function setupStateListener(room: Room) {
  store.dispatch(roomActions.roomReceived(room.information));
}

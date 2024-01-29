import { useContext } from "react";

import { RoomContext } from "./index";

export const useRoomContext = () => {
  const ctx = useContext(RoomContext);
  if (!ctx) throw new Error("do not use room context outside the room page");
  return ctx;
};

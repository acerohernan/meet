import { Room } from "@/lib/room";

import { createContext } from "react";

interface RoomContext {
  roomId: string;
  room: Room | null;
  loading: boolean;
  closed: boolean;
  token: string;
  attempConnection: () => Promise<boolean>;
  closeConnection: () => Promise<void>;
}

export const RoomContext = createContext<RoomContext | undefined>(undefined);

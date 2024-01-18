import { Room } from "@/lib/room";

import { createContext } from "react";

interface RoomContext {
  room: Room | null;
  loading: boolean;
  closed: boolean;
  token: string;
  attempConnection: () => Promise<boolean>;
  closeConnection: () => Promise<void>;
}

export const RoomContext = createContext<RoomContext | undefined>(undefined);

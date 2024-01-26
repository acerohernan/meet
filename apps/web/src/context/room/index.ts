import { createContext } from "react";

import { type IRoomContext } from "./types";

export const RoomContext = createContext<IRoomContext | undefined>(undefined);

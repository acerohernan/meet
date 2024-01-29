import { Room } from "@/lib/room";

export interface IRoomContext {
  roomId: string;
  room: Room | null;
  loading: boolean;
  closed: boolean;
  token: string;
  setToken: (token: string) => void;
  attempConnection: (customToken?: string) => Promise<boolean>;
  closeConnection: () => Promise<void>;
}

export enum DrawerSection {
  People = "People",
  Information = "Information",
}

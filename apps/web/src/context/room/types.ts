import { Room } from "@/lib/room";

export interface IRoomContext {
  roomId: string;
  room: Room | null;
  loading: boolean;
  closed: boolean;
  token: string;
  attempConnection: () => Promise<boolean>;
  closeConnection: () => Promise<void>;
  isDrawerOpen: boolean;
  drawerSection: DrawerSection | null;
  openDrawer: (section: DrawerSection) => void;
  closeDrawer: () => void;
}

export enum DrawerSection {
  People = "People",
  Information = "Information",
}

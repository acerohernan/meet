import { RoomContextProvider } from "@/context/room";
import { MeetingPage } from "./meeting";

export const RoomPage = () => {
  return (
    <RoomContextProvider>
      <MeetingPage />
    </RoomContextProvider>
  );
};

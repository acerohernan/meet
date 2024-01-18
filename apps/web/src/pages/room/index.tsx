import { useAppSelector } from "@/store/store";

import { Meeting } from "./meeting";
import { WaitRoom } from "./wait-room";

const RoomPage = () => {
  const room = useAppSelector((state) => state.room.room);

  if (!room) return <WaitRoom />;

  return <Meeting />;
};

export default RoomPage;

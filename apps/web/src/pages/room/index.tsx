import { useRoomContext } from "@/context/room/hooks";

import { Meeting } from "./meeting";
import { WaitRoom } from "./wait-room";
import { ClosedRoom } from "./closed-room";

const RoomPage = () => {
  const { room, closed } = useRoomContext();

  if (room) return <Meeting />;

  if (closed) return <ClosedRoom />;

  return <WaitRoom />;
};

export default RoomPage;

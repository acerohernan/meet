/* import { accessTokenKey } from "@/constants/auth";
import { useMemo } from "react"; */
import { useState } from "react";

import { WaitRoom } from "./wait-room";
import { Meeting } from "./meeting";

const RoomPage = () => {
  const [waitRoom] = useState(true);

  if (waitRoom) return <WaitRoom />;

  return <Meeting />;
};

export default RoomPage;

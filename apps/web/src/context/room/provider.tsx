import { useParams } from "react-router-dom";
import { type PropsWithChildren, useCallback, useMemo, useState } from "react";

import { Room } from "@/lib/room";
import { logger } from "@/lib/logger";

import { rtcService } from "@/services/rtc";

import { useToast } from "@/hooks/useToast";
import { useAccessToken } from "@/hooks/useAccessToken";

import { RoomContext } from "./index";

const RoomContextProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const params = useParams();
  const roomId = useMemo(() => params.roomId ?? "", [params]);
  const [token, setToken] = useAccessToken({ roomId });
  const toast = useToast();

  const [loading, setLoading] = useState(false);
  const [closed, setClosed] = useState(false);

  const [room, setRoom] = useState<Room | null>(null);

  const attempConnection = async (customToken?: string) => {
    const accessToken = typeof customToken === "string" ? customToken : token;
    setLoading(true);
    try {
      const room = await rtcService.connectToRoom(roomId, accessToken);
      setRoom(room);
      return true;
    } catch (error) {
      logger.error("attemp connection err", { error });
      toast.error("Error at connecting with server via websockets!");
      return false;
    } finally {
      setLoading(false);
    }
  };

  const closeConnection = useCallback(async () => {
    try {
      if (!room) throw new Error("room not found");
      room.closeConnection();
    } catch (error) {
      logger.error("close connection err", { error });
      toast.error("Error at closing room connection");
    } finally {
      setClosed(true);
      setRoom(null);
    }
  }, [room, toast]);

  return (
    <RoomContext.Provider
      value={{
        room,
        roomId,
        token,
        setToken,
        loading,
        closed,
        attempConnection,
        closeConnection,
      }}
    >
      {children}
    </RoomContext.Provider>
  );
};

export default RoomContextProvider;

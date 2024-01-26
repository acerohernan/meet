import { useParams } from "react-router-dom";
import { type PropsWithChildren, useCallback, useMemo, useState } from "react";

import { Room } from "@/lib/room";
import { logger } from "@/lib/logger";

import { rtcService } from "@/services/rtc";

import { useToast } from "@/hooks/useToast";
import { useAccessToken } from "@/hooks/useAccessToken";

import { RoomContext } from "./index";
import { DrawerSection } from "./types";

const RoomContextProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const params = useParams();
  const roomId = useMemo(() => params.roomId ?? "", [params]);
  const [token] = useAccessToken({ roomId });
  const toast = useToast();

  const [loading, setLoading] = useState(false);
  const [closed, setClosed] = useState(false);

  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [drawerSection, setDrawerSection] = useState<DrawerSection>(
    DrawerSection.People
  );

  const [room, setRoom] = useState<Room | null>(null);

  const attempConnection = useCallback(async () => {
    setLoading(true);
    try {
      const room = await rtcService.connectToRoom(roomId, token);
      setRoom(room);
      return true;
    } catch (error) {
      logger.error("attemp connection err", { error });
      toast.error("Error at connecting with server via websockets!");
      return false;
    } finally {
      setLoading(false);
    }
  }, [roomId, token, toast]);

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

  const openDrawer = useCallback((section: DrawerSection) => {
    setDrawerSection(section);
    setIsDrawerOpen(true);
  }, []);

  const closeDrawer = useCallback(() => {
    setIsDrawerOpen(false);
  }, []);

  return (
    <RoomContext.Provider
      value={{
        room,
        roomId,
        token,
        loading,
        closed,
        attempConnection,
        closeConnection,
        openDrawer,
        closeDrawer,
        drawerSection,
        isDrawerOpen,
      }}
    >
      {children}
    </RoomContext.Provider>
  );
};

export default RoomContextProvider;

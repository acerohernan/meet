import { useContext, useEffect, useState } from "react";

import { RoomEvents } from "@/lib/events";

import { Participant } from "@/proto/room_pb";

import { RoomContext } from "./index";

export const useRoomContext = () => {
  const ctx = useContext(RoomContext);
  if (!ctx) throw new Error("do not use room context outside the room page");
  return ctx;
};

export const useParticipants = (): Participant[] => {
  const { room } = useRoomContext();

  const [participants, setParticipants] = useState<Participant[]>([]);

  useEffect(() => {
    if (!room) return;

    const handleParticipantChanges = () => {
      setParticipants(Array.from(room.participants.values()));
    };

    room.on(RoomEvents.ParticipantConnected, handleParticipantChanges);
    room.on(RoomEvents.ParticipantUpdated, handleParticipantChanges);
    room.on(RoomEvents.ParticipantDisconnected, handleParticipantChanges);

    return () => {
      room.off(RoomEvents.ParticipantConnected, handleParticipantChanges);
      room.off(RoomEvents.ParticipantUpdated, handleParticipantChanges);
      room.off(RoomEvents.ParticipantDisconnected, handleParticipantChanges);
    };
  }, [room]);

  return participants;
};

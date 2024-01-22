import { useContext, useEffect, useState } from "react";

import { RoomEvents } from "@/lib/events";

import { Guest } from "@/proto/guest_pb";
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
      setParticipants(room.participants);
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

export const useGuests = (): Guest[] => {
  const { room } = useRoomContext();

  const [guests, setGuests] = useState<Guest[]>([]);

  useEffect(() => {
    if (!room) return;

    const handleParticipantChanges = () => {
      setGuests(room.guests);
    };

    room.on(RoomEvents.GuestRequestReceived, handleParticipantChanges);
    room.on(RoomEvents.GuestRequestCancelled, handleParticipantChanges);

    return () => {
      room.off(RoomEvents.GuestRequestReceived, handleParticipantChanges);
      room.off(RoomEvents.GuestRequestCancelled, handleParticipantChanges);
    };
  }, [room]);

  return guests;
};

import { useCallback, useState } from "react";
import { PlainMessage } from "@bufbuild/protobuf";
import { Avatar, Box, Button, Typography } from "@mui/material";

import { roomActions } from "@/store/room";
import { useAppDispatch } from "@/store/store";

import { useToast } from "@/hooks/useToast";

import { Guest } from "@/proto/guest_pb";

import { useRoomContext } from "@/context/room/hooks";

interface Props {
  guest: PlainMessage<Guest>;
}

export const GuestWaitingCard: React.FC<Props> = ({ guest }) => {
  const toast = useToast();
  const { room } = useRoomContext();
  const dispatch = useAppDispatch();

  const [loading, setLoading] = useState(false);

  const admitGuest = useCallback(async () => {
    if (!room) return;
    setLoading(true);
    const success = await room.acceptGuest(guest.id);
    if (!success) {
      toast.error("something went wrong");
      return;
    }
    setLoading(false);
    dispatch(roomActions.guestsReceived(room.guests));
  }, [dispatch, guest.id, room, toast]);

  const denyGuest = useCallback(async () => {
    if (!room) return;
    setLoading(true);
    const success = await room.denyGuest(guest.id);
    if (!success) {
      toast.error("something went wrong");
      return;
    }
    setLoading(false);
    dispatch(roomActions.guestsReceived(room.guests));
  }, [dispatch, guest.id, room, toast]);

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          gap: 2,
        }}
      >
        <Avatar>{guest.name.slice(0, 1).toUpperCase()}</Avatar>
        <Typography fontSize="0.875rem">{guest.name}</Typography>
      </Box>
      <Box
        sx={{
          display: "flex",
          gap: 1,
        }}
      >
        <Button onClick={admitGuest} disabled={loading}>
          Admit
        </Button>
        <Button color="error" onClick={denyGuest} disabled={loading}>
          Deny
        </Button>
      </Box>
    </Box>
  );
};

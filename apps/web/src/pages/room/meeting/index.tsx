import { useEffect } from "react";
import { Box, Typography } from "@mui/material";
import { PlainMessage } from "@bufbuild/protobuf";

import { Guest } from "@/proto/guest_pb";

import { RoomEvents } from "@/lib/events";

import { useToast } from "@/hooks/useToast";

import { ParticipantCard } from "@/components/room/participant-card";

import { useRoomContext } from "@/context/room/hooks";

import { MeetingDrawer } from "./drawer";
import { MainControls } from "./main-controls";
import { DrawerControls } from "./drawer-controls";
import { NewGuestToast } from "./toasts/new-guest-toast";

export const Meeting = () => {
  const { room } = useRoomContext();
  const toast = useToast();

  useEffect(() => {
    if (!room) return;

    const handleGuestReceived = (guest: PlainMessage<Guest>) =>
      toast.custom(NewGuestToast, { guest });

    room.on(RoomEvents.GuestRequestReceived, handleGuestReceived);
    return () => {
      room.off(RoomEvents.GuestRequestReceived, handleGuestReceived);
    };
  }, [room, toast]);

  return (
    <Box
      width="100%"
      height="100vh"
      boxSizing="border-box"
      sx={{
        display: "flex",
        flexDirection: "column",
        padding: 2,
        backgroundColor: "#202124",
      }}
    >
      <Box
        sx={{
          flex: 1,
          marginBottom: 2,
          overflowY: "auto",
          overflowX: "hidden",
          position: "relative",
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexWrap: "wrap",
            justifyContent: "center",
            gap: 2,
            height: "100%",
          }}
        >
          <ParticipantCard />
        </Box>
        <MeetingDrawer />
      </Box>

      <Box
        width="100%"
        sx={{
          display: "grid",
          gridTemplateColumns: {
            xs: "1fr",
            sm: "1fr 2fr 1fr",
          },
          gap: 2,
          alignItems: "center",
        }}
      >
        <Box
          sx={{
            display: {
              xs: "none",
              sm: "flex",
            },
          }}
          alignItems="center"
          gap={2}
        >
          <Box
            sx={{
              display: {
                xs: "none",
                lg: "flex",
              },
            }}
            alignItems="center"
            justifyContent="center"
            gap={2}
          >
            <Typography color="white" fontSize="1.1rem" fontWeight={300}>
              9:44 PM
            </Typography>
            <Box
              sx={{
                width: "1px",
                height: "18px",
                backgroundColor: "white",
              }}
            />
          </Box>
          <Typography color="white" fontSize="1.1rem" fontWeight={300}>
            wes-rpdt-dqe
          </Typography>
        </Box>
        <MainControls />
        <DrawerControls />
      </Box>
    </Box>
  );
};

import { memo } from "react";
import { Toast } from "react-hot-toast";

import { Avatar, Box, Button, Typography } from "@mui/material";

import { useToast } from "@/hooks/useToast";

import { roomActions } from "@/store/room";
import { useAppDispatch } from "@/store/store";
import { DrawerSection } from "@/store/room/types";

import { Guest } from "@/proto/guest_pb";

export const NewGuestToast: React.FC<{ toast: Toast; guest: Guest }> = memo(
  ({ toast: t, guest }) => {
    const toast = useToast();
    const dispatch = useAppDispatch();

    return (
      <Box
        display="flex"
        alignItems="center"
        gap={2}
        sx={{ background: "#404144", borderRadius: 1 }}
        paddingY={1}
        paddingX={2}
      >
        <Avatar
          sx={{
            width: "30px",
            height: "30px",
            fontSize: "0.875rem",
            border: "2px solid white",
          }}
        >
          {guest.name.slice(0, 1).toUpperCase()}
        </Avatar>
        <Typography color="white" fontSize="0.875rem">
          Someone wants to join this call
        </Typography>
        <Button
          variant="text"
          sx={{
            color: "#8ab4f8",
          }}
          onClick={() => {
            dispatch(roomActions.openDrawer(DrawerSection.People));
            toast.dismiss(t.id);
          }}
        >
          View
        </Button>
      </Box>
    );
  }
);

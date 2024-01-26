import { GuestMicVisualizer } from "@/components/room/guest-mic-visualizer";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import { Avatar, Box, IconButton, Typography } from "@mui/material";

export const ParticipantCard = () => {
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
        <Avatar>D</Avatar>
        <Typography fontSize="0.875rem">Hernan Acero (You)</Typography>
      </Box>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          gap: 1,
        }}
      >
        <GuestMicVisualizer />
        <IconButton>
          <MoreVertIcon />
        </IconButton>
      </Box>
    </Box>
  );
};

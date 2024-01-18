import { useNavigate } from "react-router-dom";
import { Box, Button, Typography } from "@mui/material";

import { useRoomContext } from "@/context/room/hooks";

export const ClosedRoom = () => {
  const navigate = useNavigate();
  const { attempConnection, loading } = useRoomContext();

  return (
    <Box width="100%" height="100vh" textAlign="center" padding={3}>
      <Typography variant="h4" fontSize="2.25rem" marginTop={4}>
        You left the meeting
      </Typography>
      <Box display="flex" gap={1} alignItems="center" justifyContent="center">
        <Button
          variant="outlined"
          sx={{ marginTop: 4, fontSize: "0.9rem", fontWeight: 600 }}
          onClick={attempConnection}
          disabled={loading}
        >
          Rejoin
        </Button>
        <Button
          variant="contained"
          sx={{ marginTop: 4, fontSize: "0.9rem", fontWeight: 600 }}
          onClick={() => navigate("/")}
        >
          Return to home screen
        </Button>
      </Box>
    </Box>
  );
};

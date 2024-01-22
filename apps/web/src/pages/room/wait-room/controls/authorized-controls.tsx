import { Box, Button, Typography } from "@mui/material";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";

import { useRoomContext } from "@/context/room/hooks";

export const AuthorizedControls = () => {
  const { loading: isConnecting, attempConnection } = useRoomContext();

  return (
    <Box display="flex" flexDirection="column" alignItems="center">
      <Typography variant="h4" fontWeight={400} fontSize="1.8rem">
        Ready to join?
      </Typography>
      <Typography marginTop={3} fontSize="0.9rem">
        No one else here
      </Typography>
      <Box marginTop={2} display="flex" alignItems="center" gap={1}>
        <Button
          variant="contained"
          disabled={isConnecting}
          sx={{
            borderRadius: "25px",
            padding: "10px 25px",
          }}
          onClick={attempConnection}
        >
          Join now
        </Button>
        <Button
          variant="outlined"
          sx={{
            borderRadius: "25px",
            padding: "10px 25px",
          }}
          startIcon={<PresentToAllOutlinedIcon />}
        >
          Present
        </Button>
      </Box>
    </Box>
  );
};

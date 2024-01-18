import {
  Box,
  Button,
  CircularProgress,
  TextField,
  Typography,
} from "@mui/material";
import { useParams } from "react-router-dom";
import { useCallback, useMemo, useState } from "react";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";

import { rtcService } from "@/services/rtc";

import { useToast } from "@/hooks/useToast";
import { useAccessToken } from "@/hooks/useAccessToken";

interface Props {
  isLoading: boolean;
}

export const WaitRoomControls: React.FC<Props> = ({ isLoading }) => {
  const params = useParams();
  const roomId = useMemo(() => params.roomId ?? "", [params]);
  const toast = useToast();
  const [token] = useAccessToken({ roomId });
  const [isConnecting, setIsConnecting] = useState(false);

  const attempConnection = useCallback(async () => {
    setIsConnecting(true);
    try {
      rtcService.connectToRoom(roomId, token);
    } catch (error) {
      console.log(error);
      toast.error("Error at connecting with server via websockets!");
    } finally {
      setIsConnecting(false);
    }
  }, [roomId, token, toast]);

  if (isLoading)
    return (
      <Box display="flex" flexDirection="column" alignItems="center">
        <Typography variant="h4" fontWeight={400} fontSize="1.8rem">
          Getting ready...
        </Typography>
        <Typography marginTop={3} fontSize="0.9rem">
          You'll be able to join in just a moment
        </Typography>
        <Box marginTop={2}>
          <CircularProgress />
        </Box>
      </Box>
    );

  if (!token)
    return (
      <Box display="flex" flexDirection="column" alignItems="center">
        <Typography variant="h4" fontWeight={400} fontSize="1.8rem">
          What's your name?
        </Typography>
        <TextField
          label="Name"
          variant="filled"
          fullWidth
          sx={{ maxWidth: "300px", marginTop: 4 }}
        />
        <Box marginTop={4} display="flex" alignItems="center" gap={1}>
          <Button
            variant="contained"
            sx={{
              borderRadius: "25px",
              padding: "10px 25px",
            }}
          >
            Ask to join
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

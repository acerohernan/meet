import { Box, Button, TextField, Typography } from "@mui/material";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";
import { useCallback, useMemo, useState } from "react";
import { rtcService } from "@/services/rtc";
import { useToast } from "@/hooks/useToast";
import { logger } from "@/lib/logger";
import { useRoomContext } from "@/context/room/hooks";

export const GuestControls = () => {
  const toast = useToast();
  const { roomId } = useRoomContext();

  const [loading, setLoading] = useState(false);
  const [name, setName] = useState("");

  const isValidName = useMemo<boolean>(() => name.length > 4, [name]);

  const askToJoin = useCallback(async () => {
    setLoading(true);
    try {
      const answer = await rtcService.askJoin(roomId, name);
      console.log({ answer });
    } catch (error) {
      logger.error("error at asking join", { error });
      toast.error("Error at asking to join, please try it later.");
    } finally {
      setLoading(false);
    }
  }, [toast, name, roomId]);

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
        value={name}
        onChange={(event) => setName(event.target.value)}
      />
      <Box marginTop={4} display="flex" alignItems="center" gap={1}>
        <Button
          variant="contained"
          sx={{
            borderRadius: "25px",
            padding: "10px 25px",
          }}
          onClick={askToJoin}
          disabled={loading || !isValidName}
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
          onClick={askToJoin}
          disabled={loading}
        >
          Present
        </Button>
      </Box>
    </Box>
  );
};

import { useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import KeyboardIcon from "@mui/icons-material/Keyboard";
import VideoCallOutlinedIcon from "@mui/icons-material/VideoCallOutlined";

import { rtcService } from "@/services/rtc";

import { useToast } from "@/hooks/useToast";
import { saveRoomToken } from "@/hooks/useAccessToken";

const HomePage = () => {
  const [loading, setLoading] = useState(false);
  const toast = useToast();

  const navigate = useNavigate();

  async function createMeeting() {
    setLoading(true);

    try {
      const res = await rtcService.createRoom();
      saveRoomToken({ roomId: res.roomId, token: res.accessToken });
      setLoading(false);
      navigate(`/${res.roomId}`);
    } catch (err) {
      toast.error(
        "Something went wrong at the server. Please try again later."
      );
      setLoading(false);
    }
  }

  return (
    <Box
      width="100%"
      height="100vh"
      sx={{
        display: "flex",
        alignItems: {
          xs: "center",
          md: "center",
        },
      }}
    >
      <Box
        textAlign="center"
        width="100%"
        padding={4}
        maxWidth={700}
        margin="0px auto"
      >
        <Typography variant="h3" marginBottom={3}>
          Premium video meetings. Now free for everyone
        </Typography>
        <Typography fontWeight={300} fontSize="1.125rem">
          We re-engineered the service we built for secure bussiness meetings,
          to make it free and avaliable for all.
        </Typography>
        <Box
          width="100%"
          sx={{
            display: "flex",
            flexDirection: {
              xs: "column",
              md: "row",
            },
            alignItems: {
              xs: "start",
              md: "center",
            },
            justifyContent: "center",
            marginTop: {
              xs: 3,
              md: 8,
            },
          }}
          gap={3}
        >
          <Button
            variant="contained"
            size="large"
            css={{ fontSize: "1rem", flexShrink: 0, fontWeight: 600 }}
            startIcon={<VideoCallOutlinedIcon />}
            onClick={createMeeting}
            disabled={loading}
          >
            New meeting
          </Button>
          <Box width="100%" display="flex" gap={1}>
            <TextField
              variant="outlined"
              fullWidth
              placeholder="Enter a code or link"
              size="medium"
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <KeyboardIcon />
                  </InputAdornment>
                ),
              }}
            />
            <Button size="large" css={{ fontSize: "1rem" }} disabled>
              Join
            </Button>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

export default HomePage;

import { Box, IconButton } from "@mui/material";

import MicIcon from "@mui/icons-material/Mic";
import MicOffIcon from "@mui/icons-material/MicOff";
import VideocamOffOutlinedIcon from "@mui/icons-material/VideocamOffOutlined";
import VideocamOutlinedIcon from "@mui/icons-material/VideocamOutlined";

import { useState } from "react";
export const MeetingPage = () => {
  const [micOn, setMicOn] = useState(false);
  const [cameraOn, setCameraOn] = useState(false);

  function toggleMicrophone() {
    setMicOn((prev) => !prev);
  }

  function toggleCamera() {
    setCameraOn((prev) => !prev);
  }

  return (
    <Box
      width="100%"
      height="100vh"
      sx={{
        display: "flex",
        flexDirection: "column",
        backgroundColor: "#202124",
        padding: 4,
      }}
    >
      <Box sx={{ flex: 1 }}></Box>
      <Box width="100%" sx={{ display: "flex", gap: 2 }}>
        <IconButton
          color={micOn ? "active" : "inactive"}
          onClick={toggleMicrophone}
        >
          {micOn ? <MicIcon /> : <MicOffIcon />}
        </IconButton>
        <IconButton
          color={cameraOn ? "active" : "inactive"}
          onClick={toggleCamera}
        >
          {cameraOn ? <VideocamOutlinedIcon /> : <VideocamOffOutlinedIcon />}
        </IconButton>
      </Box>
    </Box>
  );
};

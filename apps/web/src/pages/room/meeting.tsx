import { Box, IconButton, Tooltip, Typography } from "@mui/material";
import MicIcon from "@mui/icons-material/Mic";
import MicOffIcon from "@mui/icons-material/MicOff";
import VideocamOffOutlinedIcon from "@mui/icons-material/VideocamOffOutlined";
import VideocamOutlinedIcon from "@mui/icons-material/VideocamOutlined";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import { useState } from "react";
export const MeetingPage = () => {
  const [micOn, setMicOn] = useState(false);
  const [cameraOn, setCameraOn] = useState(false);
  const [screenShareOn, setScreenShareOn] = useState(false);
  const [participantsOpened, setParticipantsOpened] = useState(false);

  function toggleMicrophone() {
    setMicOn((prev) => !prev);
  }

  function toggleCamera() {
    setCameraOn((prev) => !prev);
  }

  function toggleScreenShare() {
    setScreenShareOn((prev) => !prev);
  }

  function toggleParticipantsModal() {
    setParticipantsOpened((prev) => !prev);
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
        <Box
          sx={{
            display: "flex",
            gap: 2,

            justifyContent: "center",
          }}
        >
          <Tooltip title={`Turn ${micOn ? "on" : "off"} microphone`}>
            <IconButton
              color={micOn ? "on" : "off"}
              size="small"
              onClick={toggleMicrophone}
            >
              {micOn ? <MicIcon /> : <MicOffIcon />}
            </IconButton>
          </Tooltip>

          <Tooltip title={`Turn ${micOn ? "on" : "off"} camera`}>
            <IconButton
              color={cameraOn ? "on" : "off"}
              onClick={toggleCamera}
              size="small"
            >
              {cameraOn ? (
                <VideocamOutlinedIcon />
              ) : (
                <VideocamOffOutlinedIcon />
              )}
            </IconButton>
          </Tooltip>
          <Tooltip title={screenShareOn ? "Present now" : "Stop presenting"}>
            <IconButton
              color={screenShareOn ? "on" : "active"}
              onClick={toggleScreenShare}
              size="small"
            >
              <PresentToAllOutlinedIcon />
            </IconButton>
          </Tooltip>
        </Box>
        <Box
          sx={{
            display: {
              xs: "none",
              sm: "flex",
            },
            alignItems: "center",
            justifyContent: "end",
          }}
        >
          <IconButton
            onClick={toggleParticipantsModal}
            sx={{
              color: participantsOpened ? "#8ab4f8" : "white",
              padding: "14px",
            }}
          >
            {participantsOpened ? <PeopleAltIcon /> : <PeopleAltOutlinedIcon />}
          </IconButton>
        </Box>
      </Box>
    </Box>
  );
};

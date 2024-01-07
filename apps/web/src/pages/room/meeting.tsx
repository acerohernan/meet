import { useState } from "react";

import { Box, IconButton, Tooltip, Typography } from "@mui/material";

import MicIcon from "@mui/icons-material/Mic";
import InfoIcon from "@mui/icons-material/Info";
import MicOffIcon from "@mui/icons-material/MicOff";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import VideocamOutlinedIcon from "@mui/icons-material/VideocamOutlined";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";
import VideocamOffOutlinedIcon from "@mui/icons-material/VideocamOffOutlined";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";

import { ParticipantCard } from "@/components/room/participant-card";

export const MeetingPage = () => {
  const [micOn, setMicOn] = useState(false);
  const [cameraOn, setCameraOn] = useState(false);
  const [screenShareOn, setScreenShareOn] = useState(false);
  const [participantsOpened, setParticipantsOpened] = useState(false);
  const [infoOpened, setInfoOpened] = useState(false);

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

  function toggleInfoModal() {
    setInfoOpened((prev) => !prev);
  }

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
          overflow: "auto",
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
          <ParticipantCard />
          <ParticipantCard />
          <ParticipantCard />
        </Box>
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
          <Tooltip title={screenShareOn ? "Stop presenting" : "Present now"}>
            <IconButton
              color={screenShareOn ? "active" : "on"}
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
            onClick={toggleInfoModal}
            sx={{
              color: infoOpened ? "#8ab4f8" : "white",
              padding: "14px",
            }}
          >
            {infoOpened ? <InfoIcon /> : <InfoOutlinedIcon />}
          </IconButton>
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

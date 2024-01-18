import { useState } from "react";

import { Box, IconButton, Typography } from "@mui/material";

import InfoIcon from "@mui/icons-material/Info";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";

import { ParticipantCard } from "@/components/room/participant-card";
import { useParticipants, useRoomContext } from "@/context/room/hooks";
import { MainControls } from "./main-controls";

export const Meeting = () => {
  const [participantsOpened, setParticipantsOpened] = useState(false);
  const [infoOpened, setInfoOpened] = useState(false);

  const { room } = useRoomContext();
  const participants = useParticipants();

  if (!room) return;

  console.log({ participants });

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
        <MainControls />
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

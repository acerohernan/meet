import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Button,
  Typography,
} from "@mui/material";

import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import PersonAddAltOutlinedIcon from "@mui/icons-material/PersonAddAltOutlined";

import { GuestWaitingCard } from "./waiting-card";
import { ParticipantCard } from "./participant-card";

export const PeopleSection = () => {
  return (
    <Box>
      <Button
        variant="outlined"
        startIcon={<PersonAddAltOutlinedIcon />}
        sx={{ marginTop: 2 }}
      >
        Add people
      </Button>
      <Box display="grid" gap={3}>
        <Box marginTop={4}>
          <Typography
            fontSize="0.75rem"
            fontWeight="600"
            marginBottom={1}
            color="#5f6368"
            letterSpacing={0.8}
          >
            WAITING TO JOIN
          </Typography>
          <Accordion elevation={0} sx={{ border: "1px solid #DADCE0" }}>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls="panel1-content"
              id="panel1-header"
              sx={{
                fontSize: "0.875rem",
              }}
            >
              Waiting to be admitted
            </AccordionSummary>
            <AccordionDetails
              sx={{
                borderTop: "1px solid #DADCE0",
              }}
            >
              <Box
                display="flex"
                gap={2}
                alignItems="center"
                justifyContent="end"
                marginTop={1}
              >
                <Button variant="outlined">Admit all</Button>
                <Button variant="outlined" color="error">
                  Deny all
                </Button>
              </Box>
              <Box marginTop={3} display="grid" gap={3}>
                <GuestWaitingCard />
                <GuestWaitingCard />
                <GuestWaitingCard />
              </Box>
            </AccordionDetails>
          </Accordion>
        </Box>

        <Box>
          <Typography
            fontSize="0.75rem"
            fontWeight="600"
            marginBottom={1}
            color="#5f6368"
            letterSpacing={0.8}
          >
            IN MEETING
          </Typography>
          <Accordion elevation={0} sx={{ border: "1px solid #DADCE0" }}>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls="panel1-content"
              id="panel1-header"
              sx={{
                fontSize: "0.875rem",
              }}
            >
              Participants
            </AccordionSummary>
            <AccordionDetails
              sx={{
                borderTop: "1px solid #DADCE0",
              }}
            >
              <Box marginTop={1} display="grid" gap={3}>
                <ParticipantCard />
                <ParticipantCard />
              </Box>
            </AccordionDetails>
          </Accordion>
        </Box>
      </Box>
    </Box>
  );
};

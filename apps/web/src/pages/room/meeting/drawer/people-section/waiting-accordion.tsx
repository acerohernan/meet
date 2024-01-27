import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Button,
  Typography,
} from "@mui/material";

import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

import { GuestWaitingCard } from "./waiting-card";
import { useAppSelector } from "@/store/store";
import { useRoomContext } from "@/context/room/hooks";

export const WaitingAccordion = () => {
  const { room } = useRoomContext();
  const guests = useAppSelector((state) => state.room.guests);

  if (!room) return;
  if (guests.length < 1) return;

  return (
    <Box>
      <Typography
        fontSize="0.75rem"
        fontWeight="600"
        marginBottom={1}
        color="#5f6368"
        letterSpacing={0.8}
      >
        WAITING TO JOIN
      </Typography>
      <Accordion
        defaultExpanded
        elevation={0}
        sx={{ border: "1px solid #DADCE0" }}
      >
        <AccordionSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="waiting-guests-accordion-content"
          id="waiting-guests-accordion-header"
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
            <Button
              variant="outlined"
              onClick={() => {
                // room.admitAllGuests()
              }}
            >
              Admit all
            </Button>
            <Button
              variant="outlined"
              color="error"
              onClick={() => {
                // room.denyAllGuests()
              }}
            >
              Deny all
            </Button>
          </Box>
          <Box marginTop={3} display="grid" gap={3}>
            {guests.map((guest) => (
              <GuestWaitingCard guest={guest} key={guest.id} />
            ))}
          </Box>
        </AccordionDetails>
      </Accordion>
    </Box>
  );
};

import { Box, Button } from "@mui/material";

import PersonAddAltOutlinedIcon from "@mui/icons-material/PersonAddAltOutlined";

import { WaitingAccordion } from "./waiting-accordion";
import { ParticipantsAccordion } from "./participants-accordion";

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
      <Box display="grid" gap={3} paddingY={4}>
        <WaitingAccordion />
        <ParticipantsAccordion />
      </Box>
    </Box>
  );
};

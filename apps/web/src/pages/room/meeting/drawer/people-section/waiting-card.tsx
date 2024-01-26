import { Avatar, Box, Button, Typography } from "@mui/material";

export const GuestWaitingCard = () => {
  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          gap: 2,
        }}
      >
        <Avatar>D</Avatar>
        <Typography fontSize="0.875rem">Hernan</Typography>
      </Box>
      <Box
        sx={{
          display: "flex",
          gap: 1,
        }}
      >
        <Button>Admit</Button>
        <Button color="error">Deny</Button>
      </Box>
    </Box>
  );
};

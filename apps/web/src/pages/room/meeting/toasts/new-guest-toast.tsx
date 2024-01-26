import { Avatar, Box, Button, Typography } from "@mui/material";

export const NewGuestToast = () => {
  return (
    <Box
      display="flex"
      alignItems="center"
      gap={2}
      sx={{ background: "#404144", borderRadius: 1 }}
      paddingY={1}
      paddingX={2}
    >
      <Avatar
        sx={{
          width: "30px",
          height: "30px",
          fontSize: "0.875rem",
          border: "2px solid white",
        }}
      >
        H
      </Avatar>
      <Typography color="white" fontSize="0.875rem">
        Someone wants to join this call
      </Typography>
      <Button
        variant="text"
        sx={{
          color: "#8ab4f8",
        }}
      >
        View
      </Button>
    </Box>
  );
};

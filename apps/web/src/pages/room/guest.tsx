import { GuestPreview } from "@/components/room/guest-preview";
import { Box, Button, Typography } from "@mui/material";

export const GuestPage = () => {
  return (
    <Box
      width="100%"
      height="100vh"
      padding={3}
      display="flex"
      alignItems="center"
      justifyContent="center"
    >
      <Box
        width="100%"
        maxWidth="1100px"
        margin="0px auto"
        alignItems="center"
        justifyContent="center"
        display="grid"
        sx={{
          gridTemplateColumns: {
            xs: "1fr",
            md: "1fr 250px",
            lg: "1fr 400px",
          },
          gap: {
            xs: "32px",
            md: "0px",
          },
        }}
      >
        <Box maxWidth="800px">
          <GuestPreview />
        </Box>
        <Box display="flex" flexDirection="column" alignItems="center">
          <Typography variant="h4" fontWeight={300} fontSize="1.75rem">
            Ready to join?
          </Typography>
          <Typography marginTop={3} fontSize="0.9rem">
            No one else here
          </Typography>
          <Box marginTop={2}>
            <Button
              variant="contained"
              sx={{
                borderRadius: "25px",
                padding: "10px 25px",
                fontWeight: 300,
                fontSize: "1rem",
              }}
            >
              Join now
            </Button>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

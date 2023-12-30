import { GuestPreview } from "@/components/room/guest-preview";
import { Box, Button, Typography } from "@mui/material";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";
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
        maxWidth="1000px"
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
          <Box marginTop={2} display="flex" alignItems="center" gap={1}>
            <Button
              variant="contained"
              sx={{
                borderRadius: "25px",
                padding: "10px 25px",
                fontWeight: 300,
                fontSize: "0.9rem",
              }}
            >
              Join now
            </Button>
            <Button
              variant="contained"
              sx={{
                borderRadius: "25px",
                backgroundColor: "white",
                padding: "10px 25px",
                fontWeight: 300,
                fontSize: "0.9rem",
                color: "#1a73e8",
              }}
              startIcon={<PresentToAllOutlinedIcon />}
            >
              Present
            </Button>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

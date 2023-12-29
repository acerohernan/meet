import { Box, Typography } from "@mui/material";

export const GuestPreview = () => {
  return (
    <Box
      width="100%"
      height="100%"
      borderRadius={3}
      sx={{
        background: "#202124",
        display: "grid",
        placeItems: "center",
        aspectRatio: "4/2.5",
        minWidth: "300px",
      }}
    >
      <Typography variant="h5" sx={{ color: "white", fontWeight: 300 }}>
        Camera is off
      </Typography>
    </Box>
  );
};

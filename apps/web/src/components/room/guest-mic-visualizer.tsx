import { Box } from "@mui/material";

export const GuestMicVisualizer = () => {
  return (
    <Box
      width="100%"
      height="100%"
      minHeight={30}
      minWidth={30}
      borderRadius="100%"
      display="flex"
      alignItems="center"
      justifyContent="center"
      gap="2px"
      sx={{
        backgroundColor: "#1a73e8e6",
      }}
    >
      <PointVisualizer />
      <PointVisualizer />
      <PointVisualizer />
    </Box>
  );
};

const PointVisualizer = () => {
  return (
    <Box
      width={5}
      height={5}
      borderRadius="100%"
      sx={{
        backgroundColor: "white",
      }}
    />
  );
};

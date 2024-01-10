import { useNavigate, useParams } from "react-router-dom";
import { Box, Button, CircularProgress, Typography } from "@mui/material";
import PresentToAllOutlinedIcon from "@mui/icons-material/PresentToAllOutlined";

import { GuestPreview } from "@/components/room/guest-preview";

import { useRoomVerification } from "@/hooks/useRoomVerification";
import { useMemo } from "react";

export const WaitRoom = () => {
  const { roomId } = useParams();
  const navigate = useNavigate();
  const { isLoading, isError, verification } = useRoomVerification({
    roomId: roomId ?? "",
  });

  const roomNotExists = useMemo(() => {
    if (!verification) return false;
    return !isLoading && !verification.exists;
  }, [isLoading, verification]);

  if (isError || roomNotExists)
    return (
      <Box width="100%" height="100vh" textAlign="center" padding={3}>
        <Typography variant="h4" fontSize="2.25rem" marginTop={4}>
          Check your meeting code
        </Typography>
        <Typography margin="auto" marginTop={4} maxWidth="450px">
          Make sure you entered the correct meeting in the URL, for example:{" "}
          {import.meta.env.VITE_URL}/<strong>xxx-yyyy-zzz</strong>
        </Typography>
        <Button
          variant="contained"
          sx={{ marginTop: 4, fontSize: "0.9rem", fontWeight: 600 }}
          onClick={() => navigate("/")}
        >
          Return to home screen
        </Button>
      </Box>
    );

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
            sm: "1fr",
            md: "1fr 300px",
            lg: "1fr 400px",
          },
          gap: {
            xs: "32px",
            md: "0px",
          },
        }}
      >
        <Box maxWidth="800px">
          <GuestPreview isLoading={isLoading} />
        </Box>
        {isLoading ? (
          <Box display="flex" flexDirection="column" alignItems="center">
            <Typography variant="h4" fontWeight={400} fontSize="1.8rem">
              Getting ready...
            </Typography>
            <Typography marginTop={3} fontSize="0.9rem">
              You'll be able to join in just a moment
            </Typography>
            <Box marginTop={2}>
              <CircularProgress />
            </Box>
          </Box>
        ) : (
          <Box display="flex" flexDirection="column" alignItems="center">
            <Typography variant="h4" fontWeight={400} fontSize="1.8rem">
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
                  fontWeight: 400,
                  fontSize: "0.9rem",
                }}
              >
                Join now
              </Button>
              <Button
                variant="outlined"
                sx={{
                  borderRadius: "25px",
                  padding: "10px 25px",
                  fontWeight: 400,
                  fontSize: "0.9rem",
                }}
                startIcon={<PresentToAllOutlinedIcon />}
              >
                Present
              </Button>
            </Box>
          </Box>
        )}
      </Box>
    </Box>
  );
};

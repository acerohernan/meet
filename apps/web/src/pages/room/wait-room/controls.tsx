import { Box, CircularProgress, Typography } from "@mui/material";

import { useRoomContext } from "@/context/room/hooks";

import { AuthorizedControls } from "./controls/authorized-controls";
import { GuestControls } from "./controls/guest-controls";

interface Props {
  isLoading: boolean;
}

export const WaitRoomControls: React.FC<Props> = ({ isLoading }) => {
  const { token } = useRoomContext();

  if (isLoading)
    return (
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
    );

  if (!token) return <GuestControls />;

  return <AuthorizedControls />;
};

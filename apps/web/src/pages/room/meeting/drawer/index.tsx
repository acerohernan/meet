import { Box, IconButton, Typography } from "@mui/material";

import CloseIcon from "@mui/icons-material/Close";

import { useRoomContext } from "@/context/room/hooks";
import { DrawerSection } from "@/context/room/types";

import { PeopleSection } from "./people-section";

import "./index.css";

export const MeetingDrawer = () => {
  const { closeDrawer, drawerSection, isDrawerOpen } = useRoomContext();

  let title = "";
  let section = <></>;

  switch (drawerSection) {
    case DrawerSection.People:
      title = "People";
      section = <PeopleSection />;
      break;
  }

  return (
    <Box
      id="meeting-drawer"
      width="100%"
      maxWidth="380px"
      height="100%"
      position="absolute"
      paddingY={1}
      paddingX={2}
      top={0}
      right={0}
      sx={{
        backgroundColor: "white",
        borderRadius: 2,
        overflowY: "auto",
        overflowX: "hidden",
        transition: "transform 0.3s ease-in-out",
        transform: isDrawerOpen ? "translateX(0%)" : "translateX(100%)",
      }}
    >
      <Box display="flex" alignItems="center" justifyContent="space-between">
        <Typography variant="h6" fontSize="1.125rem" marginLeft={1}>
          {title}
        </Typography>
        <IconButton onClick={closeDrawer}>
          <CloseIcon />
        </IconButton>
      </Box>
      {section}
    </Box>
  );
};

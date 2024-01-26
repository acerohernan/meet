import { Box, IconButton, Typography } from "@mui/material";

import CloseIcon from "@mui/icons-material/Close";

import { DrawerSection } from "@/context/room/types";

import { PeopleSection } from "./people-section";

import "./index.css";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { roomActions } from "@/store/room";

export const MeetingDrawer = () => {
  const dispatch = useAppDispatch();

  const isDrawerOpen = useAppSelector((state) => state.room.isDrawerOpen);
  const drawerSection = useAppSelector((state) => state.room.drawerSection);

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
        <IconButton onClick={() => dispatch(roomActions.closeDrawer())}>
          <CloseIcon />
        </IconButton>
      </Box>
      {section}
    </Box>
  );
};

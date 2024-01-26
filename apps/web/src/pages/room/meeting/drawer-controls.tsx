import { Box, IconButton } from "@mui/material";

import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";

import { useRoomContext } from "@/context/room/hooks";
import { DrawerSection } from "@/context/room/types";

interface DrawerControl {
  section: DrawerSection;
  activeIcon: React.ReactNode;
  inactiveIcon: React.ReactNode;
}

const controls: DrawerControl[] = [
  {
    section: DrawerSection.People,
    activeIcon: <PeopleAltIcon />,
    inactiveIcon: <PeopleAltOutlinedIcon />,
  },
];

export const DrawerControls = () => {
  const { openDrawer, drawerSection, closeDrawer, isDrawerOpen } =
    useRoomContext();

  return (
    <Box
      sx={{
        display: {
          xs: "none",
          sm: "flex",
        },
        alignItems: "center",
        justifyContent: "end",
      }}
    >
      {controls.map(({ section, activeIcon, inactiveIcon }) => {
        const isActive = drawerSection === section && isDrawerOpen;

        return (
          <IconButton
            onClick={() => (isActive ? closeDrawer() : openDrawer(section))}
            sx={{
              color: isActive ? "#8ab4f8" : "white",
              padding: "14px",
            }}
            key={`drawer-controls-${section}`}
          >
            {isActive ? activeIcon : inactiveIcon}
          </IconButton>
        );
      })}
    </Box>
  );
};

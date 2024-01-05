import { createTheme } from "@mui/material";

declare module "@mui/material/IconButton" {
  interface IconButtonPropsColorOverrides {
    active: true;
    inactive: true;
  }
}

export const theme = createTheme({
  components: {
    MuiIconButton: {
      variants: [
        {
          props: { color: "active" },
          style: {
            border: "1px solid white",
            backgroundColor: "transparent",
            padding: "14px",
            "&:hover": { backgroundColor: "#9A9A9C" },
            color: "white",
          },
        },
        {
          props: { color: "inactive" },
          style: {
            border: "1px solid #ea4335",
            padding: "14px",
            backgroundColor: "#ea4335",
            color: "white",
            "&:hover": { backgroundColor: "#C93C33" },
          },
        },
      ],
    },
  },
  typography: {
    button: {
      textTransform: "none",
    },
  },
  palette: {
    primary: {
      main: "#1a73e8",
    },
  },
  breakpoints: {
    values: {
      xs: 0,
      sm: 540,
      md: 700,
      lg: 960,
      xl: 1280,
    },
  },
});

import { createTheme } from "@mui/material";

export const theme = createTheme({
  typography: {
    button: {
      textTransform: "none",
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

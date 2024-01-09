import { Provider } from "react-redux";
import { Toaster } from "react-hot-toast";
import { RouterProvider } from "react-router-dom";
import { ThemeProvider, CssBaseline } from "@mui/material";

import { store } from "./store/store";

import { theme } from "./theme";
import { router } from "./router";

function App() {
  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <Toaster />
        <CssBaseline />
        <RouterProvider router={router} />
      </ThemeProvider>
    </Provider>
  );
}

export default App;

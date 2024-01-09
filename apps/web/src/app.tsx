import { store } from "./store/store";
import { Toaster } from "react-hot-toast";
import { RouterProvider } from "react-router-dom";
import { ThemeProvider, CssBaseline } from "@mui/material";

import { Provider } from "react-redux";
import { router } from "./router";
import { theme } from "./theme";

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

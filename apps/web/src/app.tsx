import { Suspense } from "react";
import { Provider } from "react-redux";
import { Toaster } from "react-hot-toast";
import { RouterProvider } from "react-router-dom";
import { ThemeProvider, CssBaseline } from "@mui/material";

import { LoadingPage } from "./pages/loading";

import { store } from "./store/store";

import { theme } from "./theme";
import { router } from "./router";

function App() {
  return (
    <Suspense fallback={<LoadingPage />}>
      <Provider store={store}>
        <ThemeProvider theme={theme}>
          <Toaster />
          <CssBaseline />
          <RouterProvider router={router} />
        </ThemeProvider>
      </Provider>
    </Suspense>
  );
}

export default App;

import { createBrowserRouter } from "react-router-dom";

import { HomePage } from "@/app/shared/pages/home";

export const router = createBrowserRouter([
  {
    path: "",
    element: <HomePage />,
  },
]);

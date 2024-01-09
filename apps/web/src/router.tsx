/* eslint-disable react-refresh/only-export-components */
import { lazy } from "react";
import { createBrowserRouter } from "react-router-dom";

import HomePage from "@/pages/home";

const RoomPage = lazy(() => import("@/pages/room"));

export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/:roomId",
    element: <RoomPage />,
  },
]);

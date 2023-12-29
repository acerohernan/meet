import { createBrowserRouter } from "react-router-dom";

import { HomePage } from "@/pages/home";
import { RoomPage } from "@/pages/room";
import { GuestPage } from "./pages/room/guest";
import { MeetingPage } from "./pages/room/meeting";

export const router = createBrowserRouter([
  {
    path: "",
    element: <HomePage />,
  },
  {
    path: "guest",
    element: <GuestPage />,
  },
  {
    path: "meeting",
    element: <MeetingPage />,
  },
  {
    path: ":roomId",
    element: <RoomPage />,
  },
]);

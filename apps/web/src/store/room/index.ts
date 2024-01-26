import { PayloadAction, createSlice } from "@reduxjs/toolkit";

import { DrawerSection } from "@/context/room/types";

export interface RoomState {
  drawerSection: DrawerSection;
  isDrawerOpen: boolean;
}

const initialState: RoomState = {
  drawerSection: DrawerSection.People,
  isDrawerOpen: false,
};

const roomSlice = createSlice({
  name: "room",
  initialState,
  reducers: {
    openDrawer: (state, action: PayloadAction<DrawerSection>) => {
      state.drawerSection = action.payload;
      state.isDrawerOpen = true;
    },
    closeDrawer: (state) => {
      state.isDrawerOpen = false;
    },
  },
});

export const roomActions = roomSlice.actions;
export const roomReducer = roomSlice.reducer;

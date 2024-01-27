import { PlainMessage } from "@bufbuild/protobuf";
import { PayloadAction, createSlice } from "@reduxjs/toolkit";

import { Guest } from "@/proto/guest_pb";
import { Participant } from "@/proto/room_pb";

import { DrawerSection } from "@/context/room/types";

export interface RoomState {
  drawerSection: DrawerSection;
  isDrawerOpen: boolean;
  participants: PlainMessage<Participant>[];
  guests: PlainMessage<Guest>[];
}

const initialState: RoomState = {
  drawerSection: DrawerSection.People,
  isDrawerOpen: false,
  participants: [],
  guests: [],
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
    participantsReceived: (
      state,
      action: PayloadAction<PlainMessage<Participant>[]>
    ) => {
      state.participants = action.payload;
    },
    guestsReceived: (state, action: PayloadAction<PlainMessage<Guest>[]>) => {
      state.guests = action.payload;
    },
  },
});

export const roomActions = roomSlice.actions;
export const roomReducer = roomSlice.reducer;

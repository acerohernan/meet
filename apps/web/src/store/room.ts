import { PlainMessage } from "@bufbuild/protobuf";
import { PayloadAction, createSlice } from "@reduxjs/toolkit";

import { Room } from "@/proto/room_pb";

export interface RoomState {
  room: PlainMessage<Room> | null;
}

const initialState: RoomState = {
  room: null,
};

const roomSlice = createSlice({
  name: "room",
  initialState,
  reducers: {
    roomReceived: (state, action: PayloadAction<PlainMessage<Room> | null>) => {
      state.room = action.payload;
    },
  },
});

export const roomActions = roomSlice.actions;
export const roomReducer = roomSlice.reducer;

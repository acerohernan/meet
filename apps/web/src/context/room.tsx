import React, { PropsWithChildren, createContext } from "react";

interface RoomContex {}

export const RoomContext = createContext<RoomContex | null>(null);

export const RoomContextProvider: React.FC<PropsWithChildren> = ({
  children,
}) => {
  // get token from session storage
  // if there's a token, try to reconnect to the server
  //

  return <RoomContext.Provider value={null}>{children}</RoomContext.Provider>;
};

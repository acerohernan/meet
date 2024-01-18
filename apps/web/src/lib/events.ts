export const SignalEvents = {
  RefreshToken: "refreshToken",
  ParticipantConnected: "participantConnected",
  ParticipantUpdated: "participantUpdated",
  ParticipantDisconnected: "participantDisconnected",
} as const;

export const RoomEvents = {
  /**
   * When new access token is received
   *
   * @param token string
   */
  RefreshToken: "refreshToken",

  /**
   * When new participant has joined the room
   *
   * @param participant Participant
   */
  ParticipantConnected: "participantConnected",

  /**
   * When participant information was updated
   *
   * @param participantId string
   * @param participant Participant
   */
  ParticipantUpdated: "participantUpdated",

  /**
   * When participant has left the room
   *
   * @param participantID string
   */
  ParticipantDisconnected: "participantDisconnected",
} as const;

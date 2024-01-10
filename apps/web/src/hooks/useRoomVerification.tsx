import { useEffect, useState } from "react";

import { rtcService } from "@/services/rtc";

import { VerifyRoomResponse } from "@/proto/twirp/v1/room_pb";

interface Props {
  roomId: string;
}

export const useRoomVerification = ({ roomId }: Props) => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);
  const [verification, setVerification] = useState<VerifyRoomResponse | null>(
    null
  );

  useEffect(() => {
    const fetchVerification = async () => {
      setLoading(true);

      try {
        const verification = await rtcService.verifyRoom(roomId);
        //await new Promise<void>((res) => setTimeout(res, 10000));
        setVerification(verification);
      } catch {
        setError(true);
      } finally {
        setLoading(false);
      }
    };
    fetchVerification();
  }, [roomId]);

  return { isLoading: loading, isError: error, verification };
};

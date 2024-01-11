import { useCallback, useMemo, useState } from "react";

const accessTokensKey = "__tokens";

export const useAccessToken = ({ roomId }: { roomId: string }) => {
  const [tokens, setTokens] = useState<Record<string, string>>(() => {
    const tokenMap = sessionStorage.getItem(accessTokensKey);
    if (!tokenMap) return {};
    return JSON.parse(tokenMap);
  });

  const token = useMemo(() => tokens[roomId] ?? "", [tokens, roomId]);

  const setAccessToken = useCallback(
    (token: string) => {
      setTokens((prev) => {
        const newTokens = { ...prev, [roomId]: token };
        sessionStorage.setItem(accessTokensKey, JSON.stringify(newTokens));
        return newTokens;
      });
    },
    [roomId]
  );

  return [token, setAccessToken] as const;
};

export const saveRoomToken = ({
  roomId,
  token,
}: {
  roomId: string;
  token: string;
}) => {
  const tokenMap = JSON.parse(sessionStorage.getItem(accessTokensKey) ?? "{}");
  tokenMap[roomId] = token;
  sessionStorage.setItem(accessTokensKey, JSON.stringify(tokenMap));
};

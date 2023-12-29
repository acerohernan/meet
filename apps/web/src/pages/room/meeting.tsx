import { useParams } from "react-router-dom";

export const MeetingPage = () => {
  const { roomId } = useParams();

  return <div>MeetingPage: {roomId} </div>;
};

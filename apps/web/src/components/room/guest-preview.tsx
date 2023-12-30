import { useCallback, useMemo, useRef, useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import VideocamOffOutlinedIcon from "@mui/icons-material/VideocamOffOutlined";
import MicOffOutlinedIcon from "@mui/icons-material/MicOffOutlined";
import MicNoneOutlinedIcon from "@mui/icons-material/MicNoneOutlined";
import VideocamOutlinedIcon from "@mui/icons-material/VideocamOutlined";
import { GuestMicVisualizer } from "./guest-mic-visualizer";

enum CameraState {
  OFF,
  STARTING,
  ON,
}

export const GuestPreview = () => {
  const [micOn, setMicOn] = useState(false);
  const [cameraState, setCameraState] = useState<CameraState>(CameraState.OFF);
  const [videoStream, setVideoStream] = useState<MediaStream | null>(null);
  const [audioStream, setAudioStream] = useState<MediaStream | null>(null);

  const videoRef = useRef<HTMLVideoElement | null>(null);

  const toggleMicrophone = useCallback(async () => {
    if (!micOn) {
      try {
        const stream = await navigator.mediaDevices.getUserMedia({
          audio: true,
        });
        setAudioStream(stream);
        setMicOn(true);
      } catch (error) {
        // TODO: show a tooltip
        alert("error at accessing mic");
        console.log(error);
      }
    } else {
      if (audioStream) {
        const tracks = audioStream.getTracks();
        tracks.forEach((t) => t.stop());
        setAudioStream(null);
      }
      setMicOn(false);
    }
  }, [micOn, audioStream]);

  const toggleCamera = useCallback(async () => {
    if (!videoRef.current) return;

    if (cameraState == CameraState.OFF) {
      setCameraState(CameraState.STARTING);
      try {
        const stream = await navigator.mediaDevices.getUserMedia({
          video: true,
        });
        setVideoStream(stream);
        videoRef.current.srcObject = stream;
        videoRef.current.style.display = "flex";
        videoRef.current.play();
        setCameraState(CameraState.ON);
      } catch (error) {
        // TODO: show a tooltip
        alert("error at accessing camera");
        setCameraState(CameraState.OFF);
      }
    } else {
      //if it's starting or on, turn off
      if (videoStream) {
        const tracks = videoStream.getTracks();
        tracks.forEach((t) => t.stop());
      }
      setAudioStream(null);
      videoRef.current.srcObject = null;
      videoRef.current.style.display = "hidden";
      setCameraState(CameraState.OFF);
    }
  }, [videoStream, cameraState]);

  const cameraOn = useMemo<boolean>(
    () => cameraState !== CameraState.OFF,
    [cameraState]
  );

  return (
    <Box
      width="100%"
      height="100%"
      borderRadius={3}
      position="relative"
      sx={{
        background: "#202124",
        display: "grid",
        placeItems: "center",
        aspectRatio: "4/2.5",
        minWidth: "300px",
      }}
    >
      {/* Preview display */}
      {!cameraOn ? (
        <Typography variant="h5" sx={{ color: "white", fontWeight: 300 }}>
          {cameraState == CameraState.OFF ? "Camera is off" : null}
          {cameraState == CameraState.STARTING ? "Camera is starting" : null}
        </Typography>
      ) : null}

      <video
        ref={videoRef}
        style={{
          borderRadius: 12,
          display: "hidden",
          position: "absolute",
          top: 0,
          bottom: 0,
          right: 0,
          left: 0,
          objectFit: "cover",
          width: "100%",
          height: "100%",
        }}
      />

      {/* Preview controls */}
      <Box
        position="absolute"
        top={0}
        bottom={0}
        left={0}
        right={0}
        width="auto"
        borderRadius="inherit"
        display="flex"
        flexDirection="column"
        justifyContent="space-between"
      >
        {micOn ? (
          <Box position="absolute" bottom={0} left={0} margin={2}>
            <GuestMicVisualizer />
          </Box>
        ) : null}

        {/* Top shadow */}
        <Box
          width="100%"
          height="80px"
          sx={{
            backgroundImage:
              "linear-gradient(to bottom,rgba(0,0,0,0.7) 0,rgba(0,0,0,0.3) 50%,rgba(0,0,0,0) 100%)",
          }}
          borderRadius="inherit"
        ></Box>

        {/* Bottom container */}
        <Box
          display="flex"
          alignItems="center"
          justifyContent="center"
          width="100%"
          gap={1}
          padding={1}
          sx={{
            backgroundImage:
              "linear-gradient(to top,rgba(0,0,0,0.7) 0,rgba(0,0,0,0.3) 50%,rgba(0,0,0,0) 100%)",
          }}
          borderRadius="inherit"
        >
          <Button
            aria-label="on/off microphone"
            variant="contained"
            size="small"
            sx={{
              borderRadius: "100%",
              padding: "20px",
              borderColor: micOn ? "white" : "#ea4335",
              borderWidth: "1px",
              borderStyle: "solid",
              backgroundColor: micOn ? "transparent" : "#ea4335",
              "&:hover": { backgroundColor: micOn ? "#9A9A9C" : "#C93C33" },
              transform: "scale(0.8)",
            }}
            onClick={toggleMicrophone}
          >
            {micOn ? (
              <MicNoneOutlinedIcon sx={{ color: "white", fontSize: "2rem" }} />
            ) : (
              <MicOffOutlinedIcon sx={{ color: "white", fontSize: "2rem" }} />
            )}
          </Button>
          <Button
            aria-label="on/off camera"
            variant="contained"
            size="small"
            sx={{
              borderRadius: "100%",
              padding: "20px",
              borderColor: cameraOn ? "white" : "#ea4335",
              borderWidth: "1px",
              borderStyle: "solid",
              backgroundColor: cameraOn ? "transparent" : "#ea4335",
              "&:hover": { backgroundColor: cameraOn ? "#9A9A9C" : "#C93C33" },
              transform: "scale(0.8)",
            }}
            onClick={toggleCamera}
          >
            {cameraOn ? (
              <VideocamOutlinedIcon sx={{ color: "white", fontSize: "2rem" }} />
            ) : (
              <VideocamOffOutlinedIcon
                sx={{ color: "white", fontSize: "2rem" }}
              />
            )}
          </Button>
        </Box>
      </Box>
    </Box>
  );
};

import { RTCService } from "@/lib/service";

export const rtcService = new RTCService(import.meta.env.VITE_SERVER_URL);

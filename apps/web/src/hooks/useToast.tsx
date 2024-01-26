import { toast } from "react-hot-toast";

export const useToast = () => {
  return {
    success: (msg: string) => toast.success(msg, { position: "bottom-right" }),
    error: (msg: string) => toast.error(msg, { position: "bottom-right" }),
    warning: (msg: string) =>
      toast(msg, { icon: "⚠️", position: "bottom-right" }),
    custom: (Element: React.FC) =>
      toast.custom((t) => (t.visible ? <Element /> : null), {
        position: "bottom-right",
      }),
  };
};

import { toast } from "react-hot-toast";

export const useToast = () => {
  return {
    toast: {
      success: (msg: string) => toast.success(msg),
      error: (msg: string) => toast.error(msg),
      warning: (msg: string) => toast(msg, { icon: "⚠️" }),
    },
  };
};

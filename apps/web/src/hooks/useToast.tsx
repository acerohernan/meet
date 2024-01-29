import { toast } from "react-hot-toast";

export const useToast = () => {
  return {
    /** renders a success toast */
    success: (msg: string) => toast.success(msg, { position: "bottom-right" }),
    /** renders an error toast */
    error: (msg: string) => toast.error(msg, { position: "bottom-right" }),
    /** renders a warning toast */
    warning: (msg: string) =>
      toast(msg, { icon: "⚠️", position: "bottom-right" }),
    /** renders custom toast element as toast */
    custom: (Element: React.FC<any>, props?: Record<string, any>) =>
      toast.custom(
        (t) => (t.visible ? <Element toast={t} {...props} /> : null),
        {
          position: "bottom-right",
        }
      ),
    /** removes specific toast from screen */
    dismiss: (toastId: string) => toast.dismiss(toastId),
  };
};

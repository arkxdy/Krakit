export const storage = {
  get: (key: string) => {
    if (typeof window === "undefined") return null;
    return localStorage.getItem(key);
  },
  set: (key: string, value: string) => {
    if (typeof window !== "undefined") localStorage.setItem(key, value);
  },
  remove: (key: string) => {
    if (typeof window !== "undefined") localStorage.removeItem(key);
  },
};

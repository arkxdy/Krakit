import { cookies } from "next/headers";

export const userPrefs = cookies();

export function createAppCookie(name: string, options: any) {
  return { name, ...options };
}

export const jwtCookie = createAppCookie("token", {
  httpOnly: true,
  sameSite: "lax",
  secure: process.env.NODE_ENV === "production",
  path: "/",
  maxAge: 60 * 15, // 15 mins
});

export const examCookie = createAppCookie("exam", {
  httpOnly: true,
  sameSite: "lax",
  secure: process.env.NODE_ENV === "production",
  path: "/",
});

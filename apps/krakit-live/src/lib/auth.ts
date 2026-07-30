import { cookies } from "next/headers";

export interface UserSession {
  id: string;
  name: string;
  email: string;
}

export async function getCurrentUser(): Promise<UserSession | null> {
  // const authToken = cookies().get("krakit_auth_token")?.value;
  const authToken = "dummy_token"; // For demonstration purposes, replace with actual token retrieval logic

  if (!authToken) {
    return null;
  }

  return {
    id: "user-1",
    name: "Demo User",
    email: "demo@example.com",
  };
}

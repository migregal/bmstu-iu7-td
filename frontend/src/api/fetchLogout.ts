import { client } from "./client"

export function fetchLogout() {
  return client.post<null>("/auth/logout", {})
}

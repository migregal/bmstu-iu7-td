import { client } from "./client"

export type FetchUpdateFileBody = {
    id: string
}

/**
 * @param data should have `id` and contain one of `title` and `file` keys.
 */
export function fetchUpdateFile(id: string, data: FormData) {
  return client.patchForm<FetchUpdateFileBody>(`/files/upd/${id}`, data)
}

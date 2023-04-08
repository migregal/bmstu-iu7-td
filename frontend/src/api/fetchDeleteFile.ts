import { client } from "./client"

export type FetchDeleteFileBody = {
    id: string
}

export function fetchDeleteFile(id: string) {
  return client.delete<FetchDeleteFileBody>(`/files/del/${id}`)
}


import { client } from "./client"

export type FetchAllFilesBody = {
    files: Array<{
        id: string
        length: number
        title: string
        url: string
    }>
}

export function fetchAllFiles() {
  return client.get<FetchAllFilesBody>("/files/get")
}

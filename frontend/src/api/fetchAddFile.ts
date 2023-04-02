import { client } from "./client"

export type FetchAddFileBody = {
    id: string
    url: string
}

/**
 * @param data should contain `title` and `file` keys.
 */
export function fetchAddFile(data: FormData) {
  return client.postForm<FetchAddFileBody>("/files/add", data)
}

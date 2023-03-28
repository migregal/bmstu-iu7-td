import { client } from "./client"

type Body = {
    id: string
    url: string
}

/**
 * @param data should contain `title` and `file` keys.
 */
export function fetchAddFile(data: FormData) {
  return client.postForm<Body>("/api/v1/files/add", data)
}
import { client } from "./client"

type Body = {
    login: string
    password: string
}

type Result = {
    token: string
    id: string
}

export function fetchRegistration(body: Body) {
  return client.post<Result>("/auth/registration", body)
}


type Response<B> = {
    data: B | null
    errors: Record<string, string> | null
}

export class ApiClient {
  baseUrl: string
  token: string | null = null

  constructor() {
    this.baseUrl = process.env.REACT_APP_API_URL ?? "/api/v1"
    this.token = window.localStorage?.getItem("auth-token")
  }

  setToken(token: string | null) {
    this.token = token
    if (token) {
      window.localStorage?.setItem("auth-token", token)
    } else {
      window.localStorage?.removeItem("auth-token")
    }
  }

  async post<B>(url: string, body: object): Promise<Response<B>> {
    const headers: Record<string, string> = {
      "Content-Type": "application/json",
    }

    if (this.token) {
      headers["Authorization"] = this.token
    }

    const resp = await fetch(this.baseUrl + url, {
      method: "POST",
      body: JSON.stringify(body),
      headers,    })

    if (resp.status !== 200) {
      console.error("ApiClient.post: wrong status code in response", resp)

      throw new Error("Received wrong response from server")    
    }

    const raw = await resp.json()

    return { data: raw.data ?? null, errors: raw.errors && Object.keys(raw.errors).length > 0 ? raw.errors : null }
  }
}

export const client = new ApiClient()
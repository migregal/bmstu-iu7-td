
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

  private fetch(url: string, init: RequestInit) {
    if (this.token) {
      init.headers = {
        ...init.headers,
        Authorization: `Bearer ${this.token}`,
      }
    }

    return window.fetch(this.baseUrl + url, init)
  }

  private async parseResponseJson(resp: globalThis.Response) {
    const raw = await resp.json()

    return {
      data: raw.data ?? null,
      errors: raw.errors && Object.keys(raw.errors).length > 0 ? raw.errors : null
    }
  }

  async post<B>(url: string, body: object): Promise<Response<B>> {
    const resp = await this.fetch(url, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      }
    })

    if (resp.status !== 200) {
      console.error("ApiClient.post: wrong status code in response", resp)

      throw new Error("Received wrong response from server")
    }

    return this.parseResponseJson(resp)
  }

  async postForm<B>(url: string, body: FormData): Promise<Response<B>> {
    const resp = await this.fetch(url, {
      method: "POST",
      body,
    })

    return this.parseResponseJson(resp)
  }

  async get<B>(url: string): Promise<Response<B>> {
    const resp = await this.fetch(url, {
      method: "GET",
    })

    return this.parseResponseJson(resp)
  }
}

export const client = new ApiClient()

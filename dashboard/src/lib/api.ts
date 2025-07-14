export class ApiClient {
  private baseUrl: string =
    process.env.NODE_ENV === "production"
      ? "/api"
      : "http://localhost:8080/api";

  private async request<T>(
    url: string,
    options: RequestInit,
    token?: string
  ): Promise<T> {
    const headers = new Headers(options.headers);
    headers.set("Authorization", `Bearer ${token}`);
    headers.set("Content-Type", "application/json");
    const response = await fetch(`${this.baseUrl}${url}`, {
      ...options,
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.statusText}`);
    }

    const data = await response.json();
    return data;
  }

  public async get<T>(
    endpoint: string,
    queryParams?: Record<string, string>,
    token?: string
  ): Promise<T> {
    const queryString = queryParams
      ? new URLSearchParams(queryParams).toString()
      : "";
    const url = `${endpoint}${queryString ? `?${queryString}` : ""}`;

    return this.request<T>(
      url,
      {
        method: "GET",
      },
      token
    );
  }

  public async delete<T>(
    endpoint: string,
    queryParams?: Record<string, string>,
    token?: string
  ): Promise<T> {
    const queryString = queryParams
      ? new URLSearchParams(queryParams).toString()
      : "";
    const url = `${endpoint}${queryString ? `?${queryString}` : ""}`;

    return this.request<T>(
      url,
      {
        method: "DELETE",
      },
      token
    );
  }

  public async post<T>(
    endpoint: string,
    body: object,
    token?: string
  ): Promise<T> {
    return this.request<T>(
      endpoint,
      {
        method: "POST",
        body: JSON.stringify(body),
      },
      token
    );
  }
}

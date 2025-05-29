import { useRouter } from "vue-router";

const router = useRouter();

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

type RequestOptions = {
  method?: string;
  body?: any;
  headers?: Record<string, string>;
  auth?: boolean;
};

export async function fetchClient<T = any>(
  endpoint: string,
  options: RequestOptions = {}
): Promise<T> {
  const { method = "GET", body, headers = {}, auth = true } = options;
  let token = localStorage.getItem("access_token");

  async function makeRequest(currentToken: string | null): Promise<Response> {
    return fetch(`${BASE_URL}${endpoint}`, {
      method,
      headers: {
        "Content-Type": "application/json",
        ...(auth && currentToken
          ? { Authorization: `Bearer ${currentToken}` }
          : {}),
        ...headers,
      },
      body: body ? JSON.stringify(body) : undefined,
      credentials: "include",
    });
  }

  let response = await makeRequest(token);

  if (response.status === 401 && auth) {
    const newToken = await refreshAccessToken();

    if (newToken) {
      token = newToken;
      response = await makeRequest(token);
    } else {
      router.push("/login");
      throw new Error("Unauthorized — please log in again");
    }
  }

  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ message: response.statusText }));
    throw new Error(error.error);
  }

  return response.json();
}

async function refreshAccessToken(): Promise<string | null> {
  try {
    const response = await fetch(`${BASE_URL}/auth/refresh`, {
      method: "POST",
      credentials: "include",
    });

    if (!response.ok) return null;

    const data = await response.json();
    localStorage.setItem("access_token", data.access_token);
    return data.access_token;
  } catch (error) {
    return null;
  }
}

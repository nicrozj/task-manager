import type { Router } from "vue-router";
import { router } from "@/main";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

interface RefreshTokenResponse {
  success: boolean;
  status: number;
  data: {
    access_token: string;
  };
}

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

  let isRefreshing = false;

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
    if (!isRefreshing) {
      isRefreshing = true;

      try {
        const refreshTokenResponse = await refreshAccessToken();

        if (!refreshTokenResponse?.data.access_token) {
          throw new Error("Access token not found");
        }

        const newAccessToken = refreshTokenResponse.data.access_token;
        localStorage.setItem("access_token", newAccessToken);

        response = await makeRequest(newAccessToken);
      } catch (error) {
        localStorage.removeItem("access_token");
        router.push("/login");
        throw new Error("Session expired");
      } finally {
        isRefreshing = false;
      }
    } else {
      throw new Error("Token refresh in progress");
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

async function refreshAccessToken(): Promise<RefreshTokenResponse | null> {
  return await fetchClient("/auth/refresh", {
    method: "POST",
    auth: true,
  });
}

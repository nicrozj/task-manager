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

let isRefreshing = false;
let refreshPromise: Promise<string | null> | null = null;

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
    try {
      if (isRefreshing && refreshPromise) {
        const newToken = await refreshPromise;
        if (!newToken) throw new Error("Token refresh failed");
        response = await makeRequest(newToken);
      } else {
        isRefreshing = true;
        refreshPromise = refreshAccessToken();

        const newToken = await refreshPromise;
        isRefreshing = false;
        refreshPromise = null;

        if (!newToken) {
          localStorage.removeItem("access_token");
          router.push("/login");
          throw new Error("Session expired");
        }

        localStorage.setItem("access_token", newToken);
        response = await makeRequest(newToken);
      }
    } catch (error) {
      isRefreshing = false;
      refreshPromise = null;
      localStorage.removeItem("access_token");
      router.push("/login");
      throw new Error("Session expired");
    }
  }

  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ message: response.statusText }));
    throw new Error(error?.error || error?.message || "Unknown error");
  }

  return response.json();
}

async function refreshAccessToken(): Promise<string | null> {
  try {
    const response = await fetch(`${BASE_URL}/auth/refresh`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) return null;

    const result: RefreshTokenResponse = await response.json();
    return result?.data?.access_token ?? null;
  } catch (error) {
    return null;
  }
}

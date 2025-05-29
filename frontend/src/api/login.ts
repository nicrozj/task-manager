import { fetchClient } from "./fetchClient";
import type { Router } from "vue-router";

export function login(payload: { username: string; password: string }) {
  return fetchClient("/auth/login", {
    method: "POST",
    body: payload,
    auth: false,
  }).then((response) => {
    if (response.data.access_token) {
      console.log(response.data.access_token);
      localStorage.setItem("access_token", response.data.access_token);
    }
    return response;
  });
}

export function register(payload: { username: string; password: string }) {
  return fetchClient("/auth/registration", {
    method: "POST",
    body: payload,
    auth: false,
  });
}

export function getProfile() {
  return fetchClient("/auth/me");
}

export function logout() {
  return fetchClient("/auth/logout", {
    method: "POST",
  });
}

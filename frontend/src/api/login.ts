import { fetchClient } from "./fetchClient";

export function login(payload: { username: string; password: string }) {
  return fetchClient("/auth/login", {
    method: "POST",
    body: payload,
    auth: false,
  }).then((data) => {
    if (data.access_token) {
      localStorage.setItem("access_token", data.access_token);
    }
    return data;
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

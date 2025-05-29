import { fetchClient } from "./fetchClient";
import type { Router } from "vue-router";

export function login(
  payload: { username: string; password: string },
  router: Router
) {
  return fetchClient(
    "/auth/login",
    {
      method: "POST",
      body: payload,
      auth: false,
    },
    router
  ).then((response) => {
    if (response.data.access_token) {
      console.log(response.data.access_token);
      localStorage.setItem("access_token", response.data.access_token);
    }
    return response;
  });
}

export function register(
  payload: { username: string; password: string },
  router: Router
) {
  return fetchClient(
    "/auth/registration",
    {
      method: "POST",
      body: payload,
      auth: false,
    },
    router
  );
}

export function getProfile(router: Router) {
  return fetchClient("/auth/me", {}, router);
}

export function logout(router: Router) {
  return fetchClient(
    "/auth/logout",
    {
      method: "POST",
    },
    router
  );
}

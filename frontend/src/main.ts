import { createApp } from "vue";
import { createRouter, createWebHistory, RouterView } from "vue-router";
import Index from "./pages/index.vue";
import login from "./pages/login.vue";
import registration from "./pages/registration.vue";

import "./main.css";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Index, meta: { requiresAuth: true } },
    { path: "/registration", component: registration },
    { path: "/login", component: login },
  ],
});

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem("access_token");

  if (to.meta.requiresAuth && !isAuthenticated) {
    next("/login");
  } else {
    next();
  }
});

createApp(RouterView).use(router).mount("body");

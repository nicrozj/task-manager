import { type Task } from "@/composables/useTasks";
import type { Router } from "vue-router";
import { fetchClient } from "./fetchClient";

export function getTasks() {
  return fetchClient("/tasks", {
    method: "GET",
    auth: true,
  });
}

export function getTaskById(id: number) {
  return fetchClient(`/tasks/${id}`, {
    method: "GET",
    auth: true,
  });
}

export function createTask(payload: Task) {
  return fetchClient("/tasks", {
    method: "POST",
    body: payload,
    auth: true,
  });
}

export function updateTask(payload: Task, id: number) {
  return fetchClient(`/tasks/${id}`, {
    method: "PUT",
    body: payload,
    auth: true,
  });
}

export function deleteTask(id: number) {
  return fetchClient(`/tasks/${id}`, {
    method: "DELETE",
    auth: true,
  });
}

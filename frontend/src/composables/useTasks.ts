import { getTasks } from "@/api/tasks";
import { ref } from "vue";

export type status = "new" | "in-progress" | "completed";

const tasks = ref<Task[]>([]);

export interface Task {
  id?: number;
  title: string;
  description: string;
  status: status;
  created_at?: string;
}

export const useTasks = () => {
  async function fetchTasks() {
    const response = await getTasks();
    if (response.data) {
      tasks.value.push(...response.data);
    }
  }

  function updateTaskInList(task: Task) {
    const index = tasks.value.findIndex((t) => t.id === task.id);

    if (index !== -1) {
      tasks.value[index] = task;
    } else {
      tasks.value.push(task);
    }
  }

  function deleteTaskInList(taskId: number) {
    tasks.value = tasks.value.filter((t) => t.id !== taskId);
  }
  return {
    tasks,
    fetchTasks,
    updateTaskInList,
    deleteTaskInList,
  };
};

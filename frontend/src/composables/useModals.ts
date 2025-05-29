import { ref, onMounted, onUnmounted } from "vue";
import type { Task } from "./useTasks";

const isOpen = ref(false);
const editingTask = ref<Task | null>(null);

export function useTaskModal() {
  const open = (task: Task | null = null) => {
    editingTask.value = task;
    isOpen.value = true;
  };

  const close = () => {
    isOpen.value = false;
    editingTask.value = null;
  };

  const handleEscape = (e: KeyboardEvent) => {
    if (e.key === "Escape") {
      close();
    }
  };

  onMounted(() => {
    document.addEventListener("keydown", handleEscape);
  });

  onUnmounted(() => {
    document.removeEventListener("keydown", handleEscape);
  });

  return {
    isOpen,
    editingTask,
    open,
    close,
  };
}

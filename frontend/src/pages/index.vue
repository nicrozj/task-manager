<script setup lang="ts">
import Header from "@/components/Header.vue";
import TaskModal from "@/components/tasks/TaskModal.vue";
import { ref, computed, reactive, onMounted } from "vue";
import { useTaskModal } from "@/composables/useModals";
import TaskFilter from "@/components/tasks/TaskFilter.vue";
import VStack from "@/components/stacks/VStack.vue";
import TaskCard from "@/components/tasks/TaskCard.vue";
import { getTasks } from "@/api/tasks";
import { useTasks, type Task } from "@/composables/useTasks";
import UButton from "@/components/UButton.vue";

const { isOpen, open, close } = useTaskModal();
const { tasks, fetchTasks, updateTaskInList } = useTasks();

const filters = ref({
  status: "all" as "all" | "new" | "in-progress" | "completed",
  date: "",
});

const filteredTasks = computed(() => {
  return tasks.value.filter((task) => {
    const statusMatch =
      filters.value.status === "all" || task.status === filters.value.status;

    const dateMatch =
      !filters.value.date ||
      (task.created_at &&
        new Date(task.created_at).toLocaleDateString() ===
          new Date(filters.value.date).toLocaleDateString());

    return statusMatch && dateMatch;
  });
});

onMounted(() => {
  tasks.value = [];
  fetchTasks();
});
</script>
<template>
  <Header />
  <VStack class="w-[1000px] mx-auto mt-20 gap-4">
    <span class="text-3xl">Задачи:</span>
    <TaskFilter @update:filters="(newFilters) => (filters = newFilters)" />
    <div class="grid grid-cols-2 gap-12">
      <div v-for="(task, id) in filteredTasks">
        <TaskCard :task="task" />
      </div>
    </div>
  </VStack>
  <TaskModal v-if="isOpen" />
</template>

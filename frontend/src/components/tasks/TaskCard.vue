<script setup lang="ts">
import { type Task } from "@/composables/useTasks";
import { useTaskModal } from "@/composables/useModals";
import HStack from "../stacks/HStack.vue";

const { isOpen, open, close } = useTaskModal();

interface Props {
  task: Task;
}

const props = defineProps<Props>();

const statusStyles = {
  new: "bg-green-200 text-green-800",
  "in-progress": "bg-blue-200 text-blue-800",
  completed: "bg-red-200 text-red-800",
};

const statusLabels = {
  new: "Новая",
  "in-progress": "В работе",
  completed: "Завершена",
};

function formatDate(isoString: string): string {
  const date = new Date(isoString);
  return `${date.getDate()}.${
    date.getMonth() + 1
  }.${date.getFullYear()} ${date.getHours()}:${date
    .getMinutes()
    .toString()
    .padStart(2, "0")}`;
}
</script>
<template>
  <div class="bg-white rounded-lg p-4 transition-shadow border border-gray-100">
    <div class="flex justify-between items-start">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">
          {{ props.task?.title }}
        </h3>
        <p class="mt-1 text-sm text-gray-600 line-clamp-6">
          {{ props.task.description }}
        </p>
      </div>
      <span
        :class="statusStyles[props.task.status]"
        class="px-2 py-1 rounded-md"
        >{{ statusLabels[props.task.status] }}</span
      >
    </div>

    <HStack class="mt-4 flex justify-between">
      <span>{{ formatDate(props.task.created_at!) }}</span>
      <button
        @click="open(task)"
        class="text-sm px-3 py-1 bg-indigo-100 text-indigo-700 rounded hover:bg-indigo-200 transition-colors"
      >
        Редактировать
      </button>
    </HStack>
  </div>
</template>
